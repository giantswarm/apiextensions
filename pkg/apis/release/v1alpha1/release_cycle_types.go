package v1alpha1

import (
	"github.com/giantswarm/to"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindReleaseCycle = "ReleaseCycle"

	CyclePhaseUpcoming ReleaseCyclePhase = "upcoming"
	CyclePhaseEnabled  ReleaseCyclePhase = "enabled"
	CyclePhaseDisabled ReleaseCyclePhase = "disabled"
	CyclePhaseEOL      ReleaseCyclePhase = "eol"
)

type ReleaseCyclePhase string

var releaseCycleValidation = &apiextensionsv1beta1.CustomResourceValidation{
	// See http://json-schema.org/learn.
	OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{
		Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
			"spec": {
				Type: "object",
				Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
					"disabledDate": {
						Type:   "string",
						Format: "date",
					},
					"enabledDate": {
						Type:   "string",
						Format: "date",
					},
					"phase": {
						Enum: []apiextensionsv1beta1.JSON{
							{Raw: []byte("upcoming")},
							{Raw: []byte("enabled")},
							{Raw: []byte("disabled")},
							{Raw: []byte("eol")},
						},
					},
					"release": {
						Type: "object",
						Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
							"name": {
								Type:      "string",
								MinLength: to.Int64P(3),
							},
						},
					},
				},
				Required: []string{
					"phase",
					"release",
				},
			},
			"status": {
				Type: "object",
			},
		},
	},
}

func NewReleaseCycleCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "releasecycles.release.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "release.giantswarm.io",
			Scope:   "Cluster",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "ReleaseCycle",
				Plural:   "releasecycles",
				Singular: "releasecycle",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
			Validation: releaseCycleValidation,
		},
	}
}

func NewReleaseCycleTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindReleaseCycle,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseCycle struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              ReleaseCycleSpec   `json:"spec" yaml:"spec"`
	Status            ReleaseCycleStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type ReleaseCycleSpec struct {
	// DisabledDate is the date of the cycle phase being changed to "disabled".
	DisabledDate DeepCopyTime `json:"disabledDate,omitempty" yaml:"disabledDate,omitempty"`
	// EnabledDate is the date of the cycle phase being changed to "enabled".
	EnabledDate DeepCopyTime `json:"enabledDate,omitempty" yaml:"enabledDate,omitempty"`
	// Phase is the release phase. It can be one of: "upcoming", "enabled",
	// "disabled", "eol".
	Phase ReleaseCyclePhase `json:"phase" yaml:"phase"`
	// Release contains information about Release CR referenced by this ReleaseCycle CR.
	Release ReleaseCycleSpecRelease `json:"release" yaml:"release"`
}

type ReleaseCycleSpecRelease struct {
	// Name is the name of the Release CR referenced by this ReleaseCycle CR
	Name string `json:"name" yaml:"name"`
}

type ReleaseCycleStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseCycleList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []ReleaseCycle `json:"items" yaml:"items"`
}
