package v1alpha3

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindAzureControlPlane = "AzureControlPlane"

	// TODO: change to "https://docs.giantswarm.io/reference/cp-k8s-api/azurecontrolplanes.infrastructure.giantswarm.io/"
	// after this has been first published.
	azureControlPlaneDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/"
)

func NewAzureControlPlaneCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindAzureControlPlane)
}

func NewAzureControlPlaneTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAzureControlPlane,
	}
}

// NewAzureControlPlaneCR returns an AzureControlPlane Custom Resource.
func NewAzureControlPlaneCR() *AzureControlPlane {
	return &AzureControlPlane{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: azureControlPlaneDocumentationLink,
			},
		},
		TypeMeta: NewAzureControlPlaneTypeMeta(),
	}
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureControlPlane is the infrastructure provider referenced in ControlPlane
// CRs. Represents the master nodes (also called Control Plane) of a tenant
// cluster on Azure. Reconciled by azure-operator.
type AzureControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification part of the resource.
	Spec AzureControlPlaneSpec `json:"spec"`
}

type AzureControlPlaneSpec struct {
	// +kubebuilder:validation:Optional
	// Configures which Azure availability zones to use by master nodes, as a list
	// of availability zone names like e. g. `eu-central-1c`. We support either
	// 1 or 3 availability zones.
	AvailabilityZones []string `json:"availabilityZones,omitempty"`
	// EC2 instance type identifier to use for the master node(s).
	InstanceType string `json:"instanceType"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzureControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AzureControlPlane `json:"items"`
}
