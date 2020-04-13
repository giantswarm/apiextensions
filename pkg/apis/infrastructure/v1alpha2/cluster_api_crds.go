package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
)

const (
	kindCluster                        = "Cluster"
	kindMachineDeployment              = "MachineDeployment"
	clusterDocumentationLink           = "https://pkg.go.dev/sigs.k8s.io/cluster-api/api/v1alpha2?tab=doc#Cluster"
	machineDeploymentDocumentationLink = "https://pkg.go.dev/sigs.k8s.io/cluster-api/api/v1alpha2?tab=doc#MachineDeployment"
)

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
