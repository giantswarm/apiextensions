package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindReleaseCycle = "ReleaseCycle"

	CyclePhaseUpcoming ReleaseCyclePhase = "upcoming"
	CyclePhaseEnabled  ReleaseCyclePhase = "enabled"
	CyclePhaseDisabled ReleaseCyclePhase = "disabled"
	CyclePhaseEOL      ReleaseCyclePhase = "eol"
)

type ReleaseCyclePhase string

func (r ReleaseCyclePhase) String() string {
	return string(r)
}

func NewReleaseCycleTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindReleaseCycle,
	}
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=giantswarm;common

type ReleaseCycle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ReleaseCycleSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status ReleaseCycleStatus `json:"status,omitempty"`
}

type ReleaseCycleSpec struct {
	// DisabledDate is the date of the cycle phase being changed to "disabled".
	DisabledDate DeepCopyDate `json:"disabledDate,omitempty"`
	// EnabledDate is the date of the cycle phase being changed to "enabled".
	EnabledDate DeepCopyDate `json:"enabledDate,omitempty"`
	// Phase is the release phase. It can be one of: "upcoming", "enabled",
	// "disabled", "eol".
	Phase ReleaseCyclePhase `json:"phase"`
}

type ReleaseCycleStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseCycleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ReleaseCycle `json:"items"`
}
