package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	crDocsAnnotation         = "giantswarm.io/docs"
	kindRelease              = "Release"
	releaseDocumentationLink = "https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1?tab=doc#Release"
)

type ReleaseState string

var (
	stateActive     ReleaseState = "active"
	stateDeprecated ReleaseState = "deprecated"
	stateWIP        ReleaseState = "wip"
)

func NewReleaseTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindRelease,
	}
}

func NewReleaseCR() *Release {
	return &Release{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: releaseDocumentationLink,
			},
		},
		TypeMeta: NewReleaseTypeMeta(),
	}
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Release is a Kubernetes resource (CR) which is based on the Release CRD defined above.
//
// An example Release resource can be viewed here
// https://github.com/giantswarm/apiextensions/blob/master/docs/cr/release.giantswarm.io_v1alpha1_release.yaml
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
	// Date that the release became active.
	Date *metav1.Time `json:"date" yaml:"date"`
	// State indicates the availability of the release: deprecated, active, or wip.
	State ReleaseState `json:"state" yaml:"state"`
}

type ReleaseSpecComponent struct {
	// Name of the component.
	Name string `json:"name" yaml:"name"`
	// +kubebuilder:validation:Pattern=^v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$
	// Version of the component.
	Version string `json:"version" yaml:"version"`
}

type ReleaseSpecApp struct {
	// Version of the upstream component used in the app.
	ComponentVersion string `json:"componentVersion,omitempty" yaml:"componentVersion,omitempty"`
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
