package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/apiextensions/v3/pkg/serialization"
)

const (
	kindKVMCluster = "KVMCluster"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=giantswarm;kvm
// +k8s:openapi-gen=true

type KVMCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KVMClusterSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status KVMClusterStatus `json:"status"`
}

// +k8s:openapi-gen=true
type KVMClusterSpec struct {
	Cluster       v1alpha1.Cluster            `json:"cluster"`
	KVM           KVMClusterSpecKVM           `json:"kvm"`
	VersionBundle KVMClusterSpecVersionBundle `json:"versionBundle"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecKVM struct {
	EndpointUpdater KVMClusterSpecKVMEndpointUpdater `json:"endpointUpdater"`
	K8sKVM          KVMClusterSpecKVMK8sKVM          `json:"k8sKVM" `
	Masters         []KVMClusterSpecKVMNode          `json:"masters"`
	Network         KVMClusterSpecKVMNetwork         `json:"network"`
	// NOTE THIS IS DEPRECATED
	NodeController KVMClusterSpecKVMNodeController `json:"nodeController"`
	PortMappings   []KVMClusterSpecKVMPortMappings `json:"portMappings"`
	Workers        []KVMClusterSpecKVMNode         `json:"workers"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecKVMEndpointUpdater struct {
	Docker KVMClusterSpecKVMEndpointUpdaterDocker `json:"docker"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecKVMEndpointUpdaterDocker struct {
	Image string `json:"image"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecKVMK8sKVM struct {
	Docker      KVMClusterSpecKVMK8sKVMDocker `json:"docker"`
	StorageType string                        `json:"storageType"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecKVMK8sKVMDocker struct {
	Image string `json:"image"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecKVMNode struct {
	CPUs               int                 `json:"cpus"`
	Disk               serialization.Float `json:"disk"`
	Memory             string              `json:"memory"`
	DockerVolumeSizeGB int                 `json:"dockerVolumeSizeGB"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecKVMNetwork struct {
	Flannel KVMClusterSpecKVMNetworkFlannel `json:"flannel"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecKVMNetworkFlannel struct {
	VNI int `json:"vni"`
}

// NOTE THIS IS DEPRECATED
// +k8s:openapi-gen=true
type KVMClusterSpecKVMNodeController struct {
	Docker KVMClusterSpecKVMNodeControllerDocker `json:"docker"`
}

// NOTE THIS IS DEPRECATED
// +k8s:openapi-gen=true
type KVMClusterSpecKVMNodeControllerDocker struct {
	Image string `json:"image"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecKVMPortMappings struct {
	Name       string `json:"name"`
	NodePort   int    `json:"nodePort"`
	TargetPort int    `json:"targetPort"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type KVMClusterStatus struct {
	Cluster v1alpha1.StatusCluster `json:"cluster"`
	KVM     KVMClusterStatusKVM    `json:"kvm"`
}

// +k8s:openapi-gen=true
type KVMClusterStatusKVM struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// NodeIndexes is a map from nodeID -> nodeIndex. This is used to create deterministic iSCSI initiator names.
	NodeIndexes map[string]int `json:"nodeIndexes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KVMClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KVMCluster `json:"items"`
}
