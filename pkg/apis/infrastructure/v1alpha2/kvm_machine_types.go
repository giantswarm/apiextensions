package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindKVMMachine = "KVMMachine"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=giantswarm;kvm
// +k8s:openapi-gen=true

type KVMMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KVMMachineSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status KVMMachineStatus `json:"status"`
}

// +k8s:openapi-gen=true
type KVMMachineSpec struct {
	ProviderID string `json:"providerID"`
}

// +k8s:openapi-gen=true
type KVMMachineStatus struct {
	Ready          bool   `json:"ready"`
	FailureReason  string `json:"failureReason"`
	FailureMessage string `json:"failureMessage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KVMMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KVMMachine `json:"items"`
}
