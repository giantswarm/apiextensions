package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v2/pkg/crd"
)

const (
	kindSpark = "Spark"
)

func NewSparkCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindSpark)
}

func NewSparkTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindSpark,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=common;giantswarm
// +k8s:openapi-gen=true

// Spark is a Kubernetes resource (CR) which is based on the Spark CRD defined above.
//
// An example Spark resource can be viewed here
// https://github.com/giantswarm/apiextensions/blob/master/docs/cr/core.giantswarm.io_v1alpha1_spark.yaml
type Spark struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              SparkSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status SparkStatus `json:"status"`
}

// SparkSpec is the interface which defines the input parameters for
// a newly rendered g8s ignition template.
// +k8s:openapi-gen=true
type SparkSpec struct {
	// +nullable
	Values map[string]string `json:"values,omitempty"`
}

// SparkStatus holds the rendering result.
// +k8s:openapi-gen=true
type SparkStatus struct {
	// DataSecretName is a name of the secret containing the rendered ignition once created.
	DataSecretName string `json:"dataSecretName"`
	// FailureReason is a short string indicating the reason rendering failed (if it did).
	FailureReason string `json:"failureReason"`
	// FailureMessage is a longer message indicating the reason rendering failed (if it did).
	FailureMessage string `json:"failureMessage"`
	// Ready will be true when the referenced secret contains the rendered ignition and can be used for creating nodes.
	Ready bool `json:"ready"`
	// Verification is a hash of the rendered ignition to ensure that it has
	// not been changed when loaded as a remote file by the bootstrap ignition.
	// See https://coreos.com/ignition/docs/latest/configuration-v2_2.html
	Verification SparkStatusVerification `json:"verification"`
}

// +k8s:openapi-gen=true
type SparkStatusVerification struct {
	// The content of the full rendered ignition hashed by the corresponding algorithm.
	Hash string `json:"hash"`
	// The algorithm used for hashing. Must be sha512 for now.
	Algorithm string `json:"algorithm"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SparkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Spark `json:"items"`
}
