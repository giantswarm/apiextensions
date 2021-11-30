package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster,categories=common;giantswarm,shortName=org;orgs
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +k8s:openapi-gen=true
// Organization represents schema for managed Kubernetes namespace. Reconciled by organization-operator.
type Organization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              OrganizationSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status OrganizationStatus `json:"status,omitempty"`
}

// +k8s:openapi-gen=true
type OrganizationSpec struct {
}

// +k8s:openapi-gen=true
type OrganizationStatus struct {
	// +kubebuilder:validation:Optional
	// Namespace is the namespace containing the resources for this organization.
	Namespace string `json:"namespace,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Organization `json:"items"`
}
