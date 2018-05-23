/*
Copyright 2018 The Kubernetes Authors All rights reserved.

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

package options

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	DefaultNamespaces = NamespaceList{metav1.NamespaceAll}
	DefaultCollectors = CollectorSet{
		"daemonsets":               struct{}{},
		"deployments":              struct{}{},
		"limitranges":              struct{}{},
		"nodes":                    struct{}{},
		"pods":                     struct{}{},
		"replicasets":              struct{}{},
		"replicationcontrollers":   struct{}{},
		"resourcequotas":           struct{}{},
		"services":                 struct{}{},
		"jobs":                     struct{}{},
		"cronjobs":                 struct{}{},
		"statefulsets":             struct{}{},
		"persistentvolumes":        struct{}{},
		"persistentvolumeclaims":   struct{}{},
		"namespaces":               struct{}{},
		"horizontalpodautoscalers": struct{}{},
		"endpoints":                struct{}{},
		"secrets":                  struct{}{},
		"configmaps":               struct{}{},
	}

	DefaultAllowedMetrics = MetricList{
		"kube_configmap_created",
		"kube_configmap_info",
		"kube_configmap_metadata",
		"kube_cronjob_created",
		"kube_cronjob_info",
		"kube_cronjob_labels",
		"kube_cronjob_next",
		"kube_cronjob_spec",
		"kube_cronjob_status",
		"kube_daemonset_created",
		"kube_daemonset_labels",
		"kube_daemonset_metadata",
		"kube_daemonset_status",
		"kube_daemonset_updated",
		"kube_deployment_created",
		"kube_deployment_labels",
		"kube_deployment_metadata",
		"kube_deployment_spec",
		"kube_deployment_status",
		"kube_endpoint_address",
		"kube_endpoint_created",
		"kube_endpoint_info",
		"kube_endpoint_labels",
		"kube_hpa_labels",
		"kube_hpa_metadata",
		"kube_hpa_spec",
		"kube_hpa_status",
		"kube_job_complete",
		"kube_job_created",
		"kube_job_failed",
		"kube_job_info",
		"kube_job_labels",
		"kube_job_spec",
		"kube_job_status",
		"kube_limitrange_created",
		"kube_namespace_annotations",
		"kube_namespace_created",
		"kube_namespace_labels",
		"kube_namespace_status",
		"kube_node_created",
		"kube_node_info",
		"kube_node_labels",
		"kube_node_spec",
		"kube_node_status",
		"kube_persistentvolumeclaim_info",
		"kube_persistentvolumeclaim_labels",
		"kube_persistentvolumeclaim_resource",
		"kube_persistentvolumeclaim_status",
		"kube_persistentvolume_info",
		"kube_persistentvolume_labels",
		"kube_persistentvolume_status",
		"kube_pod_completion",
		"kube_pod_container",
		"kube_pod_created",
		"kube_pod_info",
		"kube_pod_labels",
		"kube_pod_owner",
		"kube_pod_spec",
		"kube_pod_start",
		"kube_pod_status",
		"kube_replicaset_created",
		"kube_replicaset_metadata",
		"kube_replicaset_spec",
		"kube_replicaset_status",
		"kube_replicationcontroller_created",
		"kube_replicationcontroller_metadata",
		"kube_replicationcontroller_spec",
		"kube_replicationcontroller_status",
		"kube_resourcequota_created",
		"kube_secret_created",
		"kube_secret_info",
		"kube_secret_labels",
		"kube_secret_metadata",
		"kube_secret_type",
		"kube_service_created",
		"kube_service_info",
		"kube_service_labels",
		"kube_service_spec",
		"kube_statefulset_created",
		"kube_statefulset_labels",
		"kube_statefulset_metadata",
		"kube_statefulset_replicas",
		"kube_statefulset_status",
	}
)
