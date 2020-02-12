package v1alpha1

import (
	"github.com/giantswarm/apiextensions/pkg/key"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/to"
)

const (
	kindRelease   = "Release"
	semverPattern = `^(=|>=|<=|=>|=<|>|<|!=|~|~>|\^)?(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
)

var (
	namePropertySchema = apiextensionsv1beta1.JSONSchemaProps{
		Type:      "string",
		MinLength: to.Int64P(1),
	}
	versionPropertySchema = apiextensionsv1beta1.JSONSchemaProps{
		Type:    "string",
		Pattern: semverPattern,
	}
	appsPropertySchema = apiextensionsv1beta1.JSONSchemaProps{
		Type: "array",
		Required: []string{
			"name",
			"version",
		},
		Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
			"name":             namePropertySchema,
			"version":          versionPropertySchema,
			"componentVersion": versionPropertySchema,
		},
	}
	componentsPropertySchema = apiextensionsv1beta1.JSONSchemaProps{
		Type:     "array",
		Required: []string{
			"name",
			"version",
		},
		Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
			"name":    namePropertySchema,
			"version": versionPropertySchema,
		},
	}
	specPropertySchema = apiextensionsv1beta1.JSONSchemaProps{
		Type: "object",
		Required: []string{
			"components",
			"apps",
			"version",
		},
		Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
			"apps": {
				Type: "array",
				Items: &apiextensionsv1beta1.JSONSchemaPropsOrArray{
					Schema: &appsPropertySchema,
				},
			},
			"components": {
				Type: "array",
				MinItems: to.Int64P(1),
				Items: &apiextensionsv1beta1.JSONSchemaPropsOrArray{
					Schema: &componentsPropertySchema,
				},
			},
			"version":    versionPropertySchema,
		},
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
//	apiVersion: "release.giantswarm.io/v1alpha1"
//	kind: "Release"
//	metadata:
//	  name: "v6.1.0"
//	spec:
//    apps:
//	    - name: "net-exporter"
//	      version: "1.0.0"
//        componentVersion: "0.2.0"
//	  components:
//	    - name: "kubernetes"
//	      version: "1.18.0-alpha.3"
//	  version: "13.0.0"
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
	// Name of the app.
	Name string `json:"name" yaml:"name"`
	// Version of the app.
	Version string `json:"version" yaml:"version"`
	// Version of the upstream component used in the app.
	ComponentVersion string `json:"componentVersion" yaml:"componentVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Release `json:"items" yaml:"items"`
}
