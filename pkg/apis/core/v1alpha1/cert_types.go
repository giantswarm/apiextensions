package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	crDocsAnnotation            = "giantswarm.io/docs"
	kindCertConfig              = "CertConfig"
	certConfigDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/certconfigs.core.giantswarm.io/"
)

func NewCertConfigCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindCertConfig)
}

// NewCertConfigTypeMeta returns the type part for the metadata section of a
// CertConfig custom resource.
func NewCertConfigTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindCertConfig,
	}
}

// NewCertConfigCR returns an AWSCluster Custom Resource.
func NewCertConfigCR() *CertConfig {
	return &CertConfig{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: certConfigDocumentationLink,
			},
		},
		TypeMeta: NewCertConfigTypeMeta(),
	}
}

// +genclient
// +genclient:noStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=common;giantswarm
// +kubebuilder:storageversion
type CertConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              CertConfigSpec `json:"spec"`
}

type CertConfigSpec struct {
	Cert          CertConfigSpecCert          `json:"cert"`
	VersionBundle CertConfigSpecVersionBundle `json:"versionBundle"`
}

type CertConfigSpecCert struct {
	AllowBareDomains bool `json:"allowBareDomains"`
	// +kubebuilder:validation:Optional
	// +nullable
	AltNames            []string `json:"altNames,omitempty"`
	ClusterComponent    string   `json:"clusterComponent"`
	ClusterID           string   `json:"clusterID"`
	CommonName          string   `json:"commonName"`
	DisableRegeneration bool     `json:"disableRegeneration"`
	// +kubebuilder:validation:Optional
	// +nullable
	IPSANs []string `json:"ipSans,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	Organizations []string `json:"organizations,omitempty"`
	TTL           string   `json:"ttl"`
}

type CertConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CertConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CertConfig `json:"items"`
}
