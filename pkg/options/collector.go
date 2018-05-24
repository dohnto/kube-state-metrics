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
		"kube_configmap_created": true,
		"kube_configmap_info": true,
		"kube_configmap_metadata": true,
		"kube_cronjob_created": true,
		"kube_cronjob_info": true,
		"kube_cronjob_labels": true,
		"kube_cronjob_next": true,
		"kube_cronjob_spec": true,
		"kube_cronjob_status": true,
		"kube_daemonset_created": true,
		"kube_daemonset_labels": true,
		"kube_daemonset_metadata": true,
		"kube_daemonset_status": true,
		"kube_daemonset_updated": true,
		"kube_deployment_created": true,
		"kube_deployment_labels": true,
		"kube_deployment_metadata": true,
		"kube_deployment_spec": true,
		"kube_deployment_status": true,
		"kube_endpoint_address": true,
		"kube_endpoint_created": true,
		"kube_endpoint_info": true,
		"kube_endpoint_labels": true,
		"kube_hpa_labels": true,
		"kube_hpa_metadata": true,
		"kube_hpa_spec": true,
		"kube_hpa_status": true,
		"kube_job_complete": true,
		"kube_job_created": true,
		"kube_job_failed": true,
		"kube_job_info": true,
		"kube_job_labels": true,
		"kube_job_spec": true,
		"kube_job_status": true,
		"kube_limitrange_created": true,
		"kube_namespace_annotations": true,
		"kube_namespace_created": true,
		"kube_namespace_labels": true,
		"kube_namespace_status": true,
		"kube_node_created": true,
		"kube_node_info": true,
		"kube_node_labels": true,
		"kube_node_spec": true,
		"kube_node_status": true,
		"kube_persistentvolumeclaim_info": true,
		"kube_persistentvolumeclaim_labels": true,
		"kube_persistentvolumeclaim_resource": true,
		"kube_persistentvolumeclaim_status": true,
		"kube_persistentvolume_info": true,
		"kube_persistentvolume_labels": true,
		"kube_persistentvolume_status": true,
		"kube_pod_completion": true,
		"kube_pod_container": true,
		"kube_pod_created": true,
		"kube_pod_info": true,
		"kube_pod_labels": true,
		"kube_pod_owner": true,
		"kube_pod_spec": true,
		"kube_pod_start": true,
		"kube_pod_status": true,
		"kube_replicaset_created": true,
		"kube_replicaset_metadata": true,
		"kube_replicaset_spec": true,
		"kube_replicaset_status": true,
		"kube_replicationcontroller_created": true,
		"kube_replicationcontroller_metadata": true,
		"kube_replicationcontroller_spec": true,
		"kube_replicationcontroller_status": true,
		"kube_resourcequota_created": true,
		"kube_secret_created": true,
		"kube_secret_info": true,
		"kube_secret_labels": true,
		"kube_secret_metadata": true,
		"kube_secret_type": true,
		"kube_service_created": true,
		"kube_service_info": true,
		"kube_service_labels": true,
		"kube_service_spec": true,
		"kube_statefulset_created": true,
		"kube_statefulset_labels": true,
		"kube_statefulset_metadata": true,
		"kube_statefulset_replicas": true,
		"kube_statefulset_status": true,
	}

	DefaultDisabledMetrics = MetricList{}
)
