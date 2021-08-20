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

// DistributedTracingSpec defines the desired state of DistributedTracing.
type DistributedTracingSpec struct {
	// Agent defines the configuration for the agents used as part of this distributed tracing backend. Optional. Default: no agent.
	Agent *DistributedTracingAgentSpec `json:"agent,omitempty"`

	// Pipeline specifies how the collection pipeline should be for this distributed tracing backend. Required, with at least one collector.
	Pipeline []DistributedTracingCollectorSpec `json:"pipeline,omitempty"`

	// Storage defines the configuration for the persistent storage used as part of this distributed tracing backend. Optional. Default: in-memory if the last collector in the pipeline is Jaeger
	Storage *DistributedTracingStorageSpec `json:"storage,omitempty"`
}

// DistributedTracingSpec defines the desired state of DistributedTracing.
type DistributedTracingAgentSpec struct {
	// Enabled specifies whether we should deploy an agent for this distributed tracing backend. Optional. Default: true.
	Enabled *bool `json:"enabled,omitempty"`

	// Mode represents how the agent should be deployed (daemonset, sidecar). Optional. Default: sidecar.
	Mode AgentMode `json:"mode,omitempty"`

	// Type represents which agent should be used (jaeger, opentelemetry). Optional. Default: opentelemetry.
	Type AgentType `json:"type,omitempty"`
}

// DistributedTracingCollectorSpec defines how a collector of the pipeline should look like.
type DistributedTracingCollectorSpec struct {
	// Name specifies a name for this entry in the pipeline.
	Name string `json:"name,omitempty"`

	// Mode represents how the collector should be deployed (daemonset, deployment, statefulset).
	Mode CollectorMode `json:"mode,omitempty"`

	// Type represents which collector should be used (jaeger, opentelemetry, loadbalancer).
	Type CollectorType `json:"type,omitempty"`
}

// DistributedTracingStorageSpec defines where the tracing data should be persisted. At this moment, this requires Jaeger to be the last collector in the pipeline and only Elasticsearch is supported as the storage type.
type DistributedTracingStorageSpec struct {
	// Enabled specifies whether we should deploy the storage for this distributed tracing backend. When set to false, we will still use the other properties to configure the pipeline to send data to the specified location, assuming that the storage is being managed by other process.
	Enabled bool `json:"enabled,omitempty"`

	// Type represents which storage to use for this distributed tracing backend. Supported values: elasticsearch.
	Type StorageType `json:"type,omitempty"`

	// Elasticsearch defines the configuration for the Elasticsearch storage used as part of this distributed tracing backend.
	Elasticsearch *DistributedTracingStorageElasticsearchSpec `json:"elasticsearch,omitempty"`
}

// DistributedTracingStorageElasticsearchSpec defines the Elasticsearch configuration for this backend.
type DistributedTracingStorageElasticsearchSpec struct {
	// URLs defines the Elasticsearch URLs to use.
	URLs []string `json:"urls,omitempty"`
}

// DistributedTracingStatus defines the observed state of DistributedTracing.
type DistributedTracingStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DistributedTracing is the Schema for the distributedtracings API.
type DistributedTracing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DistributedTracingSpec   `json:"spec,omitempty"`
	Status DistributedTracingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DistributedTracingList contains a list of DistributedTracing.
type DistributedTracingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DistributedTracing `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DistributedTracing{}, &DistributedTracingList{})
}
