package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	crDocsAnnotation            = "giantswarm.io/docs"
	kindCertConfig              = "CertConfig"
	certConfigDocumentationLink = "https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1?tab=doc#CertConfig"
)

const certConfigCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: certconfigs.core.giantswarm.io
spec:
  conversion:
    strategy: None
  group: core.giantswarm.io
  names:
    kind: CertConfig
    listKind: CertConfigList
    plural: certconfigs
    singular: certconfig
  preserveUnknownFields: true
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
`

var certConfigCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(certConfigCRDYAML), &certConfigCRD)
	if err != nil {
		panic(err)
	}
}

// NewCertConfigCRD returns a new custom resource definition for CertConfig.
func NewCertConfigCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return certConfigCRD.DeepCopy()
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
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CertConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              CertConfigSpec `json:"spec"`
}

type CertConfigSpec struct {
	Cert          CertConfigSpecCert          `json:"cert" yaml:"cert"`
	VersionBundle CertConfigSpecVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type CertConfigSpecCert struct {
	AllowBareDomains    bool     `json:"allowBareDomains" yaml:"allowBareDomains"`
	AltNames            []string `json:"altNames" yaml:"altNames"`
	ClusterComponent    string   `json:"clusterComponent" yaml:"clusterComponent"`
	ClusterID           string   `json:"clusterID" yaml:"clusterID"`
	CommonName          string   `json:"commonName" yaml:"commonName"`
	DisableRegeneration bool     `json:"disableRegeneration" yaml:"disableRegeneration"`
	IPSANs              []string `json:"ipSans" yaml:"ipSans"`
	Organizations       []string `json:"organizations" yaml:"organizations"`
	TTL                 string   `json:"ttl" yaml:"ttl"`
}

type CertConfigSpecVersionBundle struct {
	Version string `json:"version" yaml:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CertConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CertConfig `json:"items"`
}
