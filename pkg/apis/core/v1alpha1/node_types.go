package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindNodeConfig = "NodeConfig"
)

func NewNodeConfigCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindNodeConfig)
}

func NewNodeTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindNodeConfig,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NodeConfigSpec   `json:"spec"`
	Status            NodeConfigStatus `json:"status"`
}

type NodeConfigSpec struct {
}

type NodeConfigStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []NodeConfig `json:"items"`
}
