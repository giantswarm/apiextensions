package v1alpha2

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/apis/infrastructure"
	"github.com/giantswarm/apiextensions/pkg/crd"
	"github.com/giantswarm/apiextensions/pkg/key"
)

// NewAWSMachineDeploymentCRD returns a CRD defining an AWSMachineDeployment.
func NewAWSMachineDeploymentCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(infrastructure.Group, infrastructure.KindAWSMachineDeployment)
}

// NewAWSMachineDeploymentCR returns an AWSMachineDeployment custom resource.
func NewAWSMachineDeploymentCR(name, namespace string) *AWSMachineDeployment {
	cr := AWSMachineDeployment{}
	cr.TypeMeta, cr.ObjectMeta = key.NewMeta(SchemeGroupVersion, infrastructure.KindAWSMachineDeployment, name, namespace)
	return &cr
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;giantswarm;cluster-api

// AWSMachineDeployment is the infrastructure provider referenced in Kubernetes Cluster API MachineDeployment resources.
// It contains provider-specific specification and status for a node pool.
// In use on AWS since Giant Swarm release v10.x.x and reconciled by aws-operator.
type AWSMachineDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Contains the specification.
	Spec AWSMachineDeploymentSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Holds status information.
	Status AWSMachineDeploymentStatus `json:"status,omitempty"`
}

type AWSMachineDeploymentSpec struct {
	// Specifies details of node pool and the worker nodes it should contain.
	NodePool AWSMachineDeploymentSpecNodePool `json:"nodePool"`
	// Contains AWS specific details.
	Provider AWSMachineDeploymentSpecProvider `json:"provider"`
}

type AWSMachineDeploymentSpecNodePool struct {
	// User-friendly name or description of the purpose of the node pool.
	Description string `json:"description"`
	// Specification of the worker node machine.
	Machine AWSMachineDeploymentSpecNodePoolMachine `json:"machine"`
	// Scaling settings for the node pool, configuring the cluster-autoscaler
	// determining the number of nodes to have in this node pool.
	Scaling AWSMachineDeploymentSpecNodePoolScaling `json:"scaling"`
}

type AWSMachineDeploymentSpecNodePoolMachine struct {
	// Size of the volume reserved for Docker images and overlay file systems of
	// Docker containers. Unit: 1 GB = 1,000,000,000 Bytes.
	DockerVolumeSizeGB int `json:"dockerVolumeSizeGB"`
	// Size of the volume reserved for the kubelet, which can be used by Pods via
	// volumes of type EmptyDir. Unit: 1 GB = 1,000,000,000 Bytes.
	KubeletVolumeSizeGB int `json:"kubeletVolumeSizeGB"`
}

type AWSMachineDeploymentSpecNodePoolScaling struct {
	// Maximum number of worker nodes in this node pool.
	Max int `json:"max"`
	// Minimum number of worker nodes in this node pool.
	Min int `json:"min"`
}

type AWSMachineDeploymentSpecProvider struct {
	// Name(s) of the availability zone(s) to use for worker nodes. Using multiple
	// availability zones results in higher resilience but can also result in higher
	// cost due to network traffic between availability zones.
	AvailabilityZones []string `json:"availabilityZones"`
	// +kubebuilder:validation:Optional
	// Settings defining the distribution of on-demand and spot instances in the node pool.
	InstanceDistribution AWSMachineDeploymentSpecInstanceDistribution `json:"instanceDistribution,omitempty"`
	// Specification of worker nodes.
	Worker AWSMachineDeploymentSpecProviderWorker `json:"worker"`
}

type AWSMachineDeploymentSpecInstanceDistribution struct {
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	// Base capacity of on-demand instances to use for worker nodes in this pool. When this larger
	// than 0, this value defines a number of worker nodes that will be created using on-demand
	// EC2 instances, regardless of the value configured as `onDemandPercentageAboveBaseCapacity`.
	OnDemandBaseCapacity int `json:"onDemandBaseCapacity,omitempty"`
	// +kubebuilder:default=100
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Minimum=0
	// Percentage of on-demand EC2 instances to use for worker nodes, instead of spot instances,
	// for instances exceeding `onDemandBaseCapacity`. For example, to have half of the worker nodes
	// use spot instances and half use on-demand, set this value to 50.
	OnDemandPercentageAboveBaseCapacity int `json:"onDemandPercentageAboveBaseCapacity,omitempty"`
}

type AWSMachineDeploymentSpecProviderWorker struct {
	// AWS EC2 instance type name to use for the worker nodes in this node pool.
	InstanceType string `json:"instanceType"`
	// +kubebuilder:default=false
	// If true, certain instance types with specs similar to instanceType will be used.
	UseAlikeInstanceTypes bool `json:"useAlikeInstanceTypes"`
}

type AWSMachineDeploymentStatus struct {
	// +kubebuilder:validation:Optional
	// Status specific to AWS.
	Provider AWSMachineDeploymentStatusProvider `json:"provider,omitempty"`
}

type AWSMachineDeploymentStatusProvider struct {
	// +kubebuilder:validation:Optional
	// Status of worker nodes.
	Worker AWSMachineDeploymentStatusProviderWorker `json:"worker,omitempty"`
}

type AWSMachineDeploymentStatusProviderWorker struct {
	// +kubebuilder:validation:Optional
	// AWS EC2 instance types used for the worker nodes in this node pool.
	InstanceTypes []string `json:"instanceTypes,omitempty"`
	// +kubebuilder:validation:Optional
	// Number of EC2 spot instances used in this node pool.
	SpotInstances int `json:"spotInstances,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSMachineDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AWSMachineDeployment `json:"items"`
}
