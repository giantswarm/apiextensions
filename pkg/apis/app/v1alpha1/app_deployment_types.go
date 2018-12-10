package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewAppDeploymentCRD returns a new custom resource definition for AppDeployment.
// This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: appdeployment.app.giantswarm.io
//     spec:
//       group: app.giantswarm.io
//       scope: Cluster
//       version: v1alpha1
//       names:
//         kind: AppDeployment
//         plural: appdeployments
//         singular: appdeployments
//
func NewAppDeploymentCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "appdeployments.app.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "app.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "AppDeployment",
				Plural:   "appdeployments",
				Singular: "appdeployment",
			},
		},
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AppDeploymentSpec `json:"spec"`
}

type AppDeploymentSpec struct {
	// Catalog is the name of the app deployment for this CR
	// e.g. giant-swarm
	Catalog string `json:"catalog" yaml:"catalog"`
	App     string `json:"app" yaml:"app"`
	// Release is the version of this app which we would like to use.
	Release string `json:"release" yaml:"release"`
	// KubeContext is the tenant cluster-based context name which point to specific kubeConfig as well.
	KubeContext string `json:"kubeContext" yaml:"kubeContext"`
	// Namespace is the tenant cluster-based namespace where this app would be eventually located.
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
