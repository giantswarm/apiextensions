package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewAppCatalogConfigCRD returns a new custom resource definition for AppCatalogConfig.
// This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: appcatalogconfigs.core.giantswarm.io
//     spec:
//       group: core.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: AppCatalogConfig
//         plural: appcatalogconfigs
//         singular: appcatalogconfig
//

func NewAppCatalogConfigCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "appcatalogconfigs.core.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "core.giantswarm.io",
			Scope:   "Cluster",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "AppCatalogConfig",
				Plural:   "appcatalogconfigs",
				Singular: "appcatalogconfig",
			},
		},
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AppCatalogConfigSpec `json:"spec"`
	Status            ChartConfigStatus    `json:"status"`
}

type AppCatalogConfigSpec struct {
	AppCatalog    AppCatalogSpec                    `json:"appCatalog" yaml:"appCatalog"`
	VersionBundle AppCatalogConfigSpecVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type AppCatalogSpec struct {
	// Title is the name of the app catalog for this CR
	// e.g. Catalog of Apps by Giant Swarm
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	// CatalogStorage references a map containing values that should be
	// applied to the appcatalog.
	CatalogStorage AppCatalogSpecCatalogStorage `json:"catalogStorage" yaml:"catalogStorage"`
	// LogoURL contains the links for logo image file for this app catalog
	LogoURL string `json:"logoURL" yaml:"logoURL"`
}

type AppCatalogSpecCatalogStorage struct {
	// Type indicates which package type would use for this AppCatalog
	// e.g. helm
	Type string `json:"type" yaml:"type"`
	// URL is the link where this AppCatalog's package file have been located
	// e.g. kube-system.
	URL string `json:"URL" yaml:"URL"`
}

type AppCatalogConfigSpecVersionBundle struct {
	Version string `json:"version" yaml:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AppCatalog `json:"items"`
}
