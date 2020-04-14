package v1alpha2

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	kindAWSMachineDeployment = "AWSMachineDeployment"
)

func NewAWSMachineDeploymentCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadCRD(group, kindAWSMachineDeployment)
}

func NewAWSMachineDeploymentTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAWSMachineDeployment,
	}
}

// NewAWSMachineDeploymentCR returns an AWSMachineDeployment Custom Resource.
func NewAWSMachineDeploymentCR() *AWSMachineDeployment {
	return &AWSMachineDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: awsClusterDocumentationLink,
			},
		},
		TypeMeta: NewAWSMachineDeploymentTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSMachineDeployment is the infrastructure provider referenced in upstream
// CAPI MachineDeployment CRs.
type AWSMachineDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AWSMachineDeploymentSpec   `json:"spec" yaml:"spec"`
	Status            AWSMachineDeploymentStatus `json:"status" yaml:"status"`
}

type AWSMachineDeploymentSpec struct {
	NodePool AWSMachineDeploymentSpecNodePool `json:"nodePool" yaml:"nodePool"`
	Provider AWSMachineDeploymentSpecProvider `json:"provider" yaml:"provider"`
}

type AWSMachineDeploymentSpecNodePool struct {
	Description string                                  `json:"description" yaml:"description"`
	Machine     AWSMachineDeploymentSpecNodePoolMachine `json:"machine" yaml:"machine"`
	Scaling     AWSMachineDeploymentSpecNodePoolScaling `json:"scaling" yaml:"scaling"`
}

type AWSMachineDeploymentSpecNodePoolMachine struct {
	DockerVolumeSizeGB  int `json:"dockerVolumeSizeGB" yaml:"dockerVolumeSizeGB"`
	KubeletVolumeSizeGB int `json:"kubeletVolumeSizeGB" yaml:"kubeletVolumeSizeGB"`
}

type AWSMachineDeploymentSpecNodePoolScaling struct {
	Max int `json:"max" yaml:"max"`
	Min int `json:"min" yaml:"min"`
}

type AWSMachineDeploymentSpecProvider struct {
	AvailabilityZones    []string                                     `json:"availabilityZones" yaml:"availabilityZones"`
	InstanceDistribution AWSMachineDeploymentSpecInstanceDistribution `json:"instanceDistribution" yaml:"instanceDistribution"`
	Worker               AWSMachineDeploymentSpecProviderWorker       `json:"worker" yaml:"worker"`
}

type AWSMachineDeploymentSpecInstanceDistribution struct {
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	OnDemandBaseCapacity int `json:"onDemandBaseCapacity" yaml:"onDemandBaseCapacity"`
	// +kubebuilder:default=100
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Minimum=0
	OnDemandPercentageAboveBaseCapacity int `json:"onDemandPercentageAboveBaseCapacity" yaml:"onDemandPercentageAboveBaseCapacity"`
}

type AWSMachineDeploymentSpecProviderWorker struct {
	InstanceType          string `json:"instanceType" yaml:"instanceType"`
	UseAlikeInstanceTypes bool   `json:"useAlikeInstanceTypes" yaml:"useAlikeInstanceTypes"`
}

type AWSMachineDeploymentStatus struct {
	Provider AWSMachineDeploymentStatusProvider `json:"provider" yaml:"provider"`
}

type AWSMachineDeploymentStatusProvider struct {
	Worker AWSMachineDeploymentStatusProviderWorker `json:"worker" yaml:"worker"`
}

type AWSMachineDeploymentStatusProviderWorker struct {
	InstanceTypes []string `json:"instanceTypes" yaml:"instanceTypes"`
	SpotInstances int      `json:"spotInstances" yaml:"spotInstances"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSMachineDeploymentList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []AWSMachineDeployment `json:"items" yaml:"items"`
}
