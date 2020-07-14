package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	crDocsAnnotation            = "giantswarm.io/docs"
	crReleaseNotesURLAnnotation = "giantswarm.io/release-notes"
	kindRelease                 = "Release"
	releaseDocumentationLink    = "https://docs.giantswarm.io/reference/cp-k8s-api/releases.release.giantswarm.io/"
	releaseNotesLink            = "https://github.com/giantswarm/releases"
)

type ReleaseState string

var (
	stateActive     ReleaseState = "active"     // nolint
	stateDeprecated ReleaseState = "deprecated" // nolint
	stateWIP        ReleaseState = "wip"        // nolint
)

func (r ReleaseState) String() string {
	return string(r)
}

func NewReleaseCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindRelease)
}

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
				crDocsAnnotation:            releaseDocumentationLink,
				crReleaseNotesURLAnnotation: releaseNotesLink,
			},
		},
		TypeMeta: NewReleaseTypeMeta(),
	}
}

// +kubebuilder:printcolumn:name="Kubernetes version",type=string,JSONPath=`.spec.components[?(@.name=="kubernetes")].version`,description="Version of the kubernetes component in this release"
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.spec.state`,description="State of the release"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.spec.date`,description="Time since release creation"
// +kubebuilder:printcolumn:name="Release notes",type=string,JSONPath=`.metadata.annotations['giantswarm\.io/release-notes']`,priority=1,description="Release notes for this release"
// +kubebuilder:resource:scope=Cluster,categories=common;giantswarm
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +k8s:openapi-gen=true

// Release is a Kubernetes resource (CR) representing a Giant Swarm tenant cluster release.
type Release struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ReleaseSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status ReleaseStatus `json:"status"`
}

// +k8s:openapi-gen=true
type ReleaseSpec struct {
	// Apps describes apps used in this release.
	Apps []ReleaseSpecApp `json:"apps"`
	// +kubebuilder:validation:MinItems=1
	// Components describes components used in this release.
	Components []ReleaseSpecComponent `json:"components"`
	// Date that the release became active.
	Date *metav1.Time `json:"date"`
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^(active|deprecated|wip)$`
	// State indicates the availability of the release: deprecated, active, or wip.
	State ReleaseState `json:"state"`
}

// +k8s:openapi-gen=true
type ReleaseSpecComponent struct {
	// +kubebuilder:default=control-plane-catalog
	// Catalog specifies the name of the app catalog that this component belongs to.
	Catalog string `json:"catalog,omitempty"`
	// Name of the component.
	Name string `json:"name"`
	// +kubebuilder:validation:Optional
	// Reference is the component's version in the catalog (e.g. my-app-1.2.3 or my-test-component-0.0.0-abc8675309).
	Reference string `json:"reference,omitempty"`
	// +kubebuilder:validation:Optional
	// ReleaseOperatorDeploy informs the release-operator that it should deploy the component.
	ReleaseOperatorDeploy bool `json:"releaseOperatorDeploy,omitempty"`
	// +kubebuilder:validation:Pattern=`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
	// Version of the component.
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type ReleaseSpecApp struct {
	// Version of the upstream component used in the app.
	ComponentVersion string `json:"componentVersion,omitempty"`
	// Name of the app.
	Name string `json:"name"`
	// +kubebuilder:validation:Pattern=`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
	// Version of the app.
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type ReleaseStatus struct {
	// Ready indicates if all components of the release have been deployed.
	Ready bool `json:"ready"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Release `json:"items"`
}
