package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v6/pkg/annotation"
)

const (
	kindAWSControlPlane              = "AWSControlPlane"
	awsControlPlaneDocumentationLink = "https://docs.giantswarm.io/ui-api/management-api/crd/awscontrolplanes.infrastructure.giantswarm.io/"
)

func NewAWSControlPlaneTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAWSControlPlane,
	}
}

// NewAWSControlPlaneCR returns an AWSControlPlane Custom Resource.
func NewAWSControlPlaneCR() *AWSControlPlane {
	return &AWSControlPlane{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotation.Docs: awsControlPlaneDocumentationLink,
			},
		},
		TypeMeta: NewAWSControlPlaneTypeMeta(),
	}
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=aws;giantswarm
// +kubebuilder:storageversion
// +k8s:openapi-gen=true

// AWSControlPlane is the infrastructure provider referenced in ControlPlane
// CRs. Represents the master nodes (also called Control Plane) of a workload
// cluster on AWS. Reconciled by aws-operator.
type AWSControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification part of the resource.
	Spec AWSControlPlaneSpec `json:"spec"`
}

// +k8s:openapi-gen=true
type AWSControlPlaneSpec struct {
	// +kubebuilder:validation:Optional
	// Configures which AWS availability zones to use by master nodes, as a list
	// of availability zone names like e. g. `eu-central-1c`. We support either
	// 1 or 3 availability zones.
	AvailabilityZones []string `json:"availabilityZones,omitempty"`
	// +kubebuilder:validation:Optional
	// EC2 instance type identifier to use for the master node(s).
	InstanceType string `json:"instanceType,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AWSControlPlane `json:"items"`
}
