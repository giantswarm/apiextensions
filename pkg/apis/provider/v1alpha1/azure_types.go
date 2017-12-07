package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewAzureConfigCRD returns a new custom resource definition for AzureConfig.
// This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: azureconfigs.provider.giantswarm.io
//     spec:
//       group: provider.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: AzureConfig
//         plural: azureconfigs
//         singular: azureconfig
//
func NewAzureConfigCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "azureconfigs.provider.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "provider.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "AzureConfig",
				Plural:   "azureconfigs",
				Singular: "azureconfig",
			},
		},
	}
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzureConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AzureConfigSpec `json:"spec"`
}

type AzureConfigSpec struct {
	Cluster Cluster              `json:"cluster" yaml:"cluster"`
	Azure   AzureConfigSpecAzure `json:"azure" yaml:"azure"`
}

type AzureConfigSpecAzure struct {
	KeyVault       AzureConfigSpecAzureKeyVault       `json:"keyVault" yaml:"keyVault"`
	ResourceGroup  AzureConfigSpecAzureResourceGroup  `json:"resourceGroup" yaml:"resourceGroup"`
	Storage        AzureConfigSpecAzureStorage        `json:"storage" yaml:"storage"`
	VirtualNetwork AzureConfigSpecAzureVirtualNetwork `json:"virtualNetwork" yaml:"virtualNetwork"`
	Masters        []AzureConfigSpecAzureNode         `json:"masters" yaml:"masters"`
	Workers        []AzureConfigSpecAzureNode         `json:"workers" yaml:"workers"`
	DNSZones       AzureConfigSpecAzureDNSZones       `json:"dnsZones" yaml:"dnsZones"`
}

type AzureConfigSpecAzureResourceGroup struct {
	Location string `json:"location" yaml:"location"`
}

// DNSZones contains the DNS Zones of the cluster.
type AzureConfigSpecAzureDNSZones struct {
	// API is the DNS Zone for the Kubernetes API.
	API string `json:"api" yaml:"api"`
	// Etcd is the DNS Zone for the etcd cluster.
	Etcd string `json:"etcd" yaml:"etcd"`
	// Ingress is the DNS Zone for the Ingress resource, used for customer traffic.
	Ingress string `json:"ingress" yaml:"ingress"`
}

type AzureConfigSpecAzureKeyVault struct {
	// Name is the name of the Azure Key Vault. It must be globally unique,
	// 3-24 characters in length and contain only (0-9, a-z, A-Z, and -).
	Name string `json:"name" yaml:"name"`
}

type AzureConfigSpecAzureNode struct {
	// VMSize is the master vm size (e.g. Standard_A1)
	VMSize string `json:"vmSize" yaml:"vmSize"`
	// DataDiskSizeGB is the vm data disk size in GB
	DataDiskSizeGB int `json:"dataDiskSizeGB" yaml:"dataDiskSizeGB"`
	// AdminUsername is the vm administrator username
	AdminUsername string `json:"adminUsername" yaml:"adminUsername"`
	//  AdminSSHKeyData is the vm administrator ssh public key
	AdminSSHKeyData string `json:"adminSSHKeyData" yaml:"adminSSHKeyData"`
	// OSImage is the vm OS image object
	OSImage AzureConfigSpecAzureNodeOSImage `json:"osImage" yaml:"osImage"`
}

type AzureConfigSpecAzureNodeOSImage struct {
	// Publisher is the image publisher (e.g GiantSwarm)
	Publisher string `json:"publisher" yaml:"publisher"`
	// Offer is the image offered by the publisher (e.g. CoreOS)
	Offer string `json:"offer" yaml:"offer"`
	// SKU is the image SKU (e.g. Alpha)
	SKU string `json:"sku" yaml:"sku"`
	// Version is the image version (e.g. 1465.7.0)
	Version string `json:"version" yaml:"version"`
}

type AzureConfigSpecAzureStorage struct {
	// AccountType is the Azure Storage Account Type.
	AccountType string `json:"accountType" yaml:"accountType"`
}

type AzureConfigSpecAzureVirtualNetwork struct {
	// CIDR is the CIDR for the Virtual Network.
	CIDR string `json:"cidr" yaml:"cidr"`
	// MasterSubnetCIDR is the CIDR for the master subnet,
	MasterSubnetCIDR string `json:"masterSubnetCIDR" yaml:"masterSubnetCIDR"`
	// WorkerSubnetCIDR is the CIDR for the worker subnet,
	WorkerSubnetCIDR string                                         `json:"workerSubnetCIDR" yaml:"workerSubnetCIDR"`
	LoadBalancer     AzureConfigSpecAzureVirtualNetworkLoadBalancer `json:"loadBalancer" yaml:"loadBalancer"`
}

type AzureConfigSpecAzureVirtualNetworkLoadBalancer struct {
	// EtcdCidr is the CIDR for the etcd load balancer.
	EtcdCIDR string `json:"etcdCIDR" yaml:"etcdCIDR"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzureConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AzureConfig `json:"items"`
}
