package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/apis/application"
	"github.com/giantswarm/apiextensions/pkg/crd"
	"github.com/giantswarm/apiextensions/pkg/key"
)

// NewAppCatalogCRD returns a CRD defining an AppCatalog.
func NewAppCatalogCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(application.Group, application.KindAppCatalog)
}

// NewAppCatalogCR returns an AppCatalog custom resource.
func NewAppCatalogCR(name, namespace string) *AppCatalog {
	cr := AppCatalog{}
	cr.TypeMeta, cr.ObjectMeta = key.NewMeta(SchemeGroupVersion, application.KindAppCatalog, name, namespace)
	return &cr
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=common;giantswarm
// +kubebuilder:storageversion

type AppCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AppCatalogSpec `json:"spec"`
}

type AppCatalogSpec struct {
	// Title is the name of the app catalog for this CR
	// e.g. Catalog of Apps by Giant Swarm
	Title       string `json:"title"`
	Description string `json:"description"`
	// Config is the config to be applied when apps belonging to this
	// catalog are deployed.
	Config AppCatalogSpecConfig `json:"config"`
	// LogoURL contains the links for logo image file for this app catalog
	LogoURL string `json:"logoURL"`
	// Storage references a map containing values that should be applied to
	// the appcatalog.
	Storage AppCatalogSpecStorage `json:"storage"`
}

type AppCatalogSpecConfig struct {
	// ConfigMap references a config map containing catalog values that
	// should be applied to apps in this catalog.
	ConfigMap AppCatalogSpecConfigConfigMap `json:"configMap"`
	// Secret references a secret containing catalog values that should be
	// applied to apps in this catalog.
	Secret AppCatalogSpecConfigSecret `json:"secret"`
}

type AppCatalogSpecConfigConfigMap struct {
	// Name is the name of the config map containing catalog values to
	// apply, e.g. app-catalog-values.
	Name string `json:"name"`
	// Namespace is the namespace of the catalog values config map,
	// e.g. giantswarm.
	Namespace string `json:"namespace"`
}

type AppCatalogSpecConfigSecret struct {
	// Name is the name of the secret containing catalog values to apply,
	// e.g. app-catalog-secret.
	Name string `json:"name"`
	// Namespace is the namespace of the secret,
	// e.g. giantswarm.
	Namespace string `json:"namespace"`
}

type AppCatalogSpecStorage struct {
	// Type indicates which repository type would be used for this AppCatalog.
	// e.g. helm
	Type string `json:"type"`
	// URL is the link to where this AppCatalog's repository is located
	// e.g. https://example.com/app-catalog/
	URL string `json:"URL"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AppCatalog `json:"items"`
}
