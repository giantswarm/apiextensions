package v1alpha1

import (
	"fmt"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindCluster = "Cluster"
)

// NewClusterCRD returns a new custom resource definition for Cluster. This
// might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: clusters.core.giantswarm.io
//     spec:
//       group: core.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: Cluster
//         plural: clusters
//         singular: cluster
//       subresources:
//         status: {}
//
func NewClusterCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("clusters.%s", group),
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   group,
			Scope:   "Namespaced",
			Version: version,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     kindCluster,
				Plural:   "clusters",
				Singular: "cluster",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

func NewTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindCluster,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status"`
}

// ClusterSpec is the part of the interface available to users in order to
// request a tenant cluster creation by providing necessary configurations.
// Fields here are either mandatory or optional. Optional fields left blank will
// be filled with appropriate default values which are then propagated into the
// CR status.
type ClusterSpec struct {
	// Cluster holds cluster specific configuration.
	Cluster ClusterSpecCluster `json:"cluster" yaml:"cluster"`
	// Release holds release specific configuration.
	Release ClusterSpecRelease `json:"release" yaml:"release"`
}

// ClusterSpecCluster holds cluster specific configuration users provide.
type ClusterSpecCluster struct {
	// Description is the optional cluster description users can provide. If left
	// blank a cluster description will be generated. The cluster description is
	// propagated into the CR status.
	Description string `json:"description" yaml:"description"`
	// Organization is the mandatory cluster organization in which a tenant
	// cluster will be scoped into.
	Organization string `json:"organization" yaml:"organization"`
}

// ClusterSpecRelease holds release specific configuration.
type ClusterSpecRelease struct {
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
	// Cluster holds cluster specific status information.
	Cluster ClusterStatusCluster `json:"cluster" yaml:"cluster"`
	// Conditions is a list of status conditions.
	Conditions []ClusterStatusCondition `json:"conditions" yaml:"conditions"`
	// Release holds release specific status information.
	Release ClusterStatusRelease `json:"release" yaml:"release"`
}

// ClusterStatusCluster holds cluster specific status information.
type ClusterStatusCluster struct {
	// Description is the propagated cluster description users can provide or the
	// system generates automatically if left blank.
	Description string `json:"description" yaml:"description"`
}

// ClusterStatusCondition holds a specific status condition describing certain
// aspects of the current state of the tenant cluster.
type ClusterStatusCondition struct {
	// Status may be True, False or Unknown.
	Status string `json:"status" yaml:"status"`
	// Type may be Creating, Created, Updating, Updated, or Deleting.
	Type string `json:"type" yaml:"type"`
}

// ClusterStatusRelease holds release specific status information.
type ClusterStatusRelease struct {
	// Version is the propagated release version users can provide or the system
	// sets to the current default release version.
	Version string `json:"version" yaml:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Cluster `json:"items"`
}
