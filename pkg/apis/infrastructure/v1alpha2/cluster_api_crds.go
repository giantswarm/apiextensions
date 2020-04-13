package v1alpha2

import (
	"fmt"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/yaml"
)

const (
	kindCluster                        = "Cluster"
	kindMachineDeployment              = "MachineDeployment"
	clusterDocumentationLink           = "https://pkg.go.dev/sigs.k8s.io/cluster-api/api/v1alpha2?tab=doc#Cluster"
	machineDeploymentDocumentationLink = "https://pkg.go.dev/sigs.k8s.io/cluster-api/api/v1alpha2?tab=doc#MachineDeployment"
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

const machineDeploymentCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    controller-tools.k8s.io: "1.0"
  name: machinedeployments.cluster.x-k8s.io
spec:
  conversion:
    strategy: None
  group: cluster.x-k8s.io
  names:
    kind: MachineDeployment
    listKind: MachineDeploymentList
    plural: machinedeployments
    singular: machinedeployment
  preserveUnknownFields: true
  scope: Namespaced
  versions:
  - name: v1alpha2
    served: true
    storage: true
    subresources:
      status: {}
`

var (
	clusterCRD           *apiextensionsv1beta1.CustomResourceDefinition
	machineDeploymentCRD *apiextensionsv1beta1.CustomResourceDefinition
)

func init() {
	err := yaml.Unmarshal([]byte(clusterCRDYAML), &clusterCRD)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = yaml.Unmarshal([]byte(machineDeploymentCRDYAML), &machineDeploymentCRD)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// NewClusterCRD returns a new custom resource definition for Cluster (from
// Cluster API).
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
func NewClusterCR() *apiv1alpha2.Cluster {
	return &apiv1alpha2.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: clusterDocumentationLink,
			},
		},
		TypeMeta: NewClusterTypeMeta(),
	}
}

// NewMachineDeploymentCRD returns a new custom resource definition for
// MachineDeployment (from Cluster API).
func NewMachineDeploymentCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return machineDeploymentCRD.DeepCopy()
}

// NewMachineDeploymentTypeMeta returns the type block for a MachineDeployment CR.
func NewMachineDeploymentTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindMachineDeployment,
	}
}

// NewMachineDeploymentCR returns a MachineDeployment Custom Resource.
func NewMachineDeploymentCR() *apiv1alpha2.MachineDeployment {
	return &apiv1alpha2.MachineDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: machineDeploymentDocumentationLink,
			},
		},
		TypeMeta: NewMachineDeploymentTypeMeta(),
	}
}
