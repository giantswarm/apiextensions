package v1alpha2

import (
	v1 "k8s.io/api/core/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewG8sBootstrapConfigConfigCRD returns a new G8sBootstrapConfigConfigCRD
func NewG8sBootstrapConfigConfigCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "g8sbootstrapconfigs.bootstrap.cluster.x-k8s.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "bootstrap.cluster.x-k8s.io",
			Scope:   "Namespaced",
			Version: "v1alpha2",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "G8sBootstrapConfigConfig",
				Plural:   "G8sBootstrapConfigconfigs",
				Singular: "G8sBootstrapConfigconfig",
			},
		},
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type G8sBootstrapConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              G8sBootstrapSpec   `json:"spec"`
	Status            G8sBootstrapStatus `json:"status"`
}

type G8sBootstrapSpec struct {
	G8sVersion string `json:"g8sVersion,omitempty"`
}

// G8sBootstrapStatus

type G8sBootstrapStatus struct {

	// a boolean field indicating the bootstrap config data is ready for use.
	Ready bool `json:"ready,omitempty"`

	//  A string field containing some data used for bootstrapping a cluster.
	BootstrapData []byte `json:"bootstrapData,omitempty"`

	// a string that explains why an error has occurred, if possible.
	ErrorReason string `json:"errorReason,omitempty"`

	// a string that holds the message contained by the error.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// A slice of addresses ([]v1.NodeAddress) that contains a list of apiserver endpoints.
	Addresses []v1.NodeAddress `json:"addresses,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type G8sBootstrapConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []G8sBootstrapConfig `json:"items"`
}
