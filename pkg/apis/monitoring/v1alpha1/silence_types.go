package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

const (
	kindSilence = "Silence"
)

func NewSilenceCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindSilence)
}

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster,categories=common;giantswarm
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
// Silence represents schema for managed silences in Alertmanager. Reconciled by silence-operator.
type Silence struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              SilenceSpec `json:"spec"`
}

// +k8s:openapi-gen=true
type SilenceSpec struct {
	TargetTags []TargetTag `json:"targetTags"`
	Matchers   []Matcher   `json:"matchers"`
}

type TargetTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Matcher struct {
	IsRegex bool   `json:"isRegex"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SilenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Silence `json:"items"`
}
