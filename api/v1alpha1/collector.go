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
	// CollectorMode represents how this collector should be deployed (daemonset, deployment, statefulset)
	// +kubebuilder:validation:Enum=daemonset;deployment;statefulset
	CollectorMode string

	// CollectorType represents which collector should be used (jaeger, opentelemetry). Jaeger can only be used as the final collector in a pipeline.
	// +kubebuilder:validation:Enum=jaeger;opentelemetry;loadbalancer
	CollectorType string
)

const (
	// CollectorModeDaemonSet specifies that the collector should be deployed as a Kubernetes DaemonSet.
	CollectorModeDaemonSet CollectorMode = "daemonset"

	// CollectorModeDeployment specifies that the collector should be deployed as a Kubernetes Deployment.
	CollectorModeDeployment CollectorMode = "deployment"

	// CollectorModeDeployment specifies that the collector should be deployed as a Kubernetes StatefulSet.
	CollectorModeStatefulSet CollectorMode = "statefulset"

	// CollectorTypeJaeger specifies that the Jaeger Agent should be used as the agent for this backend.
	CollectorTypeJaeger CollectorType = "jaeger"

	// CollectorTypeOpenTelemetry specifies that OpenTelemetry Collector should be used as the agent for this backend.
	CollectorTypeOpenTelemetry CollectorType = "opentelemetry"

	// CollectorTypeLoadbalancer specifies that an OpenTelemetry Collector load balancer should be used. This works best when it's used as the first entry in the pipeline.
	CollectorTypeLoadbalancer CollectorType = "loadbalancer"
)
