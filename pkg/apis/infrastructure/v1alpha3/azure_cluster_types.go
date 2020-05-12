package v1alpha3

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	crDocsAnnotation              = "giantswarm.io/docs"
	kindAzureCluster              = "AzureCluster"
	azureClusterDocumentationLink = "https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha3?tab=doc#AzureCluster"
)

func NewAzureClusterCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindAzureCluster)
}

func NewAzureClusterTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAzureCluster,
	}
}

// NewAzureClusterCR returns an AzureCluster Custom Resource.
func NewAzureClusterCR() *AzureCluster {
	return &AzureCluster{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: azureClusterDocumentationLink,
			},
		},
		TypeMeta: NewAzureClusterTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status

// AzureCluster is the infrastructure provider referenced in upstream CAPI Cluster
// CRs.
type AzureCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzureClusterSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status AzureClusterStatus `json:"status,omitempty"`
}

// AzureClusterSpec is the spec part for the AzureCluster resource.
type AzureClusterSpec struct {
	// Cluster provides cluster specification details.
	Cluster AzureClusterSpecCluster `json:"cluster"`
	// Provider holds provider-specific configuration details.
	Provider AzureClusterSpecProvider `json:"provider"`
}

// AzureClusterSpecCluster provides cluster specification details.
type AzureClusterSpecCluster struct {
	// Description is a user-friendly description that should explain the purpose of the
	// cluster to humans.
	Description string `json:"description"`
	// DNS holds DNS configuration details.
	DNS AzureClusterSpecClusterDNS `json:"dns"`
	// +kubebuilder:validation:Optional
	// KubeProxy holds flags passed to kube-proxy on each node.
	KubeProxy AzureClusterSpecClusterKubeProxy `json:"kubeProxy,omitempty"`
	// OIDC holds configuration for OpenID Connect (OIDC) authentication.
	OIDC AzureClusterSpecClusterOIDC `json:"oidc,omitempty"`
}

// AzureClusterSpecClusterDNS holds DNS configuration details.
type AzureClusterSpecClusterDNS struct {
	Domain string `json:"domain"`
}

// AzureClusterSpecClusterOIDC holds configuration for OpenID Connect (OIDC) authentication.
type AzureClusterSpecClusterOIDC struct {
	Claims    AzureClusterSpecClusterOIDCClaims `json:"claims,omitempty"`
	ClientID  string                            `json:"clientID,omitempty"`
	IssuerURL string                            `json:"issuerURL,omitempty"`
}

// AzureClusterSpecClusterOIDCClaims defines OIDC claims.
type AzureClusterSpecClusterOIDCClaims struct {
	Username string `json:"username,omitempty"`
	Groups   string `json:"groups,omitempty"`
}

// AzureClusterSpecClusterKubeProxy describes values passed to the kube-proxy running in a tenant cluster.
type AzureClusterSpecClusterKubeProxy struct {
	// Maximum number of NAT connections to track per CPU core (0 for default).
	// Passed to kube-proxy as --conntrack-max-per-core.
	ConntrackMaxPerCore int `json:"conntrackMaxPerCore,omitempty"`
}

// AzureClusterSpecProvider holds some Azure details.
type AzureClusterSpecProvider struct {
	// CredentialSecret specifies the location of the secret providing the ARN of Azure IAM identity
	// to use with this cluster.
	CredentialSecret AzureClusterSpecProviderCredentialSecret `json:"credentialSecret"`
	// Master holds master node configuration details.
	Master AzureClusterSpecProviderMaster `json:"master"`
	// +kubebuilder:validation:Optional
	// Pod network configuration.
	Pods AzureClusterSpecProviderPods `json:"pods,omitempty"`
	// Region is the Azure region the cluster is to be running in.
	Region string `json:"region"`
}

// AzureClusterSpecProviderCredentialSecret details how to chose the Azure IAM identity ARN
// to use with this cluster.
type AzureClusterSpecProviderCredentialSecret struct {
	// Name is the name of the provider credential resoure.
	Name string `json:"name"`
	// Namespace is the kubernetes namespace that holds the provider credential.
	Namespace string `json:"namespace"`
}

// AzureClusterSpecProviderMaster holds master node configuration details.
type AzureClusterSpecProviderMaster struct {
	// AvailabilityZone is the Azure availability zone to place the master node in.
	AvailabilityZone string `json:"availabilityZone"`
	// InstanceType specifies the Azure EC2 instance type to use for the master node.
	InstanceType string `json:"instanceType"`
}

// AzureClusterSpecProviderPods Pod network configuration.
type AzureClusterSpecProviderPods struct {
	// +kubebuilder:validation:Optional
	// Subnet size, expresses as the count of leading 1 bits in the subnet mask of this subnet.
	CIDRBlock string `json:"cidrBlock,omitempty"`
}

// AzureClusterStatus holds status information about the cluster, populated once the
// cluster is in creation or created.
type AzureClusterStatus struct {
	// +kubebuilder:validation:Optional
	// Cluster provides cluster-specific status details, including conditions and versions.
	Cluster CommonClusterStatus `json:"cluster,omitempty"`
	// +kubebuilder:validation:Optional
	// Provider provides provider-specific status details.
	Provider AzureClusterStatusProvider `json:"provider,omitempty"`
}

// AzureClusterStatusProvider holds provider-specific status details.
type AzureClusterStatusProvider struct {
	// +kubebuilder:validation:Optional
	// Network provides network-specific configuration details
	Network AzureClusterStatusProviderNetwork `json:"network,omitempty"`
}

// AzureClusterStatusProviderNetwork holds network details.
type AzureClusterStatusProviderNetwork struct {
	// +kubebuilder:validation:Optional
	// IPv4 address block used by the tenant cluster, in CIDR notation.
	CIDR string `json:"cidr,omitempty"`
	// +kubebuilder:validation:Optional
	// VPCID contains the ID of the tenant cluster, e.g. vpc-1234567890abcdef0.
	VPCID string `json:"vpcID,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureClusterList is the type returned when listing AzureCLuster resources.
type AzureClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AzureCluster `json:"items"`
}
