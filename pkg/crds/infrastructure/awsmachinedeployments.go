package infrastructure

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const awsmachinedeploymentsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: awsmachinedeployments.infrastructure.giantswarm.io
spec:
  group: infrastructure.giantswarm.io
  names:
    kind: AWSMachineDeployment
    listKind: AWSMachineDeploymentList
    plural: awsmachinedeployments
    singular: awsmachinedeployment
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: "AWSMachineDeployment is the infrastructure provider referenced
        in upstream CAPI MachineDeployment CRs. \n     apiVersion: infrastructure.giantswarm.io/v1alpha2
        \    kind: AWSMachineDeployment     metadata:       labels:         aws-operator.giantswarm.io/version:
        6.2.0         cluster-operator.giantswarm.io/version: 0.17.0         giantswarm.io/cluster:
        8y5kc         giantswarm.io/organization: \"giantswarm\"         giantswarm.io/machine-deployment:
        al9qy         release.giantswarm.io/version: 7.3.1       name: al9qy     spec:
        \      nodePool:         description: my fancy node pool         machine:
        \          dockerVolumeSizeGB: 100           kubeletVolumeSizeGB: 100         scaling:
        \          max: 3           min: 3       provider:         availabilityZones:
        \          - eu-central-1a         instanceDistribution:           onDemandBaseCapacity:
        0           onDemandPercentageAboveBaseCapacity: 0         worker:           instanceType:
        m4.xlarge           useAlikeInstanceTypes: true     status:       provider:
        \        worker:           instanceTypes:             - \"m4.xlarge\"             -
        \"m5.xlarge\"           spotInstances: 39"
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
            nodePool:
              properties:
                description:
                  type: string
                machine:
                  properties:
                    dockerVolumeSizeGB:
                      type: integer
                    kubeletVolumeSizeGB:
                      type: integer
                  required:
                  - dockerVolumeSizeGB
                  - kubeletVolumeSizeGB
                  type: object
                scaling:
                  properties:
                    max:
                      type: integer
                    min:
                      type: integer
                  required:
                  - max
                  - min
                  type: object
              required:
              - description
              - machine
              - scaling
              type: object
            provider:
              properties:
                availabilityZones:
                  items:
                    type: string
                  type: array
                instanceDistribution:
                  properties:
                    onDemandBaseCapacity:
                      type: integer
                    onDemandPercentageAboveBaseCapacity:
                      type: integer
                  required:
                  - onDemandBaseCapacity
                  - onDemandPercentageAboveBaseCapacity
                  type: object
                worker:
                  properties:
                    instanceType:
                      type: string
                    useAlikeInstanceTypes:
                      type: boolean
                  required:
                  - instanceType
                  - useAlikeInstanceTypes
                  type: object
              required:
              - availabilityZones
              - instanceDistribution
              - worker
              type: object
          required:
          - nodePool
          - provider
          type: object
        status:
          properties:
            provider:
              properties:
                worker:
                  properties:
                    instanceTypes:
                      items:
                        type: string
                      type: array
                    spotInstances:
                      type: integer
                  required:
                  - instanceTypes
                  - spotInstances
                  type: object
              required:
              - worker
              type: object
          required:
          - provider
          type: object
      required:
      - spec
      - status
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

func NewAWSMachineDeploymentCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(awsmachinedeploymentsYAML), &crd)
	return &crd
}
