package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewChaosMonkeyConfigCRD returns a new custom resource definition for
// ChaosMonkeyConfig. This might look something like the following.
//
//	apiVersion: apiextensions.k8s.io/v1beta1
//	kind: CustomResourceDefinition
//	metadata:
//	  name: chaosmonkeyconfigs.example.giantswarm.io
//	spec:
//	  group: example.giantswarm.io
//	  scope: Namespaced
//	  version: v1alpha1
//	  names:
//	    kind: ChaosMonkeyConfig
//	    plural: chaosmonkeyconfigs
//	    singular: chaosmonkeyconfig
//
func NewChaosMonkeyConfigCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "chaosmonkeyconfigs.example.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "example.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "ChaosMonkeyConfig",
				Plural:   "chaosmonkeyconfigs",
				Singular: "chaosmonkeyconfig",
			},
		},
	}
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChaosMonkeyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ChaosMonkeyConfigSpec `json:"spec"`
	Status            ChaosMonkeyStatus     `json:"status"`
}

type ChaosMonkeyConfigSpec struct {
	// StartTime is the beginning of the period when pods can be killed.
	StartTime DeepCopyTime `json:"startTime" yaml:"startTime"`
	// EndTime is the end of the period when pods can be killed.
	EndTime DeepCopyTime `json:"endTime" yaml:"endTime"`
	// DryRun logs what actions would have been taken.
	DryRun bool `json:"dryRun" yaml:"dryRun"`
	// NamespaceBlacklist is a list of namespaces to ignore
	NamespaceBlacklist []string `json:"namespaceBlacklist" yaml:"namespaceBlacklist"`
}

type ChaosMonkeyStatus struct {
	LastUpdateTime DeepCopyTime `json:"lastUpdateTime" yaml:"lastUpdateTime"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChaosMonkeyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ChaosMonkeyConfig `json:"items"`
}
