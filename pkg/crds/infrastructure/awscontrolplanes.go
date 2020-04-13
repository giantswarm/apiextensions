package infrastructure

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const awscontrolplanesYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: awscontrolplanes.infrastructure.giantswarm.io
spec:
  group: infrastructure.giantswarm.io
  names:
    kind: AWSControlPlane
    listKind: AWSControlPlaneList
    plural: awscontrolplanes
    singular: awscontrolplane
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: "AWSControlPlane is the infrastructure provider referenced in ControlPlane
        CRs. \n     apiVersion: infrastructure.giantswarm.io/v1alpha2     kind: AWSControlPlane
        \    metadata:       annotations:         giantswarm.io/docs: https://docs.giantswarm.io/reference/awscontrolplanes.infrastructure.giantswarm.io/v1alpha2/
        \      labels:         aws-operator.giantswarm.io/version: \"6.2.0\"         giantswarm.io/cluster:
        8y5kc         giantswarm.io/organization: giantswarm         release.giantswarm.io/version:
        \"7.3.1\"       name: 8y5kc       ownerReferences:         - apiVersion: infrastructure.giantswarm.io/v1alpha2
        \          kind: G8sControlPlane           name: 8y5kc     spec:       availabilityZones:
        \        - eu-central-1a         - eu-central-1b         - eu-central-1c       instanceType:
        m4.large"
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          properties:
            availabilityZones:
              items:
                type: string
              type: array
            instanceType:
              type: string
          required:
          - availabilityZones
          - instanceType
          type: object
        status:
          description: TODO
          properties:
            status:
              type: string
          type: object
      required:
      - spec
      type: object
  version: v1alpha2
  versions:
  - name: v1alpha2
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`

func NewAWSControlPlaneCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(awscontrolplanesYAML), &crd)
	return &crd
}