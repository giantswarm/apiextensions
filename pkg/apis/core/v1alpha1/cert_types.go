package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/apis/core"
	"github.com/giantswarm/apiextensions/pkg/crd"
	"github.com/giantswarm/apiextensions/pkg/key"
)

// NewCertConfigCRD returns a CRD defining a CertConfig.
func NewCertConfigCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(core.Group, core.KindCertConfig)
}

// NewCertConfigCR returns a CertConfig custom resource.
func NewCertConfigCR(name, namespace string) *CertConfig {
	cr := CertConfig{}
	cr.TypeMeta, cr.ObjectMeta = key.NewMeta(SchemeGroupVersion, core.KindCertConfig, name, namespace)
	return &cr
}

// +genclient
// +genclient:noStatus
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
