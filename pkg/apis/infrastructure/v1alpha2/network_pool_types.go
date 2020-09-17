package v1alpha2

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v2/pkg/annotation"
	"github.com/giantswarm/apiextensions/v2/pkg/crd"
)

const (
	kindNetworkPool              = "NetworkPool"
	networkpoolDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/networkpools.infrastructure.giantswarm.io/"
)

func NewNetworkPoolCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindNetworkPool)
}

func NewNetworkPoolTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindNetworkPool,
	}
}

// NewNetworkPoolCR returns an NetworkPool Custom Resource.
func NewNetworkPoolCR() *NetworkPool {
	return &NetworkPool{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotation.Docs: networkpoolDocumentationLink,
			},
		},
		TypeMeta: NewNetworkPoolTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;cluster-api;giantswarm
// +k8s:openapi-gen=true

// NetworkPool is the infrastructure provider referenced in upstream CAPI Cluster
// CRs.
type NetworkPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkPoolSpec `json:"spec"`
}

// NetworkPoolSpec is the spec part for the NetworkPool resource.
// +k8s:openapi-gen=true
type NetworkPoolSpec struct {
	// IPv4 address block in CIDR notation.
	CIDRBlock string `json:"cidrBlock,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NetworkPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []NetworkPool `json:"items"`
}
