// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedTracing) DeepCopyInto(out *DistributedTracing) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedTracing.
func (in *DistributedTracing) DeepCopy() *DistributedTracing {
	if in == nil {
		return nil
	}
	out := new(DistributedTracing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DistributedTracing) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedTracingAgentSpec) DeepCopyInto(out *DistributedTracingAgentSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedTracingAgentSpec.
func (in *DistributedTracingAgentSpec) DeepCopy() *DistributedTracingAgentSpec {
	if in == nil {
		return nil
	}
	out := new(DistributedTracingAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedTracingCollectorSpec) DeepCopyInto(out *DistributedTracingCollectorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedTracingCollectorSpec.
func (in *DistributedTracingCollectorSpec) DeepCopy() *DistributedTracingCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(DistributedTracingCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedTracingList) DeepCopyInto(out *DistributedTracingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DistributedTracing, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedTracingList.
func (in *DistributedTracingList) DeepCopy() *DistributedTracingList {
	if in == nil {
		return nil
	}
	out := new(DistributedTracingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DistributedTracingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedTracingSpec) DeepCopyInto(out *DistributedTracingSpec) {
	*out = *in
	if in.Agent != nil {
		in, out := &in.Agent, &out.Agent
		*out = new(DistributedTracingAgentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Pipeline != nil {
		in, out := &in.Pipeline, &out.Pipeline
		*out = make([]DistributedTracingCollectorSpec, len(*in))
		copy(*out, *in)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(DistributedTracingStorageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedTracingSpec.
func (in *DistributedTracingSpec) DeepCopy() *DistributedTracingSpec {
	if in == nil {
		return nil
	}
	out := new(DistributedTracingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedTracingStatus) DeepCopyInto(out *DistributedTracingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedTracingStatus.
func (in *DistributedTracingStatus) DeepCopy() *DistributedTracingStatus {
	if in == nil {
		return nil
	}
	out := new(DistributedTracingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedTracingStorageElasticsearchSpec) DeepCopyInto(out *DistributedTracingStorageElasticsearchSpec) {
	*out = *in
	if in.URLs != nil {
		in, out := &in.URLs, &out.URLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedTracingStorageElasticsearchSpec.
func (in *DistributedTracingStorageElasticsearchSpec) DeepCopy() *DistributedTracingStorageElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(DistributedTracingStorageElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedTracingStorageSpec) DeepCopyInto(out *DistributedTracingStorageSpec) {
	*out = *in
	if in.Elasticsearch != nil {
		in, out := &in.Elasticsearch, &out.Elasticsearch
		*out = new(DistributedTracingStorageElasticsearchSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedTracingStorageSpec.
func (in *DistributedTracingStorageSpec) DeepCopy() *DistributedTracingStorageSpec {
	if in == nil {
		return nil
	}
	out := new(DistributedTracingStorageSpec)
	in.DeepCopyInto(out)
	return out
}
