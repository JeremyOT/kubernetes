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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceExport declares that the associated service should be exported to
// other clusters.
type ServiceExport struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// +optional
	Status ServiceExportStatus `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`
}

// ServiceExportStatus contains the current status of an export.
type ServiceExportStatus struct {
	// +optional
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +listType=map
	// +listMapKey=type
	Conditions []ServiceExportCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
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
	Type ServiceExportConditionType `json:"type" protobuf:"bytes,1,opt,name=type"`
	// Status is one of {"True", "False", "Unknown"}
	Status v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status"`
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	// +optional
	Reason *string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	// +optional
	Message *string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceImport declares that the specified service should be exported to other clusters.
type ServiceImport struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// +optional
	Spec ServiceImportSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// ServiceImportSpec contains the current status of an imported service and the
// information necessary to consume it
type ServiceImportSpec struct {
	// +patchMergeKey=port
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=port
	// +listMapKey=protocol
	Ports []v1.ServicePort `json:"ports" patchStrategy:"merge" patchMergeKey:"port" protobuf:"bytes,1,rep,name=ports"`
	// +optional
	// +patchStrategy=merge
	// +patchMergeKey=cluster
	// +listType=map
	// +listMapKey=cluster
	Clusters []ClusterSpec `json:"clusters" patchStrategy:"merge" patchMergeKey:"cluster" protobuf:"bytes,2,rep,name=clusters"`
	// +optional
	IPFamily *v1.IPFamily `json:"ipFamily" protobuf:"bytes,3,opt,name=ipFamily"`
	// +optional
	IP string `json:"ip,omitempty" protobuf:"bytes,4,opt,name=IP"`
}

// ClusterSpec contains service configuration mapped to a specific cluster
type ClusterSpec struct {
	Cluster string `json:"cluster" protobuf:"bytes,1,opt,name=cluster"`
	// +listType=atomic
	// +optional
	TopologyKeys []string `json:"topologyKeys" protobuf:"bytes,2,rep,name=topologyKeys"`
	// +optional
	PublishNotReadyAddresses bool `json:"publishNotReadyAddresses" protobuf:"varint,3,opt,name=publishNotReadyAddresses"`
	// +optional
	SessionAffinity v1.ServiceAffinity `json:"sessionAffinity" protobuf:"bytes,4,opt,name=sessionAffinity"`
	// +optional
	SessionAffinityConfig *v1.SessionAffinityConfig `json:"sessionAffinityConfig" protobuf:"bytes,5,opt,name=sessionAffinityConfig"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceExportList represents a list of endpoint slices
type ServiceExportList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// List of endpoint slices
	// +listType=set
	Items []ServiceExport `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceImportList represents a list of endpoint slices
type ServiceImportList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// List of endpoint slices
	// +listType=set
	Items []ServiceImport `json:"items" protobuf:"bytes,2,rep,name=items"`
}
