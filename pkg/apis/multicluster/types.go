/*
Copyright 2019 The Kubernetes Authors.

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

package multicluster

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "k8s.io/kubernetes/pkg/apis/core"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceExport declares that the associated service should be exported to
// other clusters.
type ServiceExport struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta
	// +optional
	Status ServiceExportStatus
}

// ServiceExportStatus contains the current status of an export.
type ServiceExportStatus struct {
	// +optional
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +listType=map
	// +listMapKey=type
	Conditions []ServiceExportCondition
}

// ServiceExportConditionType identifies a specific condition.
type ServiceExportConditionType string

const (
	// ServiceExportInitialized means the service export has been noticed
	// by the controller, has passed validation, has appropriate finalizers
	// set, and any required supercluster resources like the IP have been
	// reserved
	ServiceExportInitialized ServiceExportConditionType = "Initialized"
	// ServiceExportExported means that the service referenced by this
	// service export has been synced to all clusters in the supercluster
	ServiceExportExported ServiceExportConditionType = "Exported"
)

// ServiceExportCondition contains details for the current condition of this
// service export.
//
// Once [#1624](https://github.com/kubernetes/enhancements/pull/1624) is
// merged, this will be replaced by metav1.Condition.
type ServiceExportCondition struct {
	Type ServiceExportConditionType
	// Status is one of {"True", "False", "Unknown"}
	Status api.ConditionStatus
	// +optional
	LastTransitionTime *metav1.Time
	// +optional
	Reason *string
	// +optional
	Message *string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceImport declares that the specified service should be exported to other clusters.
type ServiceImport struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta
	// +optional
	Spec ServiceImportSpec
}

// ServiceImportSpec contains the current status of an imported service and the
// information necessary to consume it
type ServiceImportSpec struct {
	// +patchMergeKey=port
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=port
	// +listMapKey=protocol
	Ports []api.ServicePort
	// +optional
	// +patchStrategy=merge
	// +patchMergeKey=cluster
	// +listType=map
	// +listMapKey=cluster
	Clusters []ClusterSpec
	// +optional
	IPFamily *api.IPFamily
	// +optional
	IP string
}

// ClusterSpec contains service configuration mapped to a specific cluster
type ClusterSpec struct {
	Cluster string
	// +listType=atomic
	// +optional
	TopologyKeys []string
	// +optional
	PublishNotReadyAddresses bool
	// +optional
	SessionAffinity api.ServiceAffinity
	// +optional
	SessionAffinityConfig *api.SessionAffinityConfig
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceExportList represents a list of endpoint slices
type ServiceExportList struct {
	metav1.TypeMeta
	// Standard list metadata.
	// +optional
	metav1.ListMeta
	// List of endpoint slices
	// +listType=set
	Items []ServiceExport
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceImportList represents a list of endpoint slices
type ServiceImportList struct {
	metav1.TypeMeta
	// Standard list metadata.
	// +optional
	metav1.ListMeta
	// List of endpoint slices
	// +listType=set
	Items []ServiceImport
}
