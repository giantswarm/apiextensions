package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=aws;giantswarm
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
// AWSClusterConfig used to represent workload cluster configuration in earlier releases. Deprecated.
type AWSClusterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AWSClusterConfigSpec `json:"spec"`
}

// +k8s:openapi-gen=true
type AWSClusterConfigSpec struct {
	Guest         AWSClusterConfigSpecGuest         `json:"guest"`
	VersionBundle AWSClusterConfigSpecVersionBundle `json:"versionBundle"`
}

// +k8s:openapi-gen=true
type AWSClusterConfigSpecGuest struct {
	ClusterGuestConfig `json:",inline"`
	CredentialSecret   AWSClusterConfigSpecGuestCredentialSecret `json:"credentialSecret"`
	Masters            []AWSClusterConfigSpecGuestMaster         `json:"masters,omitempty"`
	Workers            []AWSClusterConfigSpecGuestWorker         `json:"workers,omitempty"`
}

// AWSClusterConfigSpecGuestCredentialSecret points to the K8s Secret
// containing credentials for an AWS account in which the workload cluster should
// be created.
// +k8s:openapi-gen=true
type AWSClusterConfigSpecGuestCredentialSecret struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type AWSClusterConfigSpecGuestMaster struct {
	AWSClusterConfigSpecGuestNode `json:",inline"`
}

type AWSClusterConfigSpecGuestWorker struct {
	AWSClusterConfigSpecGuestNode `json:",inline"`
	Labels                        map[string]string `json:"labels"`
}
type AWSClusterConfigSpecGuestNode struct {
	ID           string `json:"id"`
	InstanceType string `json:"instanceType,omitempty"`
}

// +k8s:openapi-gen=true
type AWSClusterConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSClusterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AWSClusterConfig `json:"items"`
}
