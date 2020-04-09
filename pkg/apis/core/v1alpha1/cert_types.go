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
    schema:
      openAPIV3Schema:
        description: |
          The CertConfig resource is used in a Giant Swarm installation to ensure TLS communication between
          a component (e. g. prometheus) and the tenant cluster nodes. It is reconciled by cert-operator.
        properties:
          spec:
            type: object
            properties:
              cert:
                description: |
                  Defines an X.509 certificate to be ensured by cert-operator.
                type: object
                properties:
                  allowBareDomains:
                    description: |
                      Specifies if clients can request certificates matching the value of the actual
                      domains themselves.
                    type: bool
                  altNames:
                    description: |
                      Specifies requested Subject Alternative Names, in a comma-delimited list. These
                      can be host names or email addresses; they will be parsed into their respective
                      fields.
                    type: array
                    items:
                      type: string
                  clusterComponent:
                    description: |
                      Name of the component this certificate is for.
                    type: string
                  clusterID:
                    description: |
                      Unique identifier of the tenant cluster this certificate is for.
                    type: string
                  commonName:
                    description: |
                      The value of the Common Name (CN) attribute of the certificate.
                  disableRegeneration:
                    description: |
                      Toggles the automatic recreation before expiry.
                    type: bool
                  ipSans:
                    description: |
                      Specifies requested IP Subject Alternative Names to be set in the
                      certificate.
                    type: array
                    items:
                      type: string
                  organizations:
                    description: |
                      Organizations to set in the Organizations (O) attribute of the
                      certificate.
                    type: array
                    items:
                      type: string
                  ttl:
                    description: |
                      Expiry duration after creation. The value must consist of a number
                      combined with a unit, without blanks, to be parsed by the Go
                      [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration) function.
                    type: string
              versionBundle:
                description: TODO
                type: object
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
