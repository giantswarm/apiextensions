package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindAppDeployment = "AppDeployment"
)

// NewAppDeploymentCRD returns a new custom resource definition for AppDeployment.
// This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: appdeployment.application.giantswarm.io
//     spec:
//       group: application.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: AppDeployment
//         plural: appdeployments
//         singular: appdeployment
//
func NewAppDeploymentCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "appdeployments.application.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "application.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "AppDeployment",
				Plural:   "appdeployments",
				Singular: "appdeployment",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

func NewAppDeploymentTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindAppDeployment,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppDeployment CR example as below.
//
//     apiVersion: application.giantswarm.io/v1alpha1
//     kind: AppDeployment
//     metadata:
//       name: “My-Cool-Prometheus”
//       namespace: “12345”
//     spec:
//       catalog: "giant-swarm"
//       application: “kubernetes-prometheus”
//       release: 1.0.0
//       kubeContext: “giantswarm-12345”
//       namespace: “monitoring”
//     status:
//       releaseStatus: “DEPLOYED”
//
type AppDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AppDeploymentSpec `json:"spec"`
}

type AppDeploymentSpec struct {
	// Catalog is the name of the application catalog this deployment belongs to.
	// e.g. giant-swarm
	Catalog string `json:"catalog" yaml:"catalog"`
	// App is the name of the application to be deployed.
	// e.g. kubernetes-prometheus
	App string `json:"application" yaml:"application"`
	// Release is the version of this application which we would like to use.
	// e.g. 1.0.0
	Release string `json:"release" yaml:"release"`
	// KubeContext is the context name for the Kubernetes cluster where the application should be deployed.
	// e.g. giantswarm-12345
	KubeContext string `json:"kubeContext" yaml:"kubeContext"`
	// Namespace is the namespace where the application should be deployed.
	// e.g. monitoring
	Namespace string                  `json:"namespace" yaml:"namespace"`
	Status    AppDeploymentSpecStatus `json:"status" yaml:"status"`
}

type AppDeploymentSpecStatus struct {
	ReleaseStatus string `json:"releaseStatus" yaml:"releaseStatus"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AppDeployment `json:"items"`
}
