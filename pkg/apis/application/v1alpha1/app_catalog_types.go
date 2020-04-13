package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	crDocsAnnotation            = "giantswarm.io/docs"
	kindAppCatalog              = "AppCatalog"
	appCatalogDocumentationLink = "https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1?tab=doc#AppCatalog"
)

const appCatalogCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: appcatalogs.application.giantswarm.io
spec:
  group: application.giantswarm.io
  scope: Cluster
  version: v1alpha1
  names:
    kind: AppCatalog
    plural: appcatalogs
    singular: appcatalog
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: |
        Defines a location where packaged applications are stored and shared,
        ready to be installed in a Kubernetes cluster.
      type: object
      properties:
        spec:
          type: object
          properties:
            title:
              description: |
                User-friendly name of the catalog.    
              type: string
            description:
              description: |
                Additional information regarding the purpose and other details of the catalog.
              type: string
            config:
              description: |
                Defines the reference of a ConfigMap where is saved the default values that 
                will be applied to all applications contained in the catalog.
              type: object
              properties:
                configMap:
                  description: |
                    References a ConfigMap containing catalog values that should be applied to
                    apps installed from this catalog.
                  type: object
                  properties:
                    name:
                      description: |
                        Name of the ConfigMap resource.
                      type: string
                    namespace:
                      description: |
                        Namespace holding the ConfigMap resource.
                      type: string
                  required: ["name", "namespace"]
                secret:
                  description: |
                    Defines the reference of a Secret where is saved the default sensitive configuration
                    that will be applied to all applications contained in the catalog.
                  type: object
                  properties:
                    name:
                      description: |
                        Name of the Secret resource.
                      type: string
                    namespace:
                      description: |
                        Namespace holding the Secret resource.
                      type: string
                  required: ["name", "namespace"]
            logoURL:
              description: |
                The logo URL pointing to the image file to be used when displaying this catalog.
              type: string
            storage:
              description: |
                Defines the type of storage supported by the catalog.
              type: object 
              properties:
                type:
                  description: |
                    Determines the type of storage. Currently only 'helm' is available.
                    - Helm type storage use the known Helm Chart Repository format to store the chart
                      packages and expose some metadata in the file 'index.yaml' to manage the catalog.
                  type: string
                URL:
                  description: |
                    URL to the app repository.
                  type: string
                  format: uri 
              required: ["type", "URL"]
          required: ["title", "description", "storage"]
`

var appCatalogCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(appCatalogCRDYAML), &appCatalogCRD)
	if err != nil {
		panic(err)
	}
}

// NewAppCatalogCRD returns a new custom resource definition for AppCatalog.
func NewAppCatalogCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return appCatalogCRD.DeepCopy()
}

func NewAppCatalogTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAppCatalog,
	}
}

// NewAppCatalogCR returns an AppCatalog Custom Resource.
func NewAppCatalogCR() *AppCatalog {
	return &AppCatalog{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: appCatalogDocumentationLink,
			},
		},
		TypeMeta: NewAppCatalogTypeMeta(),
	}
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AppCatalogSpec `json:"spec"`
}

type AppCatalogSpec struct {
	// Title is the name of the app catalog for this CR
	// e.g. Catalog of Apps by Giant Swarm
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	// Config is the config to be applied when apps belonging to this
	// catalog are deployed.
	Config AppCatalogSpecConfig `json:"config" yaml:"config"`
	// LogoURL contains the links for logo image file for this app catalog
	LogoURL string `json:"logoURL" yaml:"logoURL"`
	// Storage references a map containing values that should be applied to
	// the appcatalog.
	Storage AppCatalogSpecStorage `json:"storage" yaml:"storage"`
}

type AppCatalogSpecConfig struct {
	// ConfigMap references a config map containing catalog values that
	// should be applied to apps in this catalog.
	ConfigMap AppCatalogSpecConfigConfigMap `json:"configMap" yaml:"configMap"`
	// Secret references a secret containing catalog values that should be
	// applied to apps in this catalog.
	Secret AppCatalogSpecConfigSecret `json:"secret" yaml:"secret"`
}

type AppCatalogSpecConfigConfigMap struct {
	// Name is the name of the config map containing catalog values to
	// apply, e.g. app-catalog-values.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the catalog values config map,
	// e.g. giantswarm.
	Namespace string `json:"namespace" yaml:"namespace"`
}

type AppCatalogSpecConfigSecret struct {
	// Name is the name of the secret containing catalog values to apply,
	// e.g. app-catalog-secret.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the secret,
	// e.g. giantswarm.
	Namespace string `json:"namespace" yaml:"namespace"`
}

type AppCatalogSpecStorage struct {
	// Type indicates which repository type would be used for this AppCatalog.
	// e.g. helm
	Type string `json:"type" yaml:"type"`
	// URL is the link to where this AppCatalog's repository is located
	// e.g. https://giantswarm.github.com/app-catalog/.
	URL string `json:"URL" yaml:"URL"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AppCatalog `json:"items"`
}
