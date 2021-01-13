// Copyright 2020 Acnodal Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package local

import (
	"fmt"
	"net"

	v1 "k8s.io/api/core/v1"

	"purelb.io/internal/election"
	"purelb.io/internal/k8s"
	"purelb.io/internal/lbnodeagent"
	purelbv1 "purelb.io/pkg/apis/v1"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vishvananda/netlink"
)

type announcer struct {
	client      k8s.ServiceEvent
	logger      log.Logger
	myNode      string
	config      *purelbv1.LBNodeAgentLocalSpec
	groups      map[string]*purelbv1.ServiceGroupLocalSpec // groupName -> ServiceGroupLocalSpec
	svcAdvs     map[string]net.IP                          // svcName -> IP
	election    *election.Election
	announceInt *netlink.Link // for local announcements
	dummyInt    *netlink.Link // for non-local announcements
}

var announcing = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: purelbv1.MetricsNamespace,
	Subsystem: "lbnodeagent",
	Name:      "announced",
	Help:      "Services announced from this node",
}, []string{
	"service",
	"node",
	"ip",
})

func init() {
	prometheus.MustRegister(announcing)
}

// NewAnnouncer returns a new local Announcer.
func NewAnnouncer(l log.Logger, node string) lbnodeagent.Announcer {
	return &announcer{logger: l, myNode: node, svcAdvs: map[string]net.IP{}}
}

// SetClient configures this announcer to use the provided client.
func (a *announcer) SetClient(client *k8s.Client) {
	a.client = client
}

func (a *announcer) SetConfig(cfg *purelbv1.Config) error {
	// the default is nil which means that we don't announce
	a.config = nil

	// if there's a "Local" agent config then we'll announce
	for _, agent := range cfg.Agents {
		if spec := agent.Spec.Local; spec != nil {
			a.logger.Log("op", "setConfig", "config", spec)
			a.config = spec

			// stash the local service group configs
			a.groups = map[string]*purelbv1.ServiceGroupLocalSpec{}
			for _, group := range cfg.Groups {
				if group.Spec.Local != nil {
					a.groups[group.ObjectMeta.Name] = group.Spec.Local
				}
			}

			// if the user specified an interface then we'll use that,
			// otherwise we'll wait until we get an IP address so we can
			// figure out the default interface
			if spec.LocalInterface != "default" {
				link, err := netlink.LinkByName(spec.LocalInterface)
				if err != nil {
					return err
				}
				a.announceInt = &link
			}

			// now that we've got a config we can create the dummy interface
			var err error
			fmt.Println("***Adding dummy int")
			if a.dummyInt, err = addDummyInterface(spec.ExtLBInterface); err != nil {
				return err
			}

			// we've got our marching orders so we don't need to continue
			// scanning
			return nil
		}
	}

	if a.config == nil {
		a.logger.Log("event", "noConfig")
	}

	return nil
}

func (a *announcer) SetBalancer(svc *v1.Service, endpoints *v1.Endpoints) error {
	var err error
	l := log.With(a.logger, "service", svc.Name)

	// if we haven't been configured then we won't announce
	if a.config == nil {
		l.Log("event", "noConfig")
		return nil
	}

	// validate the allocated address
	lbIP := net.ParseIP(svc.Status.LoadBalancer.Ingress[0].IP)
	if lbIP == nil {
		l.Log("op", "setBalancer", "error", "invalid LoadBalancer IP", "ip", svc.Status.LoadBalancer.Ingress[0].IP)
		return nil
	}

	// If the user specified an announcement interface then use that,
	// otherwise figure out a default
	announceInt := a.announceInt
	if announceInt == nil {
		announceInt, err = defaultInterface(addrFamily(lbIP))
		if err != nil {
			l.Log("event", "announceError", "err", err)
			return err
		}
	}

	if lbIPNet, defaultif, err := checkLocal(announceInt, lbIP); err == nil {

		// Local addresses do not support ExternalTrafficPolicyLocal
		// Set the service back to ExternalTrafficPolicyCluster if adding to local interface

		if svc.Spec.ExternalTrafficPolicy == v1.ServiceExternalTrafficPolicyTypeLocal {
			l.Log("op", "setBalancer", "error", "ExternalTrafficPolicy Local not supported on local Interfaces, setting to Cluster")
			svc.Spec.ExternalTrafficPolicy = v1.ServiceExternalTrafficPolicyTypeCluster
			return a.DeleteBalancer(svc.Name, "ClusterLocal")
		}

		// the service address is local, i.e., it's within the same subnet
		// as our primary interface.  We can announce the address if we
		// win the election
		if winner := a.election.Winner(lbIP.String()); winner == a.myNode {

			// we won the election so we'll add the service address to our
			// node's default interface so linux will respond to ARP
			// requests for it
			l.Log("msg", "Winner, winner, Chicken dinner", "node", a.myNode, "service", svc.Name)
			a.client.Infof(svc, "AnnouncingLocal", "Node %s announcing %s on interface %s", a.myNode, lbIP, defaultif.Attrs().Name)

			addNetwork(lbIPNet, defaultif)
			svc.Annotations[purelbv1.NodeAnnotation] = a.myNode
			svc.Annotations[purelbv1.IntAnnotation] = defaultif.Attrs().Name
			announcing.With(prometheus.Labels{
				"service": svc.Name,
				"node":    a.myNode,
				"ip":      lbIP.String(),
			}).Set(1)
		} else {
			// We lost the election so we'll withdraw any announcement that
			// we might have been making
			l.Log("msg", "notWinner", "node", a.myNode, "winner", winner, "service", svc.Name)
			return a.DeleteBalancer(svc.Name, "lostElection")
		}
	} else {

		fmt.Println("*** Addr is not local :", lbIP.String())
		// The service address is non-local, i.e., it's not on the same
		// subnet as our default interface.

		// Should we announce?
		// No, if externalTrafficPolicy is Local && there's no ready local endpoint
		// Yes, in all other cases
		if svc.Spec.ExternalTrafficPolicy == v1.ServiceExternalTrafficPolicyTypeLocal && !nodeHasHealthyEndpoint(endpoints, a.myNode) {
			l.Log("msg", "policyLocalNoEndpoints", "node", a.myNode, "service", svc.Name)
			return a.DeleteBalancer(svc.Name, "noEndpoints")
		}

		// add this address to the "dummy" interface so routing software
		// (e.g., bird) will announce routes for it
		poolName, gotName := svc.Annotations[purelbv1.PoolAnnotation]
		if gotName {
			allocPool := a.groups[poolName]
			l.Log("msg", "announcingNonLocal", "node", a.myNode, "service", svc.Name, "reason", err)
			a.client.Infof(svc, "AnnouncingNonLocal", "Announcing %s from node %s interface %s", lbIP, a.myNode, (*a.dummyInt).Attrs().Name)
			addVirtualInt(lbIP, *a.dummyInt, allocPool.Subnet, allocPool.Aggregation)
			announcing.With(prometheus.Labels{
				"service": svc.Name,
				"node":    a.myNode,
				"ip":      lbIP.String(),
			}).Set(1)
		}
	}

	// add the address to our announcement database
	a.svcAdvs[svc.Namespace+"/"+svc.Name] = lbIP

	return nil
}

func (a *announcer) DeleteBalancer(name string, reason string) error {

	fmt.Println("***Deleting Service: ", name)

	// if the service isn't in our database then we weren't announcing
	// it so we can't withdraw the address but it's OK
	svcAddr, ok := a.svcAdvs[name]
	if !ok {
		return nil
	}

	fmt.Println("*** local service array: ", a.svcAdvs)

	// delete the service from Prometheus, i.e., it won't show up in the
	// metrics anymore
	announcing.Delete(prometheus.Labels{
		"service": name,
		"node":    a.myNode,
		"ip":      svcAddr.String(),
	})

	// delete this service from our announcement database
	delete(a.svcAdvs, name)

	// if any other service is still using that address then we don't
	// want to withdraw it
	for _, addr := range a.svcAdvs {
		if addr.Equal(svcAddr) {
			a.logger.Log("event", "withdrawAnnouncement", "service", name, "reason", reason, "msg", "ip in use by other service")
			return nil
		}
	}

	a.logger.Log("event", "withdrawAnnouncement", "msg", "Delete balancer", "service", name, "reason", reason)
	deleteAddr(svcAddr)

	return nil
}

// Shutdown cleans up changes that we've made to the local networking
// configuration.
func (a *announcer) Shutdown() {
	// withdraw any announcements that we have made
	for _, ip := range a.svcAdvs {
		deleteAddr(ip)
	}

	// remove the "dummy" interface
	removeInterface(a.dummyInt)
}

func (a *announcer) SetElection(election *election.Election) {
	a.election = election
}

// nodeHasHealthyEndpoint returns true if node has at least one
// healthy endpoint.
func nodeHasHealthyEndpoint(eps *v1.Endpoints, node string) bool {
	ready := map[string]bool{}
	for _, subset := range eps.Subsets {
		for _, ep := range subset.Addresses {
			if ep.NodeName == nil || *ep.NodeName != node {
				continue
			}
			if _, ok := ready[ep.IP]; !ok {
				// Only set true if nothing else has expressed an
				// opinion. This means that false will take precedence
				// if there's any unready ports for a given endpoint.
				ready[ep.IP] = true
			}
		}
		for _, ep := range subset.NotReadyAddresses {
			ready[ep.IP] = false
		}
	}

	for _, r := range ready {
		if r {
			// At least one fully healthy endpoint on this node
			return true
		}
	}
	return false
}
