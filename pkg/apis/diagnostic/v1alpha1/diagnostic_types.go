package v1alpha1

import (
	"fmt"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindTcpdump = "Tcpdump"
)

// NewTcpdumpCRD returns a new custom resource definition for Tcpdump. This
// might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: tcpdump.diagnostic.giantswarm.io
//     spec:
//       group: diagnostic.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: Tcpdump
//         plural: tcpdumps
//         singular: tcpdump
//         shortNames:
//         - tcpdump
//
func NewTcpdumpCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("tcpdumps.%s", group),
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   group,
			Scope:   "Namespaced",
			Version: version,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     kindTcpdump,
				Plural:   "tcpdumps",
				Singular: "tcpdump",
			},
		},
	}
}

func NewTcpdumpTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindTcpdump,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Tcpdump struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              TcpdumpSpec `json:"spec"`
	// +optional
	Status TcpdumpStatus `json:"status"`
}

// TcpdumpSpec is the interface which defines the input parameters for
// a newly rendered g8s tcpdump diagnostic template.
type TcpdumpSpec struct {
	// +optional
	Application TcpdumpSpecApplication `json:"application,omitempty"`
	Exporter    TcpdumpSpecExporter    `json:"exporter,omitempty"`
}

type TcpdumpSpecApplication struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Version   string `json:"version,omitempty"`
	Catalog   string `json:"catalog,omitempty"`
}

// This only supports public S3 bucket for now
type TcpdumpSpecExporter struct {
	S3Bucket TcpdumpSpecExporterS3Bucket `json:"s3bucket"`
}

type TcpdumpSpecExporterS3Bucket struct {
	Name string `json:"name"`
}

// TcpdumpStatus holds the rendering result.
type TcpdumpStatus struct {
	State       string                   `json:"state" yaml: "state"`
	Application TcpdumpStatusApplication `json:"application"`
	Exports     TcpdumpStatusExports     `json:"exports"`
}

type TcpdumpStatusApplication struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Version   string `json:"version"`
	Catalog   string `json:"catalog"`
}

type TcpdumpStatusExports struct {
	Nodes []TcpdumpStatusExportsNode `json:"nodes,omitempty"`
}

// TODO this needs better semantics
type TcpdumpStatusExportsNode struct {
	Name      string `json:"name"`
	BucketURL string `json:"bucketURL,omitempty"`
	Key       string `json:"file,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TcpdumpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Tcpdump `json:"items"`
}
