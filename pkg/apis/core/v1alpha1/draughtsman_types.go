package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

func NewDraughtsmanConfigCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindCluster)
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DraughtsmanConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              DraughtsmanConfigSpec `json:"spec"`
}

type DraughtsmanConfigSpec struct {
	Projects []DraughtsmanConfigSpecProject `json:"projects" yaml:"projects"`
}

type DraughtsmanConfigSpecProject struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	Ref  string `json:"ref" yaml:"ref"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DraughtsmanConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []DraughtsmanConfig `json:"items"`
}
