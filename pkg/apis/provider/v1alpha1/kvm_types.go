package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
	"github.com/giantswarm/apiextensions/pkg/serialization"
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
// +kubebuilder:resource:categories=giantswarm;kvm
// +k8s:openapi-gen=true

type KVMConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KVMConfigSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status KVMConfigStatus `json:"status"`
}

// +k8s:openapi-gen=true
type KVMConfigSpec struct {
	Cluster       Cluster                    `json:"cluster"`
	KVM           KVMConfigSpecKVM           `json:"kvm"`
	VersionBundle KVMConfigSpecVersionBundle `json:"versionBundle"`
}

// +k8s:openapi-gen=true
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

// +k8s:openapi-gen=true
type KVMConfigSpecKVMEndpointUpdater struct {
	Docker KVMConfigSpecKVMEndpointUpdaterDocker `json:"docker"`
}

// +k8s:openapi-gen=true
type KVMConfigSpecKVMEndpointUpdaterDocker struct {
	Image string `json:"image"`
}

// +k8s:openapi-gen=true
type KVMConfigSpecKVMK8sKVM struct {
	Docker      KVMConfigSpecKVMK8sKVMDocker `json:"docker"`
	StorageType string                       `json:"storageType"`
}

// +k8s:openapi-gen=true
type KVMConfigSpecKVMK8sKVMDocker struct {
	Image string `json:"image"`
}

// +k8s:openapi-gen=true
type KVMConfigSpecKVMNode struct {
	CPUs               int                 `json:"cpus"`
	Disk               serialization.Float `json:"disk"`
	Memory             string              `json:"memory"`
	DockerVolumeSizeGB int                 `json:"dockerVolumeSizeGB"`
}

// +k8s:openapi-gen=true
type KVMConfigSpecKVMNetwork struct {
	Flannel KVMConfigSpecKVMNetworkFlannel `json:"flannel"`
}

// +k8s:openapi-gen=true
type KVMConfigSpecKVMNetworkFlannel struct {
	VNI int `json:"vni"`
}

// NOTE THIS IS DEPRECATED
// +k8s:openapi-gen=true
type KVMConfigSpecKVMNodeController struct {
	Docker KVMConfigSpecKVMNodeControllerDocker `json:"docker"`
}

// NOTE THIS IS DEPRECATED
// +k8s:openapi-gen=true
type KVMConfigSpecKVMNodeControllerDocker struct {
	Image string `json:"image"`
}

// +k8s:openapi-gen=true
type KVMConfigSpecKVMPortMappings struct {
	Name       string `json:"name"`
	NodePort   int    `json:"nodePort"`
	TargetPort int    `json:"targetPort"`
}

// +k8s:openapi-gen=true
type KVMConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type KVMConfigStatus struct {
	Cluster StatusCluster      `json:"cluster"`
	KVM     KVMConfigStatusKVM `json:"kvm"`
}

// +k8s:openapi-gen=true
type KVMConfigStatusKVM struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// NodeIndexes is a map from nodeID -> nodeIndex. This is used to create deterministic iSCSI initiator names.
	NodeIndexes map[string]int `json:"nodeIndexes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KVMConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KVMConfig `json:"items"`
}
