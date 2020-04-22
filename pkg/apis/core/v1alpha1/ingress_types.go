package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindIngressConfig = "IngressConfig"
)

func NewIngressConfigCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindIngressConfig)
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IngressConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              IngressConfigSpec `json:"spec"`
}

type IngressConfigSpec struct {
	GuestCluster  IngressConfigSpecGuestCluster   `json:"guestCluster"`
	HostCluster   IngressConfigSpecHostCluster    `json:"hostCluster"`
	ProtocolPorts []IngressConfigSpecProtocolPort `json:"protocolPorts"`
	VersionBundle IngressConfigSpecVersionBundle  `json:"versionBundle"`
}

type IngressConfigSpecGuestCluster struct {
	ID        string `json:"id"`
	Namespace string `json:"namespace"`
	Service   string `json:"service"`
}

type IngressConfigSpecHostCluster struct {
	IngressController IngressConfigSpecHostClusterIngressController `json:"ingressController"`
}

type IngressConfigSpecHostClusterIngressController struct {
	ConfigMap string `json:"configMap"`
	Namespace string `json:"namespace"`
	Service   string `json:"service"`
}

type IngressConfigSpecProtocolPort struct {
	IngressPort int    `json:"ingressPort"`
	LBPort      int    `json:"lbPort"`
	Protocol    string `json:"protocol"`
}

type IngressConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IngressConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []IngressConfig `json:"items"`
}
