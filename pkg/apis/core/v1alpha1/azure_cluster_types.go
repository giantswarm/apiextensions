package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindAzureClusterConfig = "AzureClusterConfig"
)

func NewAzureClusterConfigCRD() *v1.CustomResourceDefinition {
	return crd.Load(group, kindAzureClusterConfig)
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=azure;giantswarm
// +kubebuilder:storageversion

type AzureClusterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AzureClusterConfigSpec `json:"spec"`
}

type AzureClusterConfigSpec struct {
	Guest         AzureClusterConfigSpecGuest         `json:"guest"`
	VersionBundle AzureClusterConfigSpecVersionBundle `json:"versionBundle"`
}

type AzureClusterConfigSpecGuest struct {
	ClusterGuestConfig `json:",inline"`
	CredentialSecret   AzureClusterConfigSpecGuestCredentialSecret `json:"credentialSecret"`
	Masters            []AzureClusterConfigSpecGuestMaster         `json:"masters,omitempty"`
	Workers            []AzureClusterConfigSpecGuestWorker         `json:"workers,omitempty"`
}

// AzureClusterConfigSpecGuestCredentialSecret points to the K8s Secret
// containing credentials for an Azure subscription in which the tenant cluster
// should be created.
type AzureClusterConfigSpecGuestCredentialSecret struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type AzureClusterConfigSpecGuestMaster struct {
	AzureClusterConfigSpecGuestNode `json:",inline"`
}

type AzureClusterConfigSpecGuestWorker struct {
	AzureClusterConfigSpecGuestNode `json:",inline"`
	// +kubebuilder:validation:Optional
	// +nullable
	Labels map[string]string `json:"labels"`
}

type AzureClusterConfigSpecGuestNode struct {
	ID     string `json:"id"`
	VMSize string `json:"vmSize,omitempty"`
}

type AzureClusterConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzureClusterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AzureClusterConfig `json:"items"`
}
