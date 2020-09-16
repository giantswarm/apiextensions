package v1alpha2

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v2/pkg/annotation"
	"github.com/giantswarm/apiextensions/v2/pkg/crd"
)

const (
	kindAWSIPPool              = "AWSIPPool"
	awsIpPoolDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/awsippools.infrastructure.giantswarm.io/"
)

func NewAWSIPPoolCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindAWSIPPool)
}

func NewAWSIPPoolTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAWSIPPool,
	}
}

// NewAWSIPPoolCR returns an AWSIPPool Custom Resource.
func NewAWSIPPoolCR() *AWSIPPool {
	return &AWSIPPool{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotation.Docs: awsIpPoolDocumentationLink,
			},
		},
		TypeMeta: NewAWSIPPoolTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;cluster-api;giantswarm
// +k8s:openapi-gen=true

// AWSIPPool is the infrastructure provider referenced in upstream CAPI Cluster
// CRs.
type AWSIPPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AWSIPPoolSpec `json:"spec"`
}

// AWSIPPoolSpec is the spec part for the AWSIPPool resource.
// +k8s:openapi-gen=true
type AWSIPPoolSpec struct {
	// IPv4 address block in CIDR notation.
	CIDRBlock string `json:"cidrBlock,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSIPPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AWSIPPool `json:"items"`
}
