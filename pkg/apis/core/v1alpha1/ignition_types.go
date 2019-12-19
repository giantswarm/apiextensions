package v1alpha1

import (
	"fmt"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindIgnition = "Ignition"
)

// NewIgnitionCRD returns a new custom resource definition for Ignition. This
// might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: ignitions.core.giantswarm.io
//     spec:
//       group: core.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: Ignition
//         plural: ignitions
//         singular: ignition
//       subresources:
//         status: {}
//
func NewIgnitionCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("ignitions.%s", group),
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   group,
			Scope:   "Namespaced",
			Version: version,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     kindIgnition,
				Plural:   "ignitions",
				Singular: "ignition",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

func NewIgnitionTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindIgnition,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Ignition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              IgnitionSpec   `json:"spec"`
	Status            IgnitionStatus `json:"status"`
}

// IgnitionSpec is the interface which defines the input parameters for
// a newly rendered g8s ignition template.
type IgnitionSpec struct {
	BaseDomain string                 `json:"basedomain" yaml:"basedomain"`
	Calico     IgnitionSpecCalico     `json:"calico" yaml:"calico"`
	ClusterID  string                 `json:"clusterid" yaml:"clusterid"`
	Etcd       IgnitionSpecEtcd       `json:"etcd" yaml:"etcd"`
	Ingress    IgnitionSpecIngress    `json:"ingress" yaml:"ingress"`
	Kubernetes IgnitionSpecKubernetes `json:"kubernetes" yaml:"kubernetes"`
	// Defines the provider which should be rendered.
	Provider string               `json:"provider" yaml:"provider"`
	Registry IgnitionSpecRegistry `json:"registry" yaml:"registry"`
	SSO      IgnitionSpecSSO      `json:"sso" yaml:"sso"`
}

type IgnitionSpecCalico struct {
	CIDR    string `json:"cidr" yaml:"cidr"`
	Disable bool   `json:"disable" yaml:"disable"`
	MTU     string `json:"mtu" yaml:"mtu"`
	Subnet  string `json:"subnet" yaml:"subnet"`
}

type IgnitionSpecEtcd struct {
	Domain string `json:"domain" yaml:"domain"`
	Port   int    `json:"port" yaml:"port"`
	Prefix string `json:"prefix" yaml:"prefix"`
}

type IgnitionSpecIngress struct {
	Disable bool `json:"disable" yaml:"disable"`
}

type IgnitionSpecKubernetes struct {
	API     IgnitionSpecKubernetesAPI     `json:"api" yaml:"api"`
	DNS     IgnitionSpecKubernetesDNS     `json:"dns" yaml:"dns"`
	Domain  string                        `json:"domain" yaml:"domain"`
	Kubelet IgnitionSpecKubernetesKubelet `json:"kubelet" yaml:"kubelet"`
	Image   string                        `json:"image" yaml:"image"`
	IPRange string                        `json:"iprange" yaml:"iprange"`
}

type IgnitionSpecKubernetesAPI struct {
	Domain     string `json:"domain" yaml:"domain"`
	SecurePort int    `json:"secureport" yaml:"secureport"`
}

type IgnitionSpecKubernetesDNS struct {
	IP string `json:"ip" yaml:"ip"`
}

type IgnitionSpecKubernetesKubelet struct {
	Domain string `json:"domain" yaml:"domain"`
}

type IgnitionSpecRegistry struct {
	Domain string `json:"domain" yaml:"domain"`
}
type IgnitionSpecSSO struct {
	PublicKey string `json:"publicKey" yaml:"publicKey"`
}

// IgnitionStatus holds the rendering result.
type IgnitionStatus struct {
	ConfigMap IgnitionStatusConfigMap `json:"template" yaml:"template"`
}

type IgnitionStatusConfigMap struct {
	// Name is the name of the config map containing the rendered ignition.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the config map containing the rendered ignition.
	Namespace string `json:"namespace" yaml:"namespace"`
	// ResourceVersion is the Kubernetes resource version of the configmap.
	// Used to detect if the configmap has changed, e.g. 12345.
	ResourceVersion string `json:"resourceVersion" yaml:"resourceVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IgnitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Ignition `json:"items"`
}
