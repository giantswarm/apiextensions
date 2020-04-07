package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const azureToolCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: azuretool.tools.giantswarm.io
spec:
  group: tools.giantswarm.io
  scope: Namespaced
  version: v1alpha1
  names:
    kind: AzureTool
    plural: azuretools
    singular: azuretool
  validation:
    openAPIV3Schema:
      properties:
        spec:
          type: object
          properties:
            workspace:
              type: object
              properties:
                id:
                  type: string
                mode:
                  type: string
          required:
          - workspace
`

var azureToolCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(azureToolCRDYAML), &azureToolCRD)
	if err != nil {
		panic(err)
	}
}

func NewAzureToolCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return azureToolCRD.DeepCopy()
}

type AzureTool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AzureToolSpec `json:"spec"`
}

type AzureToolSpec struct {
	Workspace AzureToolWorkspace `json:"workspace" yaml:"workspace"`
}

type AzureToolWorkspace struct {
	ID   string `json:"id" yaml:"id"`
	Mode string `json:"mode" yaml:"mode"`
}
