package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/apis/core"
	"github.com/giantswarm/apiextensions/pkg/crd"
	"github.com/giantswarm/apiextensions/pkg/key"
)

// NewStorageConfigCRD returns a CRD defining a StorageConfig.
func NewStorageConfigCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(core.Group, core.KindStorageConfig)
}

// NewStorageConfigCR returns a StorageConfig custom resource.
func NewStorageConfigCR(name, namespace string) *StorageConfig {
	cr := StorageConfig{}
	cr.TypeMeta, cr.ObjectMeta = key.NewMeta(SchemeGroupVersion, core.KindStorageConfig, name, namespace)
	return &cr
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=common;giantswarm
// +kubebuilder:storageversion
// +k8s:openapi-gen=true

type StorageConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              StorageConfigSpec `json:"spec"`
}

// +k8s:openapi-gen=true
type StorageConfigSpec struct {
	Storage StorageConfigSpecStorage `json:"storage"`
}

// +k8s:openapi-gen=true
type StorageConfigSpecStorage struct {
	Data map[string]string `json:"data"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StorageConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []StorageConfig `json:"items"`
}
