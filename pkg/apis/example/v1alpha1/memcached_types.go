package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/apis/example"
	"github.com/giantswarm/apiextensions/pkg/key"
)

// NewMemcachedConfigCR returns a MemcachedConfig custom resource.
func NewMemcachedConfigCR(name string) *MemcachedConfig {
	cr := MemcachedConfig{}
	cr.TypeMeta, cr.ObjectMeta = key.NewMeta(SchemeGroupVersion, example.KindMemcachedConfig, name, "")
	return &cr
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=example;giantswarm
// +kubebuilder:storageversion

type MemcachedConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              MemcachedConfigSpec `json:"spec"`
}

type MemcachedConfigSpec struct {
	// Replicas is the number of instances of Memcache.
	Replicas int `json:"replicas"`
	// e.g. 3
	// Memory is how much RAM to use for item storage.
	// e.g. 4G
	Memory string `json:"memory"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MemcachedConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []MemcachedConfig `json:"items"`
}
