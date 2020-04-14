package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

func NewIngressConfigCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadCRD(group, kindCluster)
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
	GuestCluster  IngressConfigSpecGuestCluster   `json:"guestCluster" yaml:"guestCluster"`
	HostCluster   IngressConfigSpecHostCluster    `json:"hostCluster" yaml:"hostCluster"`
	ProtocolPorts []IngressConfigSpecProtocolPort `json:"protocolPorts" yaml:"protocolPorts"`
	VersionBundle IngressConfigSpecVersionBundle  `json:"versionBundle" yaml:"versionBundle"`
}

type IngressConfigSpecGuestCluster struct {
	ID        string `json:"id" yaml:"id"`
	Namespace string `json:"namespace" yaml:"namespace"`
	Service   string `json:"service" yaml:"service"`
}

type IngressConfigSpecHostCluster struct {
	IngressController IngressConfigSpecHostClusterIngressController `json:"ingressController" yaml:"ingressController"`
}

type IngressConfigSpecHostClusterIngressController struct {
	ConfigMap string `json:"configMap" yaml:"configMap"`
	Namespace string `json:"namespace" yaml:"namespace"`
	Service   string `json:"service" yaml:"service"`
}

type IngressConfigSpecProtocolPort struct {
	IngressPort int    `json:"ingressPort" yaml:"ingressPort"`
	LBPort      int    `json:"lbPort" yaml:"lbPort"`
	Protocol    string `json:"protocol" yaml:"protocol"`
}

type IngressConfigSpecVersionBundle struct {
	Version string `json:"version" yaml:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IngressConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []IngressConfig `json:"items"`
}
