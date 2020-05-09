package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindStorageConfig = "StorageConfig"
)

// NewStorageConfigCRD returns a new custom resource definition for StorageConfig.
func NewStorageConfigCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindStorageConfig)
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=giantswarm;common

type StorageConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              StorageConfigSpec `json:"spec"`
}

type StorageConfigSpec struct {
	Storage StorageConfigSpecStorage `json:"storage"`
}

type StorageConfigSpecStorage struct {
	Data map[string]string `json:"data"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StorageConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []StorageConfig `json:"items"`
}
