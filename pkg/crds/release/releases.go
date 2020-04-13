package release

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const releasesYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: releases.release.giantswarm.io
spec:
  group: release.giantswarm.io
  names:
    kind: Release
    listKind: ReleaseList
    plural: releases
    singular: release
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: "Release is a Kubernetes resource (CR) which is based on the Release
        CRD defined above. \n An example Release resource can be viewed here https://github.com/giantswarm/apiextensions/blob/master/docs/cr/release.giantswarm.io_v1alpha1_release.yaml"
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
            apps:
              description: Apps describes apps used in this release.
              items:
                properties:
                  componentVersion:
                    description: Version of the upstream component used in the app.
                    type: string
                  name:
                    description: Name of the app.
                    type: string
                  version:
                    description: Version of the app.
                    type: string
                required:
                - name
                - version
                type: object
              type: array
            components:
              description: Components describes components used in this release.
              items:
                properties:
                  name:
                    description: Name of the component.
                    type: string
                  version:
                    description: Version of the component.
                    type: string
                required:
                - name
                - version
                type: object
              type: array
            date:
              description: Date that the release became active.
              format: date-time
              type: string
            state:
              description: 'State indicates the availability of the release: deprecated,
                active, or wip.'
              type: string
          required:
          - apps
          - components
          - date
          - state
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

func NewReleaseCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(releasesYAML), &crd)
	return &crd
}