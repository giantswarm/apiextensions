package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

const (
	kindAzureServicePrincipal = "AzureServicePrincipal"
)

func NewAzureServicePrincipalCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindAzureServicePrincipal)
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=azure;giantswarm
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
// AzureServicePrincipal represents schema for Azure Credentials needed to talk with Azure API. Reconciled by azure-credentials-operator.
type AzureServicePrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AzureServicePrincipalSpec   `json:"spec"`
	Status            AzureServicePrincipalStatus `json:"status"`
}

// +k8s:openapi-gen=true
type AzureServicePrincipalSpec struct {
	Name      *string                 `json:"name"`
	SecretRef *corev1.ObjectReference `json:"secretRef"`
}

// +k8s:openapi-gen=true
type AzureServicePrincipalStatus struct {
	InvitationLink  *string          `json:"invitationLink,omitempty"`
	ExpirationDate  *metav1.Time     `json:"expirationDate,omitempty"`
	AccessConfirmed *AccessConfirmed `json:"accessConfirmed,omitempty"`
}

// +k8s:openapi-gen=true
type AccessConfirmed struct {
	Confirmed       *bool        `json:"confirmed,omitempty"`
	LastCheckDate   *metav1.Time `json:"lastCheckDate,omitempty"`
	LastSuccessDate *metav1.Time `json:"lastSuccessDate,omitempty"`
	LastFailureDate *metav1.Time `json:"lastFailureDate,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AzureServicePrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AzureServicePrincipal `json:"items"`
}
