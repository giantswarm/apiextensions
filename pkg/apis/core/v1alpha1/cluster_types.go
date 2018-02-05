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
	APIEndpoint       string                     `json:"apiEndpoint"`
	ID                string                     `json:"id"`
	KubernetesVersion string                     `json:"kubernetesVersion,omitempty"`
	Masters           []ClusterMasterSpec        `json:"masters,omitempty"`
	Name              string                     `json:"name,omitempty"`
	Owner             string                     `json:"owner,omitempty"`
	VersionBundles    []ClusterVersionBundleSpec `json:"versionBundles,omitempty"`
	Workers           []ClusterWorkerSpec        `json:"workers,omitempty"`
}

type ClusterMasterSpec struct {
	ClusterNodeSpec
}

type ClusterWorkerSpec struct {
	ClusterNodeSpec
	Labels map[string]string `json:"labels"`
}

type ClusterNodeSpec struct {
	InstanceType  string  `json:"instanceType,omitempty"`
	CPUCores      int     `json:"cpuCores,omitempty"`
	ID            string  `json:"id"`
	MemorySizeGB  float64 `json:"memorySizeGB,omitempty"`
	StorageSizeGB float64 `json:"storageSizeGB,omitempty"`
}

type ClusterVersionBundleSpec struct {
	Components   []ClusterComponentVersionSpec `json:"components,omitempty"`
	Dependencies []ClusterComponentVersionSpec `json:"dependencies,omitempty"`
	Name         string                        `json:"name"`
	Version      string                        `json:"version"`
	WIP          bool                          `json:"wip"`
}

type ClusterComponentVersionSpec struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ClusterConfig `json:"items"`
}
