package v1alpha3

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindAzureMachinePool = "AzureMachinePool"
)

func NewAzureMachinePoolCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindAzureMachinePool)
}

func NewAzureMachinePoolTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAzureMachinePool,
	}
}

// NewAzureMachinePoolCR returns an AzureMachinePool Custom Resource.
func NewAzureMachinePoolCR() *AzureMachinePool {
	return &AzureMachinePool{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: azureClusterDocumentationLink,
			},
		},
		TypeMeta: NewAzureMachinePoolTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status

// AzureMachinePool is the infrastructure provider referenced in Kubernetes Cluster API MachinePool resources.
// It contains provider-specific specification and status for a node pool.
// In use on Azure since Giant Swarm release v10.x.x and reconciled by azure-operator.
type AzureMachinePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Contains the specification.
	Spec AzureMachinePoolSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Holds status information.
	Status AzureMachinePoolStatus `json:"status"`
}

type AzureMachinePoolSpec struct {
	// Specifies details of node pool and the worker nodes it should contain.
	NodePool AzureMachinePoolSpecNodePool `json:"nodePool"`
	// Contains Azure specific details.
	Provider AzureMachinePoolSpecProvider `json:"provider"`
}

type AzureMachinePoolSpecNodePool struct {
	// User-friendly name or description of the purpose of the node pool.
	Description string `json:"description"`
	// Specification of the worker node machine.
	Machine AzureMachinePoolSpecNodePoolMachine `json:"machine"`
	// Scaling settings for the node pool, configuring the cluster-autoscaler
	// determining the number of nodes to have in this node pool.
	Scaling AzureMachinePoolSpecNodePoolScaling `json:"scaling"`
}

type AzureMachinePoolSpecNodePoolMachine struct {
	// Size of the volume reserved for Docker images and overlay file systems of
	// Docker containers. Unit: 1 GB = 1,000,000,000 Bytes.
	DockerVolumeSizeGB int `json:"dockerVolumeSizeGB"`
	// Size of the volume reserved for the kubelet, which can be used by Pods via
	// volumes of type EmptyDir. Unit: 1 GB = 1,000,000,000 Bytes.
	KubeletVolumeSizeGB int `json:"kubeletVolumeSizeGB"`
}

type AzureMachinePoolSpecNodePoolScaling struct {
	// Maximum number of worker nodes in this node pool.
	Max int `json:"max"`
	// Minimum number of worker nodes in this node pool.
	Min int `json:"min"`
}

type AzureMachinePoolSpecProvider struct {
	// Name(s) of the availability zone(s) to use for worker nodes. Using multiple
	// availability zones results in higher resilience but can also result in higher
	// cost due to network traffic between availability zones.
	AvailabilityZones []string `json:"availabilityZones"`
	// +kubebuilder:validation:Optional
	// Settings defining the distribution of on-demand and spot instances in the node pool.
	InstanceDistribution AzureMachinePoolSpecInstanceDistribution `json:"instanceDistribution"`
	// Specification of worker nodes.
	Worker AzureMachinePoolSpecProviderWorker `json:"worker"`
}

type AzureMachinePoolSpecInstanceDistribution struct {
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	// Base capacity of on-demand instances to use for worker nodes in this pool. When this larger
	// than 0, this value defines a number of worker nodes that will be created using on-demand
	// EC2 instances, regardless of the value configured as `onDemandPercentageAboveBaseCapacity`.
	OnDemandBaseCapacity int `json:"onDemandBaseCapacity"`
	// +kubebuilder:default=100
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Minimum=0
	// Percentage of on-demand EC2 instances to use for worker nodes, instead of spot instances,
	// for instances exceeding `onDemandBaseCapacity`. For example, to have half of the worker nodes
	// use spot instances and half use on-demand, set this value to 50.
	OnDemandPercentageAboveBaseCapacity int `json:"onDemandPercentageAboveBaseCapacity"`
}

type AzureMachinePoolSpecProviderWorker struct {
	// Azure EC2 instance type name to use for the worker nodes in this node pool.
	InstanceType string `json:"instanceType"`
	// +kubebuilder:default=false
	// If true, certain instance types with specs similar to instanceType will be used.
	UseAlikeInstanceTypes bool `json:"useAlikeInstanceTypes"`
}

type AzureMachinePoolStatus struct {
	// +kubebuilder:validation:Optional
	// Status specific to Azure.
	Provider AzureMachinePoolStatusProvider `json:"provider"`
}

type AzureMachinePoolStatusProvider struct {
	// +kubebuilder:validation:Optional
	// Status of worker nodes.
	Worker AzureMachinePoolStatusProviderWorker `json:"worker"`
}

type AzureMachinePoolStatusProviderWorker struct {
	// +kubebuilder:validation:Optional
	// Azure EC2 instance types used for the worker nodes in this node pool.
	InstanceTypes []string `json:"instanceTypes"`
	// +kubebuilder:validation:Optional
	// Number of EC2 spot instances used in this node pool.
	SpotInstances int `json:"spotInstances"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzureMachinePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AzureMachinePool `json:"items"`
}
