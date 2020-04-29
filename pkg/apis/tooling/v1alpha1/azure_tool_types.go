package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindAzureTool = "AzureTool"
)

func NewAzureToolCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindAzureTool)
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
type AzureTool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AzureToolSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status AzureToolStatus `json:"status,omitempty"`
}

type AzureToolSpec struct {
	// Workspace refers to the Azure Log Analytics Workspace.
	Workspace AzureToolWorkspace `json:"workspace" yaml:"workspace"`
}

type AzureToolStatus struct {
	// +kubebuilder:validation:Optional

	// WorkspaceStatus of the upstream Azure Log Analytics Workspace deployment (can be 'Pending', 'Active').
	WorkspaceStatus string `json:"workspace_status" yaml:"workspace_status"`
}

type AzureToolWorkspace struct {
	// ID is the Workspace ID.
	ID string `json:"id" yaml:"id"`
	// Mode is the mode that the Workspace is running in.
	Mode string `json:"mode" yaml:"mode"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureToolList is the type returned when listing AzureToolList resources.
type AzureToolList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []AzureTool `json:"items" yaml:"items"`
}
