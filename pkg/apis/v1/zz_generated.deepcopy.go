// +build !ignore_autogenerated

// Copyright 2020 Acnodal, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*ServiceGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ServiceGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Agents != nil {
		in, out := &in.Agents, &out.Agents
		*out = make([]*LBNodeAgent, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(LBNodeAgent)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LBNodeAgent) DeepCopyInto(out *LBNodeAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LBNodeAgent.
func (in *LBNodeAgent) DeepCopy() *LBNodeAgent {
	if in == nil {
		return nil
	}
	out := new(LBNodeAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LBNodeAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LBNodeAgentEPICSpec) DeepCopyInto(out *LBNodeAgentEPICSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LBNodeAgentEPICSpec.
func (in *LBNodeAgentEPICSpec) DeepCopy() *LBNodeAgentEPICSpec {
	if in == nil {
		return nil
	}
	out := new(LBNodeAgentEPICSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LBNodeAgentList) DeepCopyInto(out *LBNodeAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LBNodeAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LBNodeAgentList.
func (in *LBNodeAgentList) DeepCopy() *LBNodeAgentList {
	if in == nil {
		return nil
	}
	out := new(LBNodeAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LBNodeAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LBNodeAgentLocalSpec) DeepCopyInto(out *LBNodeAgentLocalSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LBNodeAgentLocalSpec.
func (in *LBNodeAgentLocalSpec) DeepCopy() *LBNodeAgentLocalSpec {
	if in == nil {
		return nil
	}
	out := new(LBNodeAgentLocalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LBNodeAgentSpec) DeepCopyInto(out *LBNodeAgentSpec) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LBNodeAgentLocalSpec)
		**out = **in
	}
	if in.EPIC != nil {
		in, out := &in.EPIC, &out.EPIC
		*out = new(LBNodeAgentEPICSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LBNodeAgentSpec.
func (in *LBNodeAgentSpec) DeepCopy() *LBNodeAgentSpec {
	if in == nil {
		return nil
	}
	out := new(LBNodeAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LBNodeAgentStatus) DeepCopyInto(out *LBNodeAgentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LBNodeAgentStatus.
func (in *LBNodeAgentStatus) DeepCopy() *LBNodeAgentStatus {
	if in == nil {
		return nil
	}
	out := new(LBNodeAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceGroup) DeepCopyInto(out *ServiceGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceGroup.
func (in *ServiceGroup) DeepCopy() *ServiceGroup {
	if in == nil {
		return nil
	}
	out := new(ServiceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceGroupEPICSpec) DeepCopyInto(out *ServiceGroupEPICSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceGroupEPICSpec.
func (in *ServiceGroupEPICSpec) DeepCopy() *ServiceGroupEPICSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceGroupEPICSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceGroupList) DeepCopyInto(out *ServiceGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceGroupList.
func (in *ServiceGroupList) DeepCopy() *ServiceGroupList {
	if in == nil {
		return nil
	}
	out := new(ServiceGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceGroupLocalSpec) DeepCopyInto(out *ServiceGroupLocalSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceGroupLocalSpec.
func (in *ServiceGroupLocalSpec) DeepCopy() *ServiceGroupLocalSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceGroupLocalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceGroupNetboxSpec) DeepCopyInto(out *ServiceGroupNetboxSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceGroupNetboxSpec.
func (in *ServiceGroupNetboxSpec) DeepCopy() *ServiceGroupNetboxSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceGroupNetboxSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceGroupSpec) DeepCopyInto(out *ServiceGroupSpec) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(ServiceGroupLocalSpec)
		**out = **in
	}
	if in.EPIC != nil {
		in, out := &in.EPIC, &out.EPIC
		*out = new(ServiceGroupEPICSpec)
		**out = **in
	}
	if in.Netbox != nil {
		in, out := &in.Netbox, &out.Netbox
		*out = new(ServiceGroupNetboxSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceGroupSpec.
func (in *ServiceGroupSpec) DeepCopy() *ServiceGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceGroupStatus) DeepCopyInto(out *ServiceGroupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceGroupStatus.
func (in *ServiceGroupStatus) DeepCopy() *ServiceGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceGroupStatus)
	in.DeepCopyInto(out)
	return out
}
