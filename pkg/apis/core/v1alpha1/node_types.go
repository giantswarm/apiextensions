package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewNodeConfigCRD returns a new custom resource definition for NodeConfig.
// This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: nodeconfigs.core.giantswarm.io
//     spec:
//       group: core.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: NodeConfig
//         plural: nodeconfigs
//         singular: nodeconfig
//
func NewNodeConfigCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "nodeconfigs.core.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "core.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "NodeConfig",
				Plural:   "nodeconfigs",
				Singular: "nodeconfig",
			},
		},
	}
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NodeConfigSpec   `json:"spec"`
	Status            NodeConfigStatus `json:"status"`
}

type NodeConfigSpec struct {
	Patch NodeConfigSpecPatch `json:"patch" yaml:"patch"`
}

type NodeConfigSpecPatch struct {
	// SelfLink is the self link of the provider config CR. This is used by the
	// node-operator to patch the status of the provider config CR. Without the
	// self link the node-operator would not know which runtime object to patch.
	SelfLink string `json:"selfLink" yaml:"selfLink"`
}

type NodeConfigStatus struct {
	Conditions []NodeConfigStatusCondition `json:"conditions" yaml:"conditions"`
}

// NodeConfigStatusCondition expresses a condition in which a node may is.
type NodeConfigStatusCondition struct {
	Status string `json:"status" yaml:"status"`
	Type   string `json:"type" yaml:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []NodeConfig `json:"items"`
}
