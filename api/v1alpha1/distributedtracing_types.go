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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DistributedTracingSpec defines the desired state of DistributedTracing
type DistributedTracingSpec struct {
	// AgentType defines whether the agent should be an OpenTelemetry Collector or a Jaeger Agent. Possible values: "jaeger" or "opentelemetry".
	AgentType string `json:"agentType,omitempty"`
}

// DistributedTracingStatus defines the observed state of DistributedTracing
type DistributedTracingStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DistributedTracing is the Schema for the distributedtracings API
type DistributedTracing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DistributedTracingSpec   `json:"spec,omitempty"`
	Status DistributedTracingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DistributedTracingList contains a list of DistributedTracing
type DistributedTracingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DistributedTracing `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DistributedTracing{}, &DistributedTracingList{})
}
