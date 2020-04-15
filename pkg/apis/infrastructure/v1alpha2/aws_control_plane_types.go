package v1alpha2

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindAWSControlPlane = "AWSControlPlane"

	// TODO: change to "https://docs.giantswarm.io/reference/cp-k8s-api/awscontrolplanes.infrastructure.giantswarm.io/"
	// after this has been first published.
	awsControlPlaneDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/"
)

func NewAWSControlPlaneCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindAWSControlPlane)
}

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
				crDocsAnnotation: awsControlPlaneDocumentationLink,
			},
		},
		TypeMeta: NewAWSControlPlaneTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSControlPlane is the infrastructure provider referenced in ControlPlane
// CRs.
type AWSControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AWSControlPlaneSpec   `json:"spec" yaml:"spec"`
	Status            AWSControlPlaneStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type AWSControlPlaneSpec struct {
	AvailabilityZones []string `json:"availabilityZones" yaml:"availabilityZones"`
	InstanceType      string   `json:"instanceType" yaml:"instanceType"`
}

// TODO
type AWSControlPlaneStatus struct {
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSControlPlaneList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []AWSControlPlane `json:"items" yaml:"items"`
}
