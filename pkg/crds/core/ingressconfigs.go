package core

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const ingressconfigsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: ingressconfigs.core.giantswarm.io
spec:
  group: core.giantswarm.io
  names:
    kind: IngressConfig
    listKind: IngressConfigList
    plural: ingressconfigs
    singular: ingressconfig
  scope: Namespaced
  validation:
    openAPIV3Schema:
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
            guestCluster:
              properties:
                id:
                  type: string
                namespace:
                  type: string
                service:
                  type: string
              required:
              - id
              - namespace
              - service
              type: object
            hostCluster:
              properties:
                ingressController:
                  properties:
                    configMap:
                      type: string
                    namespace:
                      type: string
                    service:
                      type: string
                  required:
                  - configMap
                  - namespace
                  - service
                  type: object
              required:
              - ingressController
              type: object
            protocolPorts:
              items:
                properties:
                  ingressPort:
                    type: integer
                  lbPort:
                    type: integer
                  protocol:
                    type: string
                required:
                - ingressPort
                - lbPort
                - protocol
                type: object
              type: array
            versionBundle:
              properties:
                version:
                  type: string
              required:
              - version
              type: object
          required:
          - guestCluster
          - hostCluster
          - protocolPorts
          - versionBundle
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

func NewIngressConfigCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(ingressconfigsYAML), &crd)
	return &crd
}