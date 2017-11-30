package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewFlannelCRD returns a new custom resource definition for Flannel. This
// might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: flannels.core.giantswarm.io
//     spec:
//       group: core.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: Flannel
//         plural: flannels
//         singular: flannel
//
func NewFlannelCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "flannels.core.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "core.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "Flannel",
				Plural:   "flannels",
				Singular: "flannel",
			},
		},
	}
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Flannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              FlannelSpec `json:"spec"`
}

type FlannelSpec struct {
	Bridge        FlannelSpecBridge        `json:"bridge" yaml:"bridge"`
	Cluster       FlannelSpecCluster       `json:"cluster" yaml:"cluster"`
	Flannel       FlannelSpecFlannel       `json:"flannel" yaml:"flannel"`
	Health        FlannelSpecHealth        `json:"health" yaml:"health"`
	VersionBundle FlannelSpecVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type FlannelSpecBridge struct {
	Docker FlannelSpecBridgeDocker `json:"docker" yaml:"docker"`
	Spec   FlannelSpecBridgeSpec   `json:"spec" yaml:"spec"`
}

type FlannelSpecBridgeDocker struct {
	Image string `json:"image" yaml:"image"`
}

type FlannelSpecBridgeSpec struct {
	Interface      string                   `json:"interface" yaml:"interface"`
	PrivateNetwork string                   `json:"privateNetwork" yaml:"privateNetwork"`
	DNS            FlannelSpecBridgeSpecDNS `json:"dns" yaml:"dns"`
	NTP            FlannelSpecBridgeSpecNTP `json:"ntp" yaml:"ntp"`
}

type FlannelSpecBridgeSpecDNS struct {
	Servers []string `json:"servers" yaml:"servers"`
}

type FlannelSpecBridgeSpecNTP struct {
	Servers []string `json:"servers" yaml:"servers"`
}

type FlannelSpecCluster struct {
	ID        string `json:"id" yaml:"id"`
	Customer  string `json:"customer" yaml:"customer"`
	Namespace string `json:"namespace" yaml:"namespace"`
}

type FlannelSpecFlannel struct {
	Spec FlannelSpecFlannelSpec `json:"spec" yaml:"spec"`
}

type FlannelSpecFlannelSpec struct {
	Network   string `json:"network" yaml:"network"`
	SubnetLen int    `json:"subnetLen" yaml:"subnetLen"`
	RunDir    string `json:"runDir" yaml:"runDir"`
	VNI       int    `json:"vni" yaml:"vni"`
}

type FlannelSpecHealth struct {
	Docker FlannelSpecHealthDocker `json:"docker" yaml:"docker"`
}

type FlannelSpecHealthDocker struct {
	Image string `json:"image" yaml:"image"`
}

type FlannelSpecVersionBundle struct {
	Version string `json:"version" yaml:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FlannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Flannel `json:"items"`
}
