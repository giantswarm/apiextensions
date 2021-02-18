package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/cluster-api/api/v1alpha3"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:object:root=true

// KVMCluster is the infrastructure provider referenced in upstream CAPI Cluster
// CRs.
type KVMCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KVMClusterSpec `json:"spec"`
}

// +k8s:openapi-gen=true

// KVMClusterSpec is the spec part for the KVMCluster resource.
type KVMClusterSpec struct {
	// Endpoint used to connect to the target cluster's Kubernetes API server.
	ControlPlaneEndpoint v1alpha3.APIEndpoint `json:"controlPlaneEndpoint"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// KVMClusterSpecCluster provides cluster specification details.
type KVMClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KVMCluster `json:"items"`
}
