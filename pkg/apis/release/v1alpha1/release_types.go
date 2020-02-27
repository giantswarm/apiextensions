package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	kindRelease    = "Release"
	releaseCRDYAML = `apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: Release.release.giantswarm.io
spec:
  group: release.giantswarm.io
  names:
    kind: Release
    plural: releases
    singular: release
    shortNames:
    - rel
  preserveUnknownFields: false
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        required:
        - metadata
        properties:
          metadata:
            type: object
            required:
            - name
            properties:
              name:
                pattern: ^v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$
                type: string
          spec:
            properties:
              apps:
                items:
                  properties:
                    componentVersion:
                      pattern: ^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$
                      type: string
                    name:
                      minLength: 1
                      type: string
                    version:
                      pattern: ^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$
                      type: string
                  required:
                  - name
                  - version
                  type: object
                type: array
              components:
                items:
                  properties:
                    name:
                      minLength: 1
                      type: string
                    version:
                      pattern: ^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$
                      type: string
                  required:
                  - name
                  - version
                  type: object
                minItems: 1
                type: array
              state:
                pattern: ^(active|deprecated|wip)$
                type: string
              version:
                pattern: ^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$
                type: string
            required:
            - components
            - apps
            - state
            - version
            type: object
    served: true
    storage: true
`
)

type ReleaseState string

var (
	StateActive     ReleaseState = "active"
	StateDeprecated ReleaseState = "deprecated"
	StateWIP        ReleaseState = "wip"
)

var releaseCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.UnmarshalStrict([]byte(releaseCRDYAML), &releaseCRD)
	if err != nil {
		panic(err)
	}
}

// NewReleaseCRD returns a new custom resource definition for Release.
func NewReleaseCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return releaseCRD.DeepCopy()
}

func NewReleaseTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindRelease,
	}
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Release struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              ReleaseSpec `json:"spec" yaml:"spec"`
}

type ReleaseSpec struct {
	// Apps describes apps used in this release.
	Apps []ReleaseSpecApp `json:"apps" yaml:"apps"`
	// Components describes components used in this release.
	Components []ReleaseSpecComponent `json:"components" yaml:"components"`
	// State indicates the availability of the release: deprecated, active, or wip.
	State ReleaseState `json:"state" yaml:"state"`
	// Version is the version of the release.
	Version string `json:"version" yaml:"version"`
}

type ReleaseSpecComponent struct {
	// Name of the component.
	Name string `json:"name" yaml:"name"`
	// Version of the component.
	Version string `json:"version" yaml:"version"`
}

type ReleaseSpecApp struct {
	// Version of the upstream component used in the app.
	ComponentVersion string `json:"componentVersion" yaml:"componentVersion"`
	// Name of the app.
	Name string `json:"name" yaml:"name"`
	// Version of the app.
	Version string `json:"version" yaml:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Release `json:"items" yaml:"items"`
}
