package v1alpha1

import (
	"fmt"
	"strings"

	"github.com/giantswarm/to"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

const (
	semverPattern = `^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
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

	statePropertySchema = apiextensionsv1beta1.JSONSchemaProps{
		Type:    "string",
		Pattern: fmt.Sprintf("^(%s)$", strings.Join(validStates, "|")),
	}

	appsPropertySchema = apiextensionsv1beta1.JSONSchemaProps{
		Type: "object",
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
		Type: "object",
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
			"state",
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
				Type:     "array",
				MinItems: to.Int64P(1),
				Items: &apiextensionsv1beta1.JSONSchemaPropsOrArray{
					Schema: &componentsPropertySchema,
				},
			},
			"state":   statePropertySchema,
			"version": versionPropertySchema,
		},
	}
)
