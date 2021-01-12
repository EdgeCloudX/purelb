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

package v1

const (
	// Annotations that the user can set to drive PureLB.

	// SharingAnnotation is the key for the annotation that indicates
	// whether a service can share its IP address with other
	// services. If two or more services have the same value in their
	// SharingAnnotation, and if they use different ports, then they can
	// share their IP address.
	SharingAnnotation string = "purelb.io/allow-shared-ip"

	// DesiredGroupAnnotation is the key for the annotation that
	// indicates the pool from which the user would like PureLB to
	// allocate this service's IP address.
	DesiredGroupAnnotation string = "purelb.io/service-group"

	// Annotations that PureLB sets that might be useful to users.

	// BrandAnnotation is the key for the PureLB "brand" annotation.
	// It's set when PureLB allocates an IP address for a service.
	BrandAnnotation string = "purelb.io/allocated-by"

	// Brand is the brand name of this product. It's the value of the
	// BrandAnnotation annotation.
	Brand string = "PureLB"

	// PoolAnnotation is the key for the annotation that indicates from
	// which pool the IP address was allocated. Pools are configured by
	// the PureLB ServiceGroup custom resource.
	PoolAnnotation string = "purelb.io/allocated-from"

	// NodeAnnotation is the key for the annotation that indicates which
	// node is announcing this service's IP address.
	NodeAnnotation string = "purelb.io/announcing-node"

	// IntAnnotation is the key for the annotation that indicates which
	// interface is announcing this service's IP address.
	IntAnnotation string = "purelb.io/announcing-interface"

	// Internal annotations that probably aren't useful to users.

	// GroupAnnotation describes the URL of the EGW service group to
	// which this service belongs. It's used internally by PureLB and is
	// unlikely to be useful to the user.
	GroupAnnotation string = "acnodal.io/groupURL"

	// ServiceAnnotation describes this service's EGW URL. It's used
	// internally by PureLB and is unlikely to be useful to the user.
	ServiceAnnotation string = "acnodal.io/serviceURL"

	// EndpointAnnotation describes the EGW URL used to create endpoints
	// for this service. It's used internally by PureLB and is unlikely
	// to be useful to the user.
	EndpointAnnotation string = "acnodal.io/endpointcreateURL"
)