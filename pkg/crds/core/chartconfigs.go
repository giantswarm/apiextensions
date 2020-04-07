package core

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const chartconfigsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: chartconfigs.core.giantswarm.io
spec:
  group: core.giantswarm.io
  names:
    kind: ChartConfig
    listKind: ChartConfigList
    plural: chartconfigs
    singular: chartconfig
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
            chart:
              properties:
                channel:
                  description: Channel is the name of the Appr channel to reconcile
                    against, e.g. 1-0-stable.
                  type: string
                configMap:
                  description: ConfigMap references a config map containing values
                    that should be applied to the chart.
                  properties:
                    name:
                      description: Name is the name of the config map containing chart
                        values to apply, e.g. node-exporter-chart-values.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the values config
                        map, e.g. kube-system.
                      type: string
                    resourceVersion:
                      description: ResourceVersion is the Kubernetes resource version
                        of the configmap. Used to detect if the configmap has changed,
                        e.g. 12345.
                      type: string
                  required:
                  - name
                  - namespace
                  - resourceVersion
                  type: object
                name:
                  description: Name is the name of the Helm chart to deploy, e.g.
                    kubernetes-node-exporter.
                  type: string
                namespace:
                  description: Namespace is the namespace where the Helm chart is
                    to be deployed, e.g. giantswarm.
                  type: string
                release:
                  description: Release is the name of the Helm release when the chart
                    is deployed, e.g. node-exporter.
                  type: string
                secret:
                  description: Secret references a secret containing secret values
                    that should be applied to the chart.
                  properties:
                    name:
                      description: Name is the name of the secret containing chart
                        values to apply, e.g. node-exporter-chart-secret.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the secret, e.g.
                        kube-system.
                      type: string
                    resourceVersion:
                      description: ResourceVersion is the Kubernetes resource version
                        of the secret. Used to detect if the secret has changed, e.g.
                        12345.
                      type: string
                  required:
                  - name
                  - namespace
                  - resourceVersion
                  type: object
                userConfigMap:
                  description: UserConfigMap references a config map containing custom
                    values. These custom values are specified by the user to override
                    default values.
                  properties:
                    name:
                      description: Name is the name of the config map containing chart
                        values to apply, e.g. node-exporter-chart-values.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the values config
                        map, e.g. kube-system.
                      type: string
                    resourceVersion:
                      description: ResourceVersion is the Kubernetes resource version
                        of the configmap. Used to detect if the configmap has changed,
                        e.g. 12345.
                      type: string
                  required:
                  - name
                  - namespace
                  - resourceVersion
                  type: object
              required:
              - channel
              - configMap
              - name
              - namespace
              - release
              - secret
              - userConfigMap
              type: object
            versionBundle:
              properties:
                version:
                  type: string
              required:
              - version
              type: object
          required:
          - chart
          - versionBundle
          type: object
        status:
          properties:
            reason:
              description: Reason is the description of the last status of helm release
                when the chart is not installed successfully, e.g. deploy resource
                already exists.
              type: string
            releaseStatus:
              description: ReleaseStatus is the status of the Helm release when the
                chart is installed, e.g. DEPLOYED.
              type: string
          required:
          - releaseStatus
          type: object
      required:
      - metadata
      - spec
      - status
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

func NewChartConfigCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(chartconfigsYAML), &crd)
	return &crd
}
