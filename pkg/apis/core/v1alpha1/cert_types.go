package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/key"
)

// NewCertConfigCR returns a CertConfig Custom Resource.
func NewCertConfigCR(name string) *CertConfig {
	certConfig := CertConfig{}
	groupVersionKind := metav1.GroupVersionKind{
		Group:   key.GroupApplication,
		Version: version,
		Kind:    key.KindApp,
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
