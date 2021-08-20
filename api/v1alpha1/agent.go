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

type (
	// AgentMode represents how the agent should be deployed (daemonset, sidecar)
	// +kubebuilder:validation:Enum=daemonset;sidecar
	AgentMode string

	// AgentType represents which agent should be used (jaeger, opentelemetry)
	// +kubebuilder:validation:Enum=jaeger;opentelemetry
	AgentType string
)

const (
	// AgentModeDaemonSet specifies that the collector should be deployed as a Kubernetes DaemonSet.
	AgentModeDaemonSet AgentMode = "daemonset"

	// AgentModeSidecar specifies that the collector should be deployed as a sidecar to pods.
	AgentModeSidecar AgentMode = "sidecar"

	// AgentTypeJaeger specifies that the Jaeger Agent should be used as the agent for this backend.
	AgentTypeJaeger AgentType = "jaeger"

	// AgentTypeOpenTelemetry specifies that OpenTelemetry Collector should be used as the agent for this backend.
	AgentTypeOpenTelemetry AgentType = "opentelemetry"
)
