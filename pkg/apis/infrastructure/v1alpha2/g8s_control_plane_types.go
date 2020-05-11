package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/key"
)

// NewG8sControlPlaneCR returns a G8sControlPlane Custom Resource.
func NewG8sControlPlaneCR(name string) *G8sControlPlane {
	g8sControlPlane := G8sControlPlane{}
	groupVersionKind := metav1.GroupVersionKind{
		Group:   key.GroupApplication,
		Version: version,
		Kind:    key.KindApp,
	}
	g8sControlPlane.TypeMeta = key.NewTypeMeta(groupVersionKind)
	g8sControlPlane.ObjectMeta = key.NewObjectMeta(groupVersionKind)
	g8sControlPlane.Name = name
	return &g8sControlPlane
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;giantswarm;cluster-api

// The G8sControlPlane resource defines the Control Plane nodes (Kubernetes master nodes) of
// a Giant Swarm tenant cluster. It is reconciled by cluster-operator.
type G8sControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification part.
	Spec G8sControlPlaneSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Status information.
	Status G8sControlPlaneStatus `json:"status"`
}

type G8sControlPlaneSpec struct {
	// +kubebuilder:validation:Enum=1;3
	// +kubebuilder:validation:Optional
	// Number of master nodes.
	Replicas int `json:"replicas,omitempty"`
	// Reference to a provider-specific resource. On AWS, this would be of kind
	// [AWSControlPlane](https://docs.giantswarm.io/reference/cp-k8s-api/awscontrolplanes.infrastructure.giantswarm.io/).
	InfrastructureRef corev1.ObjectReference `json:"infrastructureRef"`
}

// G8sControlPlaneStatus defines the observed state of G8sControlPlane.
type G8sControlPlaneStatus struct {
	// +kubebuilder:validation:Optional
	// Total number of non-terminated machines targeted by this control plane
	// (their labels match the selector).
	Replicas int32 `json:"replicas,omitempty"`
	// +kubebuilder:validation:Optional
	// Total number of fully running and ready control plane machines.
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type G8sControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []G8sControlPlane `json:"items"`
}
