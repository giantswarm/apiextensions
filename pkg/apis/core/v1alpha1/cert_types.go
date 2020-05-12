package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
	"github.com/giantswarm/apiextensions/pkg/key"
)

const (
	crDocsAnnotation            = "giantswarm.io/docs"
	kindCertConfig              = "CertConfig"
	certConfigDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/certconfigs.core.giantswarm.io/"
)

func NewCertConfigCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindCertConfig)
}

// NewCertConfigTypeMeta returns the type part for the metadata section of a
// CertConfig custom resource.
func NewCertConfigTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindCertConfig,
	}
}

// NewCertConfigCR returns a CertConfig Custom Resource.
func NewCertConfigCR(name string) *CertConfig {
	certConfig := CertConfig{}
	groupVersionKind := metav1.GroupVersionKind{
		Group:   group,
		Version: version,
		Kind:    key.KindCertConfig,
	}
	certConfig.TypeMeta = key.NewTypeMeta(groupVersionKind)
	certConfig.ObjectMeta = key.NewObjectMeta(groupVersionKind)
	certConfig.Name = name
	return &certConfig
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=giantswarm;common
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
