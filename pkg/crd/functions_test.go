package crd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var testCRDs = []runtime.Object{
	&v1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: crdV1GVK.Version,
			Kind:       crdV1GVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "examples.example.giantswarm.io",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "example.giantswarm.io",
			Names: v1.CustomResourceDefinitionNames{
				Plural:   "examples",
				Singular: "example",
				Kind:     "Example",
				ListKind: "ExampleList",
				Categories: []string{
					"example",
				},
			},
		},
		Status: v1.CustomResourceDefinitionStatus{},
	},
}

const testCRDsYAML = `
---
apiVersion: v1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  name: examples.example.giantswarm.io
spec:
  group: example.giantswarm.io
  names:
    categories:
    - example
    kind: Example
    listKind: ExampleList
    plural: examples
    singular: example
  scope: ""
  versions: null
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`

func Test_writeCRDs(t *testing.T) {
	var writer strings.Builder
	err := writeObjects(&writer, testCRDs)
	require.Nil(t, err, err)
	require.Equal(t, writer.String(), testCRDsYAML)
}
