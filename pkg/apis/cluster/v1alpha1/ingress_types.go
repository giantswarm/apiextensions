package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Ingress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              IngressSpec `json:"spec"`
}

type IngressSpec struct {
	GuestCluster  IngressSpecGuestCluster   `json:"guestcluster" yaml:"guestcluster"`
	HostCluster   IngressSpecHostCluster    `json:"hostcluster" yaml:"hostcluster"`
	ProtocolPorts []IngressSpecProtocolPort `json:"protocolPorts" yaml:"protocolPorts"`
}

type IngressSpecGuestCluster struct {
	ID        string `json:"id" yaml:"id"`
	Namespace string `json:"namespace" yaml:"namespace"`
	Service   string `json:"service" yaml:"service"`
}

type IngressSpecHostCluster struct {
	IngressController IngressSpecHostClusterIngressController `json:"ingressController" yaml:"ingressController"`
}

type IngressSpecHostClusterIngressController struct {
	ConfigMap string `json:"configMap" yaml:"configMap"`
	Namespace string `json:"namespace" yaml:"namespace"`
	Service   string `json:"service" yaml:"service"`
}

type IngressSpecProtocolPort struct {
	IngressPort int    `json:"ingressPort" yaml:"ingressPort"`
	LBPort      int    `json:"lbPort" yaml:"lbPort"`
	Protocol    string `json:"protocol" yaml:"protocol"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IngressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Ingress `json:"items"`
}
