package v1alpha1

import (
	"github.com/ghodss/yaml"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindAppCatalogEntry = "AppCatalogEntry"
)

const appCatalogEntryCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: appcatalogentries.application.giantswarm.io
spec:
  group: application.giantswarm.io
  scope: Cluster
  version: v1alpha1
  names:
    kind: AppCatalogEntry
    plural: appcatalogentries
    singular: appcatalogentry
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          type: object
          properties:
            name:
              type: string
            versions:
              type: array
              items:
                type: string
                pattern: "^v.+"
          required: ["description", "versions"]
`

var appCatalogEntryCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(appCatalogEntryCRDYAML), &appCatalogEntryCRD)
	if err != nil {
		panic(err)
	}
}

// NewAppCatalogEntryCRD returns a new custom resource definition for AppCatalogEntry.
// This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: appcatalog.application.giantswarm.io
//     spec:
//       group: application.giantswarm.io
//       scope: Cluster
//       version: v1alpha1
//       names:
//         kind: AppCatalogEntry
//         plural: appcatalogentries
//         singular: appcatalogentry
//
func NewAppCatalogEntryCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return appCatalogEntryCRD.DeepCopy()
}

func NewAppCatalogEntryTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindAppCatalogEntry,
	}
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppCatalogEntry CRs might look something like the following.
//
//    apiVersion: application.giantswarm.io/v1alpha1
//    kind: AppCatalogEntry
//    metadata:
//      name: "chart-operator"
//      labels:
//        app-operator.giantswarm.io/version: "1.0.0"
//
//    spec:
//      description: A Helm chart for the chart-operator
//      versions:
//        - v0.9.0
//        - v0.9.1
//
type AppCatalogEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AppCatalogEntrySpec `json:"spec"`
}

type AppCatalogEntrySpec struct {
	// Apps is the list which contains app name and available version list
	Description string
	Versions    []string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCatalogEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AppCatalogEntry `json:"items"`
}
