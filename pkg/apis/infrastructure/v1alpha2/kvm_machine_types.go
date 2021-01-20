package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/serialization"
)

const (
	KindKVMMachine = "KVMMachine"
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
	// A cloud provider ID identifying the machine.
	ProviderID string `json:"providerID"`
	// Sizing information about the machine.
	Size KVMMachineSpecSize `json:"node"`
}

// +k8s:openapi-gen=true
type KVMMachineSpecSize struct {
	CPUs               int                 `json:"cpus"`
	Disk               serialization.Float `json:"disk"`
	Memory             string              `json:"memory"`
	DockerVolumeSizeGB int                 `json:"dockerVolumeSizeGB"`
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
