package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewClusterNetworkConfigCRD returns a new custom resource definition for ClusterNetworkConfig.
// This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: clusternetworkconfigs.core.giantswarm.io
//     spec:
//       group: core.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: ClusterNetworkConfig
//         plural: clusternetworkconfigs
//         singular: clusternetworkconfig
//       # subresources describes the subresources for custom resource.
//       subresources:
//          # status enables the status subresource.
//         status: {}
//
func NewClusterNetworkConfigCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "clusternetworkconfigs.core.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "core.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "ClusterNetworkConfig",
				Plural:   "clusternetworkconfigs",
				Singular: "clusternetworkconfig",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterNetworkConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ClusterNetworkConfigSpec   `json:"spec" yaml:"spec"`
	Status            ClusterNetworkConfigStatus `json:"status" yaml:"status"`
}

type ClusterNetworkConfigSpec struct {
	Cluster       ClusterNetworkConfigSpecCluster       `json:"cluster" yaml:"cluster"`
	VersionBundle ClusterNetworkConfigSpecVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type ClusterNetworkConfigSpecCluster struct {
	// ID contains Cluster ID of the guest cluster this ClusterNetworkConfig is
	// created for.
	ID      string                                 `json:"id" yaml:"id"`
	Network ClusterNetworkConfigSpecClusterNetwork `json:"network" yaml:"network"`
}

type ClusterNetworkConfigSpecClusterNetwork struct {
	// MaskBits is the number of ones in network mask that defines the
	// requested guest cluster network size. E.g. MaskBits:24 requests /24
	// network that can contain at max. 254 hosts minus environment specific
	// restrictions.
	MaskBits int `json:"maskBits" yaml:"maskBits"`
}

type ClusterNetworkConfigSpecVersionBundle struct {
	// Version contains version bundle version for cluster-operator network
	// controller implementation.
	Version string `json:"version" yaml:"version"`
}

type ClusterNetworkConfigStatus struct {
	// IP contains the network IP for allocated guest cluster subnet.
	IP string `json:"ip" yaml:"ip"`
	// Mask contains the network mask for allocated guest cluster subnet.
	Mask string `json:"mask" yaml:"mask"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterNetworkConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ClusterNetworkConfig `json:"items"`
}
