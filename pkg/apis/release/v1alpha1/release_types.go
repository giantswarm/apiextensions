package v1alpha1

import (
	"github.com/giantswarm/apiextensions/pkg/key"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindRelease = "Release"
)

type ReleaseState string

var (
	StateActive     ReleaseState = "active"
	StateDeprecated ReleaseState = "deprecated"
	StateWIP        ReleaseState = "wip"

	validStates = []string{
		string(StateActive),
		string(StateDeprecated),
		string(StateWIP),
	}
)

// NewReleaseCRD returns a new custom resource definition for Release.
func NewReleaseCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	schema := apiextensionsv1beta1.JSONSchemaProps{
		Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
			"spec": specPropertySchema,
		},
	}
	return key.NewCRD(kindRelease, group, version, "Cluster", schema)
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

// Release CRs might look something like the following.
//
//	apiVersion: release.giantswarm.io/v1alpha1
//	kind: Release
//	metadata:
//	  name: 13.0.0
//	spec:
//	  version: 13.0.0
//    state: active
//    apps:
//	    - name: net-exporter
//	      version: 1.0.0
//        componentVersion: 0.2.0
//	  components:
//	    - name: kubernetes
//	      version: 1.18.0-alpha.3
//
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
