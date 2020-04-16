package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindCluster = "Cluster"
)

func NewClusterCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindCluster)
}

func NewClusterTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindCluster,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ClusterSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status ClusterStatus `json:"status"`
}

// ClusterSpec is the part of the interface available to users in order to
// request a tenant cluster creation by providing necessary configurations.
// Fields here are either mandatory or optional. Optional fields left blank will
// be filled with appropriate default values which are then propagated into the
// CR status.
type ClusterSpec struct {
	// Description is the optional cluster description users can provide. If left
	// blank a cluster description will be generated. The cluster description is
	// propagated into the CR status.
	Description string `json:"description" yaml:"description"`
	// Organization is the mandatory cluster organization in which a tenant
	// cluster will be scoped into.
	Organization string `json:"organization" yaml:"organization"`
	// Version is the optional release version users can provide. If left blank
	// the current default release version will be used. The release version is
	// propagated into the CR status.
	Version string `json:"version" yaml:"version"`
}

// ClusterStatus is the part of the interface available to users in order to
// fetch a tenant cluster's status information after creation was requested.
// Fields here are automatically filled and can only ever be read. For instance
// the tenant cluster description will be generated if left blank upon cluster
// creation and made available here.
type ClusterStatus struct {
	// LastHeartbeatTime is the last time we got an update on a given condition.
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime" yaml:"lastHeartbeatTime"`
	// LastTransitionTime is the last time the condition transitioned from one
	// status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime" yaml:"lastTransitionTime"`
	// Cluster holds cluster specific status information.
	Cluster ClusterStatusCluster `json:"cluster" yaml:"cluster"`
	// Conditions is a list of status conditions.
	Conditions []ClusterStatusCondition `json:"conditions" yaml:"conditions"`
}

// ClusterStatusCluster holds cluster specific status information. Some of the
// fields from this structure may move back to the spec in the future once we
// make more use of mutating admission controllers for defaulting reasons. For
// instance the cluster ID and version are candidates for this.
type ClusterStatusCluster struct {
	// Description is the propagated cluster description users can provide or the
	// system generates automatically if left blank.
	Description string `json:"description" yaml:"description"`
	// ID is the internal cluster ID automatically generated upon cluster
	// creation.
	ID string `json:"id" yaml:"id"`
	// Version is the propagated release version users can provide or the system
	// sets to the current default release version.
	Version string `json:"version" yaml:"version"`
}

// ClusterStatusCondition holds a specific status condition describing certain
// aspects of the current state of the tenant cluster.
type ClusterStatusCondition struct {
	// Status may be True, False or Unknown.
	Status string `json:"status" yaml:"status"`
	// Type may be Creating, Created, Updating, Updated, or Deleting.
	Type string `json:"type" yaml:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Cluster `json:"items"`
}
