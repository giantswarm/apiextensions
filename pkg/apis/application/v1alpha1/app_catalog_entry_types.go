package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v2/pkg/annotation"
	"github.com/giantswarm/apiextensions/v2/pkg/crd"
)

const (
	kindAppCatalogEntry              = "AppCatalogEntry"
	AppCatalogEntryDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/AppCatalogEntrys.application.giantswarm.io/"
)

func NewAppCatalogEntryCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindAppCatalogEntry)
}

func NewAppCatalogEntryTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAppCatalogEntry,
	}
}

// NewAppCatalogEntryCR returns an AppCatalogEntry Custom Resource.
func NewAppCatalogEntryCR() *AppCatalogEntry {
	return &AppCatalogEntry{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotation.Docs: AppCatalogEntryDocumentationLink,
			},
		},
		TypeMeta: NewAppCatalogEntryTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=common;giantswarm
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
// AppCatalogEntry represents an entry of an app in a catalog of managed apps.
type AppCatalogEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppCatalogEntrySpec `json:"spec"`
}

// +k8s:openapi-gen=true
type AppCatalogEntrySpec struct {
	// AppName is the name of the app this entry belongs to.
	// e.g. nginx-ingress-controller-app
	AppName string `json:"appName"`
	// AppVersion is the upstream version of the app for this entry.
	// e.g. v0.35.0
	AppVersion string `json:"appVersion"`
	// Catalog is the name of the app catalog this entry belongs to.
	// e.g. giantswarm
	Catalog AppCatalogEntrySpecCatalog `json:"catalog"`
	// Chart is metadata from the Chart.yaml of the app this entry belongs to.
	Chart AppCatalogEntrySpecChart `json:"chart,omitempty"`
	// DateCreated is when this entry was first created.
	// e.g. 2020-09-02T09:40:39.223638219Z
	DateCreated *metav1.Time `json:"dateCreated"`
	// DateUpdated is when this entry was last updated.
	// e.g. 2020-09-02T09:40:39.223638219Z
	DateUpdated *metav1.Time `json:"dateUpdated"`
	// Version is the version of the app chart for this entry.
	// e.g. 1.9.2
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type AppCatalogEntrySpecCatalog struct {
	// Name is the name of the app catalog this entry belongs to.
	// e.g. giantswarm-catalog
	Name string `json:"name"`
	// +kubebuilder:validation:Optional
	// Namespace is the namespace of the catalog. It is empty while the
	// appcatalog CRD is cluster scoped.
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type AppCatalogEntrySpecChart struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// Home is the URL of this projects home page.
	Home string `json:"home,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Icon is a URL to an SVG or PNG image to be used as an icon.
	Icon string `json:"icon,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCatalogEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AppCatalogEntry `json:"items"`
}
