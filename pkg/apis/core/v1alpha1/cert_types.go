package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/annotation"
)

const (
	kindCertConfig              = "CertConfig"
	certConfigDocumentationLink = "https://docs.giantswarm.io/ui-api/management-api/crd/certconfigs.core.giantswarm.io/"
)

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
				annotation.Docs: certConfigDocumentationLink,
			},
		},
		TypeMeta: NewCertConfigTypeMeta(),
	}
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=common;giantswarm
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
// CertConfig specifies details for an X.509 certificate to be issued, handled by cert-operator.
type CertConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              CertConfigSpec `json:"spec"`
}

// +k8s:openapi-gen=true
type CertConfigSpec struct {
	// Specifies the configurable certificate details.
	Cert CertConfigSpecCert `json:"cert"`
	// Specifies the cert-operator version to use.
	VersionBundle CertConfigSpecVersionBundle `json:"versionBundle"`
}

// +k8s:openapi-gen=true
type CertConfigSpecCert struct {
	AllowBareDomains bool `json:"allowBareDomains"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Subject Alternative Names to be set in the certificate.
	AltNames []string `json:"altNames,omitempty"`
	// Host name of the service to create the certificate for.
	ClusterComponent string `json:"clusterComponent"`
	// Workload cluster ID to issue the certificate for.
	ClusterID string `json:"clusterID"`
	// Full common name (CN).
	CommonName string `json:"commonName"`
	// If set, cert-operator will forbid updating this certificate.
	DisableRegeneration bool `json:"disableRegeneration"`
	// +kubebuilder:validation:Optional
	// +nullable
	// List of IP addresses to be set as SANs (Subject Alternative Names) in the certificate.
	IPSANs []string `json:"ipSans,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// List of organizations to set in the certificate.
	Organizations []string `json:"organizations,omitempty"`
	// Expiry time as a Golang duration string, e. g. "1d" for one day.
	TTL string `json:"ttl"`
}

// +k8s:openapi-gen=true
type CertConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CertConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CertConfig `json:"items"`
}
