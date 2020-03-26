package v1alpha2

import (
	"fmt"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterapiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/yaml"
)

const (
	kindCluster = "Cluster"
	// TODO: Change to this CRD's docs URL once published.
	clusterDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/"
)

const clusterCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    controller-tools.k8s.io: "1.0"
  name: clusters.cluster.x-k8s.io
spec:
  conversion:
    strategy: None
  group: cluster.x-k8s.io
  names:
    kind: Cluster
    listKind: ClusterList
    plural: clusters
    singular: cluster
  preserveUnknownFields: true
  scope: Namespaced
  versions:
  - name: v1alpha2
    served: true
    storage: true
    subresources:
      status: {}
`

var clusterCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(clusterCRDYAML), &clusterCRD)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// NewClusterCRD returns a new custom resource definition for Cluster (from
// Cluster API). This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: clusters.cluster.x-k8s.io
//     spec:
//       group: cluster.x-k8s.io
//       scope: Namespaced
//       version: v1alpha2
//       names:
//         kind: Cluster
//         plural: clusters
//         singular: cluster
//       subresources:
//         status: {}
//
func NewClusterCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return clusterCRD.DeepCopy()
}

func NewClusterTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindCluster,
	}
}

// NewClusterCR returns a Cluster Custom Resource.
func NewClusterCR() *clusterapiv1alpha2.Cluster {
	return &clusterapiv1alpha2.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: clusterDocumentationLink,
			},
		},
		TypeMeta: NewClusterTypeMeta(),
	}
}

// NewMachineDeploymentCRD returns a new custom resource definition for
// MachineDeployment (from Cluster API). This might look something like the
// following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: machinedeployments.cluster.x-k8s.io
//     spec:
//       group: cluster.x-k8s.io
//       scope: Namespaced
//       version: v1alpha2
//       names:
//         kind: MachineDeployment
//         plural: machinedeployments
//         singular: machinedeployment
//       subresources:
//         status: {}
//
func NewMachineDeploymentCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"controller-tools.k8s.io": "1.0",
			},
			Name: "machinedeployments.cluster.x-k8s.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group: "cluster.x-k8s.io",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:   "MachineDeployment",
				Plural: "machinedeployments",
			},
			Scope: apiextensionsv1beta1.NamespaceScoped,
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
			Version: "v1alpha2",
		},
	}
}
