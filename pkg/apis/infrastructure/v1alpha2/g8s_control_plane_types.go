package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindG8sControlPlane              = "G8sControlPlane"
	g8sControlPlaneDocumentationLink = "https://pkg.go.dev/github.com/giantswarm/apiextensions@v0.2.5/pkg/apis/infrastructure/v1alpha2?tab=doc#G8sControlPlane"
)

func NewG8sControlPlaneCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindG8sControlPlane)
}

func NewG8sControlPlaneTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindG8sControlPlane,
	}
}

// NewG8sControlPlaneCR returns a G8sControlPlane Custom Resource.
func NewG8sControlPlaneCR() *G8sControlPlane {
	return &G8sControlPlane{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: g8sControlPlaneDocumentationLink,
			},
		},
		TypeMeta: NewClusterTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status

// G8sControlPlane defines the Control Plane Nodes (Kubernetes Master Nodes) of
// a Giant Swarm Tenant Cluster.
type G8sControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              G8sControlPlaneSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status G8sControlPlaneStatus `json:"status"`
}

type G8sControlPlaneSpec struct {
	// Replicas is the number replicas of the master node.
	Replicas int `json:"replicas" yaml:"replicas"`
	// InfrastructureRef is a required reference to provider-specific
	// Infrastructure.
	InfrastructureRef corev1.ObjectReference `json:"infrastructureRef"`
}

// G8sControlPlaneStatus defines the observed state of G8sControlPlane.
type G8sControlPlaneStatus struct {
	// +kubebuilder:validation:Enum=1;3
	// Total number of non-terminated machines targeted by this control plane
	// (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// Total number of fully running and ready control plane machines.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type G8sControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []G8sControlPlane `json:"items"`
}
