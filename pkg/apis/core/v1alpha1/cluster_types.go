package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewClusterConfigCRD returns a new custom resource definition for
// ClusterConfig. This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: clusterconfigs.core.giantswarm.io
//     spec:
//       group: core.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: ClusterConfig
//         plural: clusterconfigs
//         singular: clusterconfig
//
func NewClusterConfigCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "clusterconfigs.core.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "core.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "ClusterConfig",
				Plural:   "clusterconfigs",
				Singular: "clusterconfig",
			},
		},
	}
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ClusterConfigSpec `json:"spec"`
}

type ClusterConfigSpec struct {
	Guest ClusterConfigSpecGuest `json:"guest" yaml:"guest"`
}

type ClusterConfigSpecGuest struct {
	API            ClusterConfigSpecGuestAPI        `json:"api" yaml:"api"`
	ID             string                           `json:"id" yaml:"id"`
	Masters        []ClusterConfigSpecGuestMaster   `json:"masters,omitempty" yaml:"masters,omitempty"`
	Name           string                           `json:"name,omitempty" yaml:"name,omitempty"`
	Owner          string                           `json:"owner,omitempty" yaml:"owner,omitempty"`
	VersionBundles []ClusterConfigSpecVersionBundle `json:"versionBundles,omitempty" yaml:"versionBundles,omitempty"`
	Workers        []ClusterConfigSpecGuestWorker   `json:"workers,omitempty" yaml:"workers,omitempty"`
}

type ClusterConfigSpecGuestAPI struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

type ClusterConfigSpecGuestMaster struct {
	ClusterConfigSpecGuestNode
}

type ClusterConfigSpecGuestWorker struct {
	ClusterConfigSpecGuestNode
	Labels map[string]string `json:"labels" yaml:"labels"`
}

type ClusterConfigSpecGuestNode struct {
	AWS   ClusterConfigSpecAWS   `json:"aws,omitempty" yaml:"aws,omitempty"`
	Azure ClusterConfigSpecAzure `json:"azure,omitempty" yaml:"azure,omitempty"`
	ID    string                 `json:"id" yaml:"id"`
	KVM   ClusterConfigSpecKVM   `json:"kvm,omitempty" yaml:"kvm,omitempty"`
}

type ClusterConfigSpecAWS struct {
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
}

type ClusterConfigSpecAzure struct {
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
	VMSize       string `json:"vmSize,omitempty" yaml:"vmSize,omitempty"`
}

type ClusterConfigSpecKVM struct {
	CPUCores      int     `json:"cpuCores,omitempty" yaml:"cpuCores,omitempty"`
	MemorySizeGB  float64 `json:"memorySizeGB,omitempty" yaml:"memorySizeGB,omitempty"`
	StorageSizeGB float64 `json:"storageSizeGB,omitempty" yaml:"storageSizeGB,omitempty"`
}

type ClusterConfigSpecVersionBundle struct {
	Components   []ClusterConfigSpecComponentVersion `json:"components,omitempty" yaml:"components,omitempty"`
	Dependencies []ClusterConfigSpecComponentVersion `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Name         string                              `json:"name" yaml:"name"`
	Version      string                              `json:"version" yaml:"version"`
	WIP          bool                                `json:"wip" yaml:"wip"`
}

type ClusterConfigSpecComponentVersion struct {
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ClusterConfig `json:"items"`
}
