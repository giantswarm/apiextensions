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
	// +kubebuilder:validation:Optional
	// +nullable
	Spec AzureServicePrincipalSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// +nullable
	Status AzureServicePrincipalStatus `json:"status"`
}

// +k8s:openapi-gen=true
type AzureServicePrincipalSpec struct {
	SecretRef corev1.ObjectReference `json:"secretRef"`
}

// +k8s:openapi-gen=true
type AzureServicePrincipalStatus struct {
	// +kubebuilder:validation:Optional
	ServicePrincipalName string `json:"servicePrincipalName,omitempty"`
	// +kubebuilder:validation:Optional
	InvitationLink string `json:"invitationLink,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	ExpirationDate *metav1.Time `json:"expirationDate,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	AccessConfirmed *AccessConfirmed `json:"accessConfirmed,omitempty"`
}

// +k8s:openapi-gen=true
type AccessConfirmed struct {
	// +kubebuilder:validation:Optional
	Confirmed *bool `json:"confirmed,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	LastCheckDate *metav1.Time `json:"lastCheckDate,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	LastSuccessDate *metav1.Time `json:"lastSuccessDate,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	LastFailureDate *metav1.Time `json:"lastFailureDate,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AzureServicePrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AzureServicePrincipal `json:"items"`
}
