package application

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const appsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: apps.application.giantswarm.io
spec:
  group: application.giantswarm.io
  names:
    kind: App
    listKind: AppList
    plural: apps
    singular: app
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
            catalog:
              description: Catalog is the name of the app catalog this app belongs
                to. e.g. giantswarm
              type: string
            config:
              description: Config is the config to be applied when the app is deployed.
              properties:
                configMap:
                  description: ConfigMap references a config map containing values
                    that should be applied to the app.
                  properties:
                    name:
                      description: Name is the name of the config map containing app
                        values to apply, e.g. prometheus-values.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the values config
                        map, e.g. monitoring.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
                secret:
                  description: Secret references a secret containing secret values
                    that should be applied to the app.
                  properties:
                    name:
                      description: Name is the name of the secret containing app values
                        to apply, e.g. prometheus-secret.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the secret, e.g.
                        kube-system.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
              required:
              - configMap
              - secret
              type: object
            kubeConfig:
              description: KubeConfig is the kubeconfig to connect to the cluster
                when deploying the app.
              properties:
                context:
                  description: Context is the kubeconfig context.
                  properties:
                    name:
                      description: Name is the name of the kubeconfig context. e.g.
                        giantswarm-12345.
                      type: string
                  required:
                  - name
                  type: object
                inCluster:
                  description: InCluster is a flag for whether to use InCluster credentials.
                    When true the context name and secret should not be set.
                  type: boolean
                secret:
                  description: Secret references a secret containing the kubconfig.
                  properties:
                    name:
                      description: Name is the name of the secret containing the kubeconfig,
                        e.g. app-operator-kubeconfig.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the secret containing
                        the kubeconfig, e.g. giantswarm.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
              required:
              - context
              - inCluster
              - secret
              type: object
            name:
              description: Name is the name of the app to be deployed. e.g. kubernetes-prometheus
              type: string
            namespace:
              description: Namespace is the namespace where the app should be deployed.
                e.g. monitoring
              type: string
            userConfig:
              description: UserConfig is the user config to be applied when the app
                is deployed.
              properties:
                configMap:
                  description: ConfigMap references a config map containing user values
                    that should be applied to the app.
                  properties:
                    name:
                      description: Name is the name of the config map containing user
                        values to apply, e.g. prometheus-user-values.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the user values config
                        map on the control plane, e.g. 123ab.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
                secret:
                  description: Secret references a secret containing user secret values
                    that should be applied to the app.
                  properties:
                    name:
                      description: Name is the name of the secret containing user
                        values to apply, e.g. prometheus-user-secret.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the secret, e.g.
                        kube-system.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
              required:
              - configMap
              - secret
              type: object
            version:
              description: Version is the version of the app that should be deployed.
                e.g. 1.0.0
              type: string
          required:
          - catalog
          - config
          - kubeConfig
          - name
          - namespace
          - userConfig
          - version
          type: object
        status:
          properties:
            appVersion:
              description: AppVersion is the value of the AppVersion field in the
                Chart.yaml of the deployed app. This is an optional field with the
                version of the component being deployed. e.g. 0.21.0. https://docs.helm.sh/developing_charts/#the-chart-yaml-file
              type: string
            release:
              description: Release is the status of the Helm release for the deployed
                app.
              properties:
                lastDeployed:
                  description: LastDeployed is the time when the app was last deployed.
                  format: date-time
                  type: string
                reason:
                  description: Reason is the description of the last status of helm
                    release when the app is not installed successfully, e.g. deploy
                    resource already exists.
                  type: string
                status:
                  description: Status is the status of the deployed app, e.g. DEPLOYED.
                  type: string
              required:
              - lastDeployed
              - status
              type: object
            version:
              description: Version is the value of the Version field in the Chart.yaml
                of the deployed app. e.g. 1.0.0.
              type: string
          required:
          - appVersion
          - release
          - version
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

func NewAppCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(appsYAML), &crd)
	return &crd
}
