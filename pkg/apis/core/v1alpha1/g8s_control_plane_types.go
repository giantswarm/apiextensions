package v1alpha1

import (
	"fmt"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewG8sControlPlaneCRD returns a new custom resource definition for
// G8sControlPlane. This might look something like the following.
//
//	apiVersion: apiextensions.k8s.io/v1beta1
//	kind: CustomResourceDefinition
//	metadata:
//	  name: g8sControlPlanes.core.giantswarm.io
//	spec:
//	  group: infrastructure.giantswarm.io
//	  scope: Namespaced
//	  version: v1alpha1
//	  names:
//	    kind: G8sControlPlane
//	    plural: g8sControlPlanes
//	    singular: g8sControlPlane
//    subresources:
//      status: {}
//
func NewG8sControlPlaneCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("g8sControlPlanes.%s", group),
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   group,
			Scope:   "Namespaced",
			Version: version,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "G8sControlPlane",
				Plural:   "g8sControlPlanes",
				Singular: "g8sControlPlane",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type G8sControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              G8sControlPlaneSpec   `json:"spec"`
	Status            G8sControlPlaneStatus `json:"status"`
}

type G8sControlPlaneSpec struct {
	// Replicas is the number replicas of the master node.
	Replicas int `json:"replicas" yaml:"replicas"`
	// InfrastructureRef is a required reference to provider-specific
	// Infrastructure.
	InfrastructureReference string `json:"infrastructureReference"`
}

// KubeadmControlPlaneStatus defines the observed state of G8sControlPlane.
type G8sControlPlaneStatus struct {
	// Selector is the label selector in string format to avoid introspection
	// by clients, and is used to provide the CRD-based integration for the
	// scale subresource and additional integrations for things like kubectl
	// describe.. The string will be in the same format as the query-param syntax.
	// More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors
	// +optional
	Selector string `json:"selector,omitempty"`

	// Total number of non-terminated machines targeted by this control plane
	// (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas,omitempty"`

	// Total number of non-terminated machines targeted by this control plane
	// that have the desired template spec.
	// +optional
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty"`

	// Total number of fully running and ready control plane machines.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// Total number of unavailable machines targeted by this control plane.
	// This is the total number of machines that are still required for
	// the deployment to have 100% available capacity. They may either
	// be machines that are running but not yet ready or machines
	// that still have not been created.
	// +optional
	UnavailableReplicas int32 `json:"unavailableReplicas,omitempty"`

	// Initialized denotes whether or not the control plane has the
	// uploaded configmap.
	// +optional
	Initialized bool `json:"initialized"`

	// Ready denotes that the G8sControlPlane API Server is ready to
	// receive requests.
	// +optional
	Ready bool `json:"ready"`

	// FailureReason indicates that there is a terminal problem reconciling the
	// state, and will be set to a token value suitable for
	// programmatic interpretation.
	// +optional
	FailureReason error `json:"failureReason,omitempty"`

	// ErrorMessage indicates that there is a terminal problem reconciling the
	// state, and will be set to a descriptive error message.
	// +optional
	FailureMessage *string `json:"failureMessage,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type G8sControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []G8sControlPlane `json:"items"`
}
