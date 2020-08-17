package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v2/pkg/crd"
)

const (
	kindOrganization = "Organization"
)

func NewOrganizationCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindOrganization)
}

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster,categories=common;giantswarm
// +kubebuilder:storageversion
// +k8s:openapi-gen=true

type Organization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              OrganizationSpec `json:"spec"`
}

// +k8s:openapi-gen=true
type OrganizationSpec struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Organization `json:"items"`
}
