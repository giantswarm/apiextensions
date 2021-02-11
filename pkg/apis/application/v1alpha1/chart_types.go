package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/annotation"
	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

const (
	kindChart              = "Chart"
	chartDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/charts.application.giantswarm.io/"
)

func NewChartCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindChart)
}

func NewChartTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindChart,
	}
}

// NewChartCR returns an Chart Custom Resource.
func NewChartCR() *Chart {
	return &Chart{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotation.Docs: chartDocumentationLink,
			},
		},
		TypeMeta: NewChartTypeMeta(),
	}
}

// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.spec.version`,description="Version of the app"
// +kubebuilder:printcolumn:name="Last Deployed",type=date,JSONPath=`.status.release.lastDeployed`,description="Time since last deployment"
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.release.status`,description="Deployment status of the app"
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=common;giantswarm
// +k8s:openapi-gen=true
// Chart represents a Helm chart to be deployed as a Helm release. It is reconciled by chart-operator.
type Chart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ChartSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status ChartStatus `json:"status"`
}

// +k8s:openapi-gen=true
type ChartSpec struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// Config is the config to be applied when the chart is deployed.
	Config ChartSpecConfig `json:"config,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Install is the config used to deploy the app and is passed to Helm.
	Install ChartSpecInstall `json:"install,omitempty"`
	// Name is the name of the Helm chart to be deployed.
	// e.g. kubernetes-prometheus
	Name string `json:"name"`
	// Namespace is the namespace where the chart should be deployed.
	// e.g. monitoring
	Namespace string `json:"namespace"`
	// TarballURL is the URL for the Helm chart tarball to be deployed.
	// e.g. https://example.com/path/to/prom-1-0-0.tgz
	TarballURL string `json:"tarballURL"`
	// Version is the version of the chart that should be deployed.
	// e.g. 1.0.0
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type ChartSpecInstall struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// SkipCRDs when true decides that CRDs which are supplied with the chart are not installed. Default: false.
	SkipCRDs bool `json:"skipCRDs,omitempty"`
}

// +k8s:openapi-gen=true
type ChartSpecConfig struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// ConfigMap references a config map containing values that should be
	// applied to the chart.
	ConfigMap ChartSpecConfigConfigMap `json:"configMap,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Secret references a secret containing secret values that should be
	// applied to the chart.
	Secret ChartSpecConfigSecret `json:"secret,omitempty"`
}

// +k8s:openapi-gen=true
type ChartSpecConfigConfigMap struct {
	// Name is the name of the config map containing chart values to apply,
	// e.g. prometheus-chart-values.
	Name string `json:"name"`
	// Namespace is the namespace of the values config map,
	// e.g. monitoring.
	Namespace string `json:"namespace"`
	// ResourceVersion is the Kubernetes resource version of the configmap.
	// Used to detect if the configmap has changed, e.g. 12345.
	ResourceVersion string `json:"resourceVersion"`
}

// +k8s:openapi-gen=true
type ChartSpecConfigSecret struct {
	// Name is the name of the secret containing chart values to apply,
	// e.g. prometheus-chart-secret.
	Name string `json:"name"`
	// Namespace is the namespace of the secret,
	// e.g. kube-system.
	Namespace string `json:"namespace"`
	// ResourceVersion is the Kubernetes resource version of the secret.
	// Used to detect if the secret has changed, e.g. 12345.
	ResourceVersion string `json:"resourceVersion"`
}

// +k8s:openapi-gen=true
type ChartStatus struct {
	// AppVersion is the value of the AppVersion field in the Chart.yaml of the
	// deployed chart. This is an optional field with the version of the
	// component being deployed.
	// e.g. 0.21.0.
	// https://helm.sh/docs/topics/charts/#the-chartyaml-file
	AppVersion string `json:"appVersion"`
	// Reason is the description of the last status of helm release when the chart is
	// not installed successfully, e.g. deploy resource already exists.
	Reason string `json:"reason,omitempty"`
	// Release is the status of the Helm release for the deployed chart.
	Release ChartStatusRelease `json:"release"`
	// Version is the value of the Version field in the Chart.yaml of the
	// deployed chart.
	// e.g. 1.0.0.
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type ChartStatusRelease struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// LastDeployed is the time when the deployed chart was last deployed.
	LastDeployed *metav1.Time `json:"lastDeployed,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Revision is the revision number for this deployed chart.
	Revision *int `json:"revision,omitempty"`
	// Status is the status of the deployed chart,
	// e.g. DEPLOYED.
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Chart `json:"items"`
}
