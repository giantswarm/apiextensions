package v1alpha1

import (
	"fmt"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindIgnition = "Ignition"
)

// NewIgnitionCRD returns a new custom resource definition for Ignition. This
// might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: ignitions.core.giantswarm.io
//     spec:
//       group: core.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: Ignition
//         plural: ignitions
//         singular: ignition
//       subresources:
//         status: {}
//
func NewIgnitionCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("ignitions.%s", group),
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   group,
			Scope:   "Namespaced",
			Version: version,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     kindIgnition,
				Plural:   "ignitions",
				Singular: "ignition",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

func NewIgnitionTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindIgnition,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Ignition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              IgnitionSpec   `json:"spec"`
	Status            IgnitionStatus `json:"status"`
}

// IgnitionSpec is the interface which defines the input parameters for
// a newly rendered g8s ignition template.
type IgnitionSpec struct {
	// Defines the provider which should be rendered.
	Provider string `json:"provider" yaml:"provider"`
}

// IgnitionStatus holds the rendering result.
type IgnitionStatus struct {
	ConfigMap IgnitionStatusConfigMap `json:"template" yaml:"template"`
}

type IgnitionStatusConfigMap struct {
	// Name is the name of the config map containing the rendered ignition.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the config map containing the rendered ignition.
	Namespace string `json:"namespace" yaml:"namespace"`
	// ResourceVersion is the Kubernetes resource version of the configmap.
	// Used to detect if the configmap has changed, e.g. 12345.
	ResourceVersion string `json:"resourceVersion" yaml:"resourceVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IgnitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Ignition `json:"items"`
}
