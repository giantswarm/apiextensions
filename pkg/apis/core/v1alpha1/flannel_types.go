package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindFlannelConfig = "FlannelConfig"
)

func NewFlannelConfigCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindFlannelConfig)
}

// +genclient
// +genclient:noStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=giantswarm;kvm
// +kubebuilder:storageversion
type FlannelConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              FlannelConfigSpec `json:"spec"`
}

type FlannelConfigSpec struct {
	Bridge        FlannelConfigSpecBridge        `json:"bridge"`
	Cluster       FlannelConfigSpecCluster       `json:"cluster"`
	Flannel       FlannelConfigSpecFlannel       `json:"flannel"`
	Health        FlannelConfigSpecHealth        `json:"health"`
	VersionBundle FlannelConfigSpecVersionBundle `json:"versionBundle"`
}

type FlannelConfigSpecBridge struct {
	Docker FlannelConfigSpecBridgeDocker `json:"docker"`
	Spec   FlannelConfigSpecBridgeSpec   `json:"spec"`
}

type FlannelConfigSpecBridgeDocker struct {
	Image string `json:"image"`
}

type FlannelConfigSpecBridgeSpec struct {
	Interface      string                         `json:"interface"`
	PrivateNetwork string                         `json:"privateNetwork"`
	DNS            FlannelConfigSpecBridgeSpecDNS `json:"dns"`
	NTP            FlannelConfigSpecBridgeSpecNTP `json:"ntp"`
}

type FlannelConfigSpecBridgeSpecDNS struct {
	Servers []string `json:"servers"`
}

type FlannelConfigSpecBridgeSpecNTP struct {
	Servers []string `json:"servers"`
}

type FlannelConfigSpecCluster struct {
	ID        string `json:"id"`
	Customer  string `json:"customer"`
	Namespace string `json:"namespace"`
}

type FlannelConfigSpecFlannel struct {
	Spec FlannelConfigSpecFlannelSpec `json:"spec"`
}

type FlannelConfigSpecFlannelSpec struct {
	Network   string `json:"network"`
	SubnetLen int    `json:"subnetLen"`
	RunDir    string `json:"runDir"`
	VNI       int    `json:"vni"`
}

type FlannelConfigSpecHealth struct {
	Docker FlannelConfigSpecHealthDocker `json:"docker"`
}

type FlannelConfigSpecHealthDocker struct {
	Image string `json:"image"`
}

type FlannelConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FlannelConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []FlannelConfig `json:"items"`
}
