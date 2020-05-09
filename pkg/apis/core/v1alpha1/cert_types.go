package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
				crDocsAnnotation: certConfigDocumentationLink,
			},
		},
		TypeMeta: NewCertConfigTypeMeta(),
	}
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	AllowBareDomains    bool     `json:"allowBareDomains"`
	AltNames            []string `json:"altNames"`
	ClusterComponent    string   `json:"clusterComponent"`
	ClusterID           string   `json:"clusterID"`
	CommonName          string   `json:"commonName"`
	DisableRegeneration bool     `json:"disableRegeneration"`
	IPSANs              []string `json:"ipSans"`
	Organizations       []string `json:"organizations"`
	TTL                 string   `json:"ttl"`
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
