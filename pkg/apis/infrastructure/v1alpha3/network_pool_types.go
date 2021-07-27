package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;giantswarm;cluster-api
// +k8s:openapi-gen=true

// NetworkPool is the infrastructure provider referenced in upstream CAPI Cluster
// CRs.
type NetworkPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkPoolSpec `json:"spec"`
}

// NetworkPoolSpec is the spec part for the NetworkPool resource.
// +k8s:openapi-gen=true
type NetworkPoolSpec struct {
	// IPv4 address block in CIDR notation.
	CIDRBlock string `json:"cidrBlock,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NetworkPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []NetworkPool `json:"items"`
}
