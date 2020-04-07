package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const crdYAML = `apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  name: apps.application.giantswarm.io
spec:
  group: application.giantswarm.io
  names:
    kind: App
    plural: apps
    singular: app
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: |
        Defines an App resource, which represents an application to be running in a Kubernetes cluster.
        Reconciled by app-operator.
      properties:
        spec:
          properties:
            catalog:
              description: |
                Name of the AppCatalog to install this app from. Find more information in the AppCatalog
                CRD documentation.
              type: string
            config:
              description: |
                Configuration details for the app.
              properties:
                configMap:
                  description: |
                    If present, points to a ConfigMap resource that holds configuration data
                    used by the app.
                  properties:
                    name:
                      description: |
                        Name of the ConfigMap.
                      type: string
                    namespace:
                      description: |
                        Namespace to find the ConfigMap in.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
                secret:
                  description: |
                    If present, points to a Secret resoure that can be used by the app.
                  properties:
                    name:
                      description: |
                        Name of the Secret.
                      type: string
                    namespace:
                      description: |
                        Namespace to find the Secret in.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
              type: object
            kubeConfig:
              description: |
                The kubeconfig to use to connect to the tenant cluster when deploying the app.
              properties:
                context:
                  description: |
                    Kubeconfig context part to use when not using inCluster credentials.
                  properties:
                    name:
                      description: |
                        Context name.
                      type: string
                  type: object
                inCluster:
                  description: |
                    Defines whether to use inCluster credentials. If true, the context and secret
                    properties must not be set.
                  type: boolean
                secret:
                  description: |
                    References a Secret resource holding the kubeconfig details, if not using inCluster credentials.
                  properties:
                    name:
                      description: |
                        Name of the Secret resource.
                      type: string
                    namespace:
                      description: |
                        Namespace holding the Secret resource.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
              type: object
            name:
              description: |
                Name of this App.
              type: string
            namespace:
              description: |
                Kubernetes namespace in which to install the workloads defined by this App.
              type: string
            userConfig:
              description: |
                Additional and optional user-provided configuration for the app.
              properties:
                configMap:
                  description: |
                    Reference to an optional ConfigMap.
                  properties:
                    name:
                      description: |
                        Name of the ConfigMap resource.
                      type: string
                    namespace:
                      description: |
                        Namespace holding the ConfigMap resource.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
                secret:
                  description: |
                    Reference to an optional Secret resource.
                  properties:
                    name:
                      description: |
                        Name of the Secret resource.
                      type: string
                    namespace:
                      description: |
                        Namespace holding the Secret resource.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
              type: object
            version:
              description: Version of the app to be deployed.
              type: string
          required:
          - catalog
          - name
          - namespace
          - version
          type: object
      type: object
  version: v1alpha1`

func NewApplicationCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(crdYAML), &crd)
	return &crd
}
