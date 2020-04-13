package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	kindApp              = "App"
	appDocumentationLink = "https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1?tab=doc#App"
)

const appCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: apps.application.giantswarm.io
spec:
  group: application.giantswarm.io
  scope: Namespaced
  version: v1alpha1
  names:
    kind: App
    plural: apps
    singular: app
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: |
        Defines an App resource, which represents an application to be running in a Kubernetes cluster.
        Reconciled by app-operator running on the Control Plane cluster.
      type: object
      properties:
        spec:
          type: object
          properties:
            catalog:
              description: |
                Name of the AppCatalog resource to install this app from. The AppCatalog resource must
                exist upfront in the cluster.
              type: string
            name:
              description: |
                Name of this App.
              type: string
            namespace:
              description: |
                Kubernetes namespace in which to install the workloads defined by this App in the 
                given cluster.
              type: string
            version:
              description: |
                Version of the app to be deployed. It must exist within the list of App packages
                stored in the App Catalog.
              type: string
            config:
              description: |
                Cluster configuration details for the application injected by cluster-operator with basic
                information of the cluster where the App is deployed. 
              type: object
              properties:
                configMap:
                  description: |
                    Defines a reference to a ConfigMap where is saved the cluster configuration values 
                    that will be applied to the application when it is deployed in the cluster. 
                  type: object
                  properties:
                    name:
                      description: |
                        Name of the ConfigMap resource.
                      type: string
                    namespace:
                      description: |
                        Namespace holding the ConfigMap resource.
                      type: string
                  required: ["name", "namespace"]
                secret:
                  description: |
                    Defines the reference of a Secret where is saved the sensitive configuration
                    that will be applied to the application when it is deployed in the cluster.
                  type: object
                  properties:
                    name:
                      description: |
                        Name of the Secret resource.
                      type: string
                    namespace:
                      description: |
                        Namespace holding the Secret resource.
                      type: string
                  required: ["name", "namespace"]
            kubeConfig:
              description: |
                The kubeconfig to use to connect to the cluster when deploying the app.
              type: object
              properties:
                inCluster:
                  description: |
                    Defines whether to use inCluster credentials. If true, it uses the service account to 
                    authenticate against the Kubernetes API in the same cluster where the app-operator
                    is running. In that case, the context and secret properties must not be set. If false,
                    secret and context is used by the app-operator to access the external cluster API.
                  type: boolean
                context:
                  description: |
                    Kubeconfig context part to use when not using inCluster credentials.
                  type: object
                  properties:
                    name:
                      description: |
                        Context name.
                      type: string
                secret:
                  description: |
                    Defines the reference to a Secret where is saved the kubeconfig configuration
                    that will be applied when accessing the cluster to manage the application.
                  type: object
                  properties:
                    name:
                      description: |
                        Name of the Secret resource.
                      type: string
                    namespace:
                      description: |
                        Namespace holding the Secret resource.
                      type: string
                  required: ["name", "namespace"]
            userConfig:
              description: |
                User configuration for the App. This configuration will be merged with the catalog and
                cluster configuration to generate a single values ConfigMap in the target cluster.
              type: object
              properties:
                configMap:
                  description: |
                    Defines the reference of a ConfigMap where is saved the user configuration values 
                    that will be applied to the application when it is deployed in the cluster. 
                  type: object
                  properties:
                    name:
                      description: |
                        Name of the ConfigMap resource.
                      type: string
                    namespace:
                      description: |
                        Namespace holding the ConfigMap resource.
                      type: string
                  required: ["name", "namespace"]
                secret:
                  description: |
                    Defines the reference of a Secret where is saved the user sensitive configuration 
                    that will be applied to the application when it is deployed in the cluster. 
                  type: object
                  properties:
                    name:
                      description: |
                        Name of the Secret resource.
                      type: string
                    namespace:
                      description: |
                        Namespace holding the Secret resource.
                      type: string
                  required: ["name", "namespace"]
          required: ["catalog", "name", "namespace", "version"]
`

var appCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(appCRDYAML), &appCRD)
	if err != nil {
		panic(err)
	}
}

// +gencrd

// NewAppCRD returns a new custom resource definition for App.
func NewAppCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return appCRD.DeepCopy()
}

func NewAppTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindApp,
	}
}

// NewAppCR returns an App Custom Resource.
func NewAppCR() *App {
	return &App{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: appDocumentationLink,
			},
		},
		TypeMeta: NewAppTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AppSpec   `json:"spec"`
	Status            AppStatus `json:"status" yaml:"status"`
}

type AppSpec struct {
	// Catalog is the name of the app catalog this app belongs to.
	// e.g. giantswarm
	Catalog string `json:"catalog" yaml:"catalog"`
	// Config is the config to be applied when the app is deployed.
	Config AppSpecConfig `json:"config" yaml:"config"`
	// KubeConfig is the kubeconfig to connect to the cluster when deploying
	// the app.
	KubeConfig AppSpecKubeConfig `json:"kubeConfig" yaml:"kubeConfig"`
	// Name is the name of the app to be deployed.
	// e.g. kubernetes-prometheus
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace where the app should be deployed.
	// e.g. monitoring
	Namespace string `json:"namespace" yaml:"namespace"`
	// UserConfig is the user config to be applied when the app is deployed.
	UserConfig AppSpecUserConfig `json:"userConfig" yaml:"userConfig"`
	// Version is the version of the app that should be deployed.
	// e.g. 1.0.0
	Version string `json:"version" yaml:"version"`
}

type AppSpecConfig struct {
	// ConfigMap references a config map containing values that should be
	// applied to the app.
	ConfigMap AppSpecConfigConfigMap `json:"configMap" yaml:"configMap"`
	// Secret references a secret containing secret values that should be
	// applied to the app.
	Secret AppSpecConfigSecret `json:"secret" yaml:"secret"`
}

type AppSpecConfigConfigMap struct {
	// Name is the name of the config map containing app values to apply,
	// e.g. prometheus-values.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the values config map,
	// e.g. monitoring.
	Namespace string `json:"namespace" yaml:"namespace"`
}

type AppSpecConfigSecret struct {
	// Name is the name of the secret containing app values to apply,
	// e.g. prometheus-secret.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the secret,
	// e.g. kube-system.
	Namespace string `json:"namespace" yaml:"namespace"`
}

type AppSpecKubeConfig struct {
	// InCluster is a flag for whether to use InCluster credentials. When true the
	// context name and secret should not be set.
	InCluster bool `json:"inCluster" yaml:"inCluster"`
	// Context is the kubeconfig context.
	Context AppSpecKubeConfigContext `json:"context" yaml:"context"`
	// Secret references a secret containing the kubconfig.
	Secret AppSpecKubeConfigSecret `json:"secret" yaml:"secret"`
}

type AppSpecKubeConfigContext struct {
	// Name is the name of the kubeconfig context.
	// e.g. giantswarm-12345.
	Name string `json:"name" yaml:"name"`
}

type AppSpecKubeConfigSecret struct {
	// Name is the name of the secret containing the kubeconfig,
	// e.g. app-operator-kubeconfig.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the secret containing the kubeconfig,
	// e.g. giantswarm.
	Namespace string `json:"namespace" yaml:"namespace"`
}

type AppSpecUserConfig struct {
	// ConfigMap references a config map containing user values that should be
	// applied to the app.
	ConfigMap AppSpecUserConfigConfigMap `json:"configMap" yaml:"configMap"`
	// Secret references a secret containing user secret values that should be
	// applied to the app.
	Secret AppSpecUserConfigSecret `json:"secret" yaml:"secret"`
}

type AppSpecUserConfigConfigMap struct {
	// Name is the name of the config map containing user values to apply,
	// e.g. prometheus-user-values.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the user values config map on the control plane,
	// e.g. 123ab.
	Namespace string `json:"namespace" yaml:"namespace"`
}

type AppSpecUserConfigSecret struct {
	// Name is the name of the secret containing user values to apply,
	// e.g. prometheus-user-secret.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the secret,
	// e.g. kube-system.
	Namespace string `json:"namespace" yaml:"namespace"`
}

type AppStatus struct {
	// AppVersion is the value of the AppVersion field in the Chart.yaml of the
	// deployed app. This is an optional field with the version of the
	// component being deployed.
	// e.g. 0.21.0.
	// https://docs.helm.sh/developing_charts/#the-chart-yaml-file
	AppVersion string `json:"appVersion" yaml:"appVersion"`
	// Release is the status of the Helm release for the deployed app.
	Release AppStatusRelease `json:"release" yaml:"release"`
	// Version is the value of the Version field in the Chart.yaml of the
	// deployed app.
	// e.g. 1.0.0.
	Version string `json:"version" yaml:"version"`
}

type AppStatusRelease struct {
	// LastDeployed is the time when the app was last deployed.
	LastDeployed metav1.Time `json:"lastDeployed" yaml:"lastDeployed"`
	// Reason is the description of the last status of helm release when the app is
	// not installed successfully, e.g. deploy resource already exists.
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
	// Status is the status of the deployed app,
	// e.g. DEPLOYED.
	Status string `json:"status" yaml:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []App `json:"items"`
}
