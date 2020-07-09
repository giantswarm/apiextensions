package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindAzureConfig = "AzureConfig"
)

func NewAzureConfigCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindAzureConfig)
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=azure;giantswarm
// +k8s:openapi-gen=true

type AzureConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AzureConfigSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status AzureConfigStatus `json:"status"`
}

// +k8s:openapi-gen=true
type AzureConfigSpec struct {
	Cluster       Cluster                      `json:"cluster"`
	Azure         AzureConfigSpecAzure         `json:"azure"`
	VersionBundle AzureConfigSpecVersionBundle `json:"versionBundle"`
}

// +k8s:openapi-gen=true
type AzureConfigSpecAzure struct {
	// +kubebuilder:validation:Optional
	// +nullable
	AvailabilityZones []int                              `json:"availabilityZones"`
	CredentialSecret  CredentialSecret                   `json:"credentialSecret"`
	DNSZones          AzureConfigSpecAzureDNSZones       `json:"dnsZones"`
	Masters           []AzureConfigSpecAzureNode         `json:"masters"`
	VirtualNetwork    AzureConfigSpecAzureVirtualNetwork `json:"virtualNetwork"`
	// +kubebuilder:validation:Optional
	// +nullable
	Workers []AzureConfigSpecAzureNode `json:"workers"`
}

// AzureConfigSpecAzureDNSZones contains the DNS Zones of the cluster.
// +k8s:openapi-gen=true
type AzureConfigSpecAzureDNSZones struct {
	// API is the DNS Zone for the Kubernetes API.
	API AzureConfigSpecAzureDNSZonesDNSZone `json:"api"`
	// Etcd is the DNS Zone for the etcd cluster.
	Etcd AzureConfigSpecAzureDNSZonesDNSZone `json:"etcd"`
	// Ingress is the DNS Zone for the Ingress resource, used for customer traffic.
	Ingress AzureConfigSpecAzureDNSZonesDNSZone `json:"ingress"`
}

// AzureConfigSpecAzureDNSZonesDNSZone points to a DNS Zone in Azure.
// +k8s:openapi-gen=true
type AzureConfigSpecAzureDNSZonesDNSZone struct {
	// ResourceGroup is the resource group of the zone.
	ResourceGroup string `json:"resourceGroup"`
	// Name is the name of the zone.
	Name string `json:"name"`
}

// +k8s:openapi-gen=true
type AzureConfigSpecAzureVirtualNetwork struct {
	// CIDR is the CIDR for the Virtual Network.
	CIDR string `json:"cidr"`

	// TODO: remove Master, Worker and Calico subnet cidr after azure-operator v2
	// is deleted. MasterSubnetCIDR is the CIDR for the master subnet.
	//
	//     https://github.com/giantswarm/giantswarm/issues/4358
	//
	MasterSubnetCIDR string `json:"masterSubnetCIDR"`
	// WorkerSubnetCIDR is the CIDR for the worker subnet.
	WorkerSubnetCIDR string `json:"workerSubnetCIDR"`

	// CalicoSubnetCIDR is the CIDR for the calico subnet. It has to be
	// also a worker subnet (Azure limitation).
	CalicoSubnetCIDR string `json:"calicoSubnetCIDR"`
}

// +k8s:openapi-gen=true
type AzureConfigSpecAzureNode struct {
	// VMSize is the master vm size (e.g. Standard_A1)
	VMSize string `json:"vmSize"`
	// DockerVolumeSizeGB is the size of a volume mounted to /var/lib/docker.
	DockerVolumeSizeGB int `json:"dockerVolumeSizeGB"`
	// KubeletVolumeSizeGB is the size of a volume mounted to /var/lib/kubelet.
	KubeletVolumeSizeGB int `json:"kubeletVolumeSizeGB"`
}

// +k8s:openapi-gen=true
type AzureConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type AzureConfigStatus struct {
	// +kubebuilder:validation:Optional
	Cluster StatusCluster `json:"cluster"`
	// +kubebuilder:validation:Optional
	Provider AzureConfigStatusProvider `json:"provider"`
}

// +k8s:openapi-gen=true
type AzureConfigStatusProvider struct {
	// +kubebuilder:validation:Optional
	// +nullable
	AvailabilityZones []int `json:"availabilityZones,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	Ingress AzureConfigStatusProviderIngress `json:"ingress"`
}

// +k8s:openapi-gen=true
type AzureConfigStatusProviderIngress struct {
	// +kubebuilder:validation:Optional
	// +nullable
	LoadBalancer AzureConfigStatusProviderIngressLoadBalancer `json:"loadBalancer"`
}

// +k8s:openapi-gen=true
type AzureConfigStatusProviderIngressLoadBalancer struct {
	// +kubebuilder:validation:Optional
	PublicIPName string `json:"publicIPName"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzureConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AzureConfig `json:"items"`
}
