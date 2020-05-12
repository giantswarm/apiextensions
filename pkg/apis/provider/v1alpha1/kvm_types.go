package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindKVMConfig = "KVMConfig"
)

func NewKVMConfigCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindKVMConfig)
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

type KVMConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KVMConfigSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status KVMConfigStatus `json:"status"`
}

type KVMConfigSpec struct {
	Cluster       Cluster                    `json:"cluster"`
	KVM           KVMConfigSpecKVM           `json:"kvm"`
	VersionBundle KVMConfigSpecVersionBundle `json:"versionBundle"`
}

type KVMConfigSpecKVM struct {
	EndpointUpdater KVMConfigSpecKVMEndpointUpdater `json:"endpointUpdater"`
	K8sKVM          KVMConfigSpecKVMK8sKVM          `json:"k8sKVM" `
	Masters         []KVMConfigSpecKVMNode          `json:"masters"`
	Network         KVMConfigSpecKVMNetwork         `json:"network"`
	// NOTE THIS IS DEPRECATED
	NodeController KVMConfigSpecKVMNodeController `json:"nodeController"`
	PortMappings   []KVMConfigSpecKVMPortMappings `json:"portMappings"`
	Workers        []KVMConfigSpecKVMNode         `json:"workers"`
}

type KVMConfigSpecKVMEndpointUpdater struct {
	Docker KVMConfigSpecKVMEndpointUpdaterDocker `json:"docker"`
}

type KVMConfigSpecKVMEndpointUpdaterDocker struct {
	Image string `json:"image"`
}

type KVMConfigSpecKVMK8sKVM struct {
	Docker      KVMConfigSpecKVMK8sKVMDocker `json:"docker"`
	StorageType string                       `json:"storageType"`
}

type KVMConfigSpecKVMK8sKVMDocker struct {
	Image string `json:"image"`
}

type KVMConfigSpecKVMNode struct {
	CPUs int `json:"cpus"`
	// +kubebuilder:validation:Type=number
	Disk               string `json:"disk"`
	Memory             string `json:"memory"`
	DockerVolumeSizeGB int    `json:"dockerVolumeSizeGB"`
}

type KVMConfigSpecKVMNetwork struct {
	Flannel KVMConfigSpecKVMNetworkFlannel `json:"flannel"`
}

type KVMConfigSpecKVMNetworkFlannel struct {
	VNI int `json:"vni"`
}

// NOTE THIS IS DEPRECATED
type KVMConfigSpecKVMNodeController struct {
	Docker KVMConfigSpecKVMNodeControllerDocker `json:"docker"`
}

// NOTE THIS IS DEPRECATED
type KVMConfigSpecKVMNodeControllerDocker struct {
	Image string `json:"image"`
}

type KVMConfigSpecKVMPortMappings struct {
	Name       string `json:"name"`
	NodePort   int    `json:"nodePort"`
	TargetPort int    `json:"targetPort"`
}

type KVMConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

type KVMConfigStatus struct {
	Cluster StatusCluster      `json:"cluster"`
	KVM     KVMConfigStatusKVM `json:"kvm"`
}

type KVMConfigStatusKVM struct {
	// NodeIndexes is a map from nodeID -> nodeIndex. This is used to create deterministic iSCSI initiator names.
	NodeIndexes map[string]int `json:"nodeIndexes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KVMConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KVMConfig `json:"items"`
}
