package v1alpha2

import (
	"fmt"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	kindAWSControlPlane = "AWSControlPlane"

	// TODO: change to "https://docs.giantswarm.io/reference/cp-k8s-api/awscontrolplanes.infrastructure.giantswarm.io/"
	// after this has been first published.
	awsControlPlaneDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/"
)

const awsControlPlaneCRDYAML = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: awscontrolplanes.infrastructure.giantswarm.io
spec:
  group: infrastructure.giantswarm.io
  names:
    kind: AWSControlPlane
    plural: awscontrolplanes
    singular: awscontrolplane
  scope: Namespaced
  subresources:
    status: {}
  versions:
  - name: v1alpha1
    served: false
    storage: false
    schema:
      openAPIV3Schema:
        type: object
        properties: {}
  - name: v1alpha2
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        description: |
          Configuration for the master nodes (also called Control Plane) of a
          tenant cluster on AWS.
        type: object
        properties:
          spec:
            type: object
            properties:
              availabilityZones:
                description: |
                  Configures which AWS availability zones to use by master nodes.
                  We support either 1 or 3 availability zones.
                type: array
                items:
                  description: |
                    Identifier of an availability zone to use.
                  type: string
              instanceType:
                description: |
                  EC2 instance type to use for all master nodes.
                type: string
  conversion:
    strategy: None
`

var awsControlPlaneCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(awsControlPlaneCRDYAML), &awsControlPlaneCRD)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func NewAWSControlPlaneCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return awsControlPlaneCRD.DeepCopy()
}

func NewAWSControlPlaneTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAWSControlPlane,
	}
}

// NewAWSControlPlaneCR returns an AWSControlPlane Custom Resource.
func NewAWSControlPlaneCR() *AWSControlPlane {
	return &AWSControlPlane{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: awsControlPlaneDocumentationLink,
			},
		},
		TypeMeta: NewAWSControlPlaneTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSControlPlane is the infrastructure provider referenced in ControlPlane
// CRs.
//
//     apiVersion: infrastructure.giantswarm.io/v1alpha2
//     kind: AWSControlPlane
//     metadata:
//       annotations:
//         giantswarm.io/docs: https://docs.giantswarm.io/reference/awscontrolplanes.infrastructure.giantswarm.io/v1alpha2/
//       labels:
//         aws-operator.giantswarm.io/version: "6.2.0"
//         giantswarm.io/cluster: 8y5kc
//         giantswarm.io/organization: giantswarm
//         release.giantswarm.io/version: "7.3.1"
//       name: 8y5kc
//       ownerReferences:
//         - apiVersion: infrastructure.giantswarm.io/v1alpha2
//           kind: G8sControlPlane
//           name: 8y5kc
//     spec:
//       availabilityZones:
//         - eu-central-1a
//         - eu-central-1b
//         - eu-central-1c
//       instanceType: m4.large
//
type AWSControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AWSControlPlaneSpec   `json:"spec" yaml:"spec"`
	Status            AWSControlPlaneStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type AWSControlPlaneSpec struct {
	AvailabilityZones []string `json:"availabilityZones" yaml:"availabilityZones"`
	InstanceType      string   `json:"instanceType" yaml:"instanceType"`
}

// TODO
type AWSControlPlaneStatus struct {
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSControlPlaneList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []AWSControlPlane `json:"items" yaml:"items"`
}
