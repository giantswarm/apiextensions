package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/annotation"
	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

const (
	kindConfig              = "Config"
	configDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/configs." + group + "/"
)

func NewConfigCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindConfig)
}

func NewConfigTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindConfig,
	}
}

func NewConfigCR() *Config {
	return &Config{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotation.Docs: configDocumentationLink,
			},
		},
		TypeMeta: NewConfigTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=config;giantswarm
// +k8s:openapi-gen=true

// Config represents configuration of a Management Cluster App.
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Spec part of the Config resource.
	Status ConfigStatus `json:"status,omitempty"`
}

// ConfigSpec is the spec part for the Config resource.
// +k8s:openapi-gen=true
type ConfigSpec struct {
	// App details for which the configuration should be generated.
	App ConfigSpecApp `json:"app"`
}

// ConfigSpecApp holds the information about the App to be configured.
// +k8s:openapi-gen=true
type ConfigSpecApp struct {
	// Name is the name of the App.
	Name string `json:"name"`
	// Version is the version of the App.
	Version string `json:"version"`
	// Catalog is the name of App's App Catalog.
	Catalog string `json:"catalog"`
}

// ConfigStatus holds status information about the generated configuration.
// +k8s:openapi-gen=true
type ConfigStatus struct {
	// +kubebuilder:validation:Optional
	// App details for which the configuration was generated.
	App ConfigStatusApp `json:"app,omitempty"`
	// +kubebuilder:validation:Optional
	// Config holds the references to the generated configuration.
	Config ConfigStatusConfig `json:"config,omitempty"`
	// +kubebuilder:validation:Optional
	// Version of the giantswarm/config repository used to generate the
	// configuration.
	Version string `json:"version,omitempty"`
}

// ConfigStatusApp holds the information about the App used to generate
// referenced configuration.
// +k8s:openapi-gen=true
type ConfigStatusApp struct {
	// Name is the name of the App.
	Name string `json:"name"`
	// Version is the version of the App.
	Version string `json:"version"`
	// Catalog is the name of App's App Catalog.
	Catalog string `json:"catalog"`
}

// ConfigStatusConfig holds configuration ConfigMap and Secret references to be
// used to configure the App.
// +k8s:openapi-gen=true
type ConfigStatusConfig struct {
	ConfigMapRef ConfigStatusConfigConfigMapRef `json:"configMapRef"`
	SecretRef    ConfigStatusConfigSecretRef    `json:"secretRef"`
}

// ConfigStatusConfigConfigMapRef contains a reference to the generated ConfigMap.
// +k8s:openapi-gen=true
type ConfigStatusConfigConfigMapRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// ConfigStatusConfigSecretRef contains a reference to the generated Secret.
// +k8s:openapi-gen=true
type ConfigStatusConfigSecretRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Config `json:"items"`
}
