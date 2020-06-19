package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/apis/core"
	"github.com/giantswarm/apiextensions/pkg/crd"
	"github.com/giantswarm/apiextensions/pkg/key"
)

const (
	DrainerConfigStatusStatusTrue  = "True"
	DrainerConfigStatusTypeDrained = "Drained"
	DrainerConfigStatusTypeTimeout = "Timeout"
)

// NewAppCatalogCRD returns a CRD defining a DrainerConfig.
func NewDrainerConfigCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(core.Group, core.KindDrainerConfig)
}

// NewDrainerConfigCR returns a DrainerConfig custom resource.
func NewDrainerConfigCR(name, namespace string) *DrainerConfig {
	cr := DrainerConfig{}
	cr.TypeMeta, cr.ObjectMeta = key.NewMeta(SchemeGroupVersion, core.KindDrainerConfig, name, namespace)
	return &cr
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=common;giantswarm
// +k8s:openapi-gen=true

type DrainerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              DrainerConfigSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status DrainerConfigStatus `json:"status"`
}

// +k8s:openapi-gen=true
type DrainerConfigSpec struct {
	Guest         DrainerConfigSpecGuest         `json:"guest"`
	VersionBundle DrainerConfigSpecVersionBundle `json:"versionBundle"`
}

// +k8s:openapi-gen=true
type DrainerConfigSpecGuest struct {
	Cluster DrainerConfigSpecGuestCluster `json:"cluster"`
	Node    DrainerConfigSpecGuestNode    `json:"node"`
}

// +k8s:openapi-gen=true
type DrainerConfigSpecGuestCluster struct {
	API DrainerConfigSpecGuestClusterAPI `json:"api"`
	// ID is the guest cluster ID of which a node should be drained.
	ID string `json:"id"`
}

// +k8s:openapi-gen=true
type DrainerConfigSpecGuestClusterAPI struct {
	// Endpoint is the guest cluster API endpoint.
	Endpoint string `json:"endpoint"`
}

// +k8s:openapi-gen=true
type DrainerConfigSpecGuestNode struct {
	// Name is the identifier of the guest cluster's master and worker nodes. In
	// Kubernetes/Kubectl they are represented as node names. The names are manage
	// in an abstracted way because of provider specific differences.
	//
	//     AWS: EC2 instance DNS.
	//     Azure: VM name.
	//     KVM: host cluster pod name.
	//
	Name string `json:"name"`
}

// +k8s:openapi-gen=true
type DrainerConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type DrainerConfigStatus struct {
	Conditions []DrainerConfigStatusCondition `json:"conditions"`
}

// DrainerConfigStatusCondition expresses a condition in which a node may is.
// +k8s:openapi-gen=true
type DrainerConfigStatusCondition struct {
	// LastHeartbeatTime is the last time we got an update on a given condition.
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime"`
	// LastTransitionTime is the last time the condition transitioned from one
	// status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// Status may be True, False or Unknown.
	Status string `json:"status"`
	// Type may be Pending, Ready, Draining, Drained.
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DrainerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []DrainerConfig `json:"items"`
}
