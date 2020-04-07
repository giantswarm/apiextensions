package release

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const releasecyclesYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: releasecycles.release.giantswarm.io
spec:
  group: release.giantswarm.io
  names:
    kind: ReleaseCycle
    listKind: ReleaseCycleList
    plural: releasecycles
    singular: releasecycle
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: "ReleaseCycle CRs might look something like the following. \n \tapiVersion:
        \"release.giantswarm.io/v1alpha1\" \tkind: \"ReleaseCycle\" \tmetadata: \t
        \ name: \"aws.v6.1.0\" \t  labels: \t    giantswarm.io/managed-by: \"opsctl\"
        \t    giantswarm.io/provider: \"aws\" \tspec: \t  disabledDate: 2019-01-12
        \t  enabledDate: 2019-01-08 \t  phase: \"enabled\""
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
            disabledDate:
              description: DisabledDate is the date of the cycle phase being changed
                to "disabled".
              format: date-time
              type: string
            enabledDate:
              description: EnabledDate is the date of the cycle phase being changed
                to "enabled".
              format: date-time
              type: string
            phase:
              description: 'Phase is the release phase. It can be one of: "upcoming",
                "enabled", "disabled", "eol".'
              type: string
          required:
          - phase
          type: object
        status:
          type: object
      required:
      - metadata
      - spec
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`

func NewReleaseCycleCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(releasecyclesYAML), &crd)
	return &crd
}
