package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/annotation"
	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

const (
	kindAWSConfig              = "AWSConfig"
	awsConfigDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/awsconfigs.provider.giantswarm.io/"
)

func NewAWSConfigCRD() *v1beta1.CustomResourceDefinition {
	return crd.LoadV1Beta1(group, kindAWSConfig)
}

// NewAWSClusterTypeMeta returns the populated metav1 metadata object for this CRD.
func NewAWSClusterTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAWSConfig,
	}
}

// NewAWSConfigCR returns a custom resource of type AWSConfig.
func NewAWSConfigCR() *AWSConfig {
	return &AWSConfig{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotation.Docs: awsConfigDocumentationLink,
			},
		},
		TypeMeta: NewAWSClusterTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;giantswarm
// +k8s:openapi-gen=true

// AWSConfig used to represent workload cluster configuration in earlier releases. Deprecated.
type AWSConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AWSConfigSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status AWSConfigStatus `json:"status"`
}

// +k8s:openapi-gen=true
type AWSConfigSpec struct {
	Cluster       Cluster                    `json:"cluster"`
	AWS           AWSConfigSpecAWS           `json:"aws"`
	VersionBundle AWSConfigSpecVersionBundle `json:"versionBundle"`
}

// +k8s:openapi-gen=true
type AWSConfigSpecAWS struct {
	API AWSConfigSpecAWSAPI `json:"api"`
	// TODO remove the deprecated AZ field due to AvailabilityZones.
	//
	//     https://github.com/giantswarm/giantswarm/issues/4507
	//
	AZ string `json:"az"`
	// AvailabilityZones is the number of AWS availability zones used to spread
	// the workload cluster's worker nodes across. There are limitations on
	// availability zone settings due to binary IP range splitting and provider
	// specific region capabilities. When for instance choosing 3 availability
	// zones, the configured IP range will be split into 4 ranges and thus one of
	// it will not be able to be utilized. Such limitations have to be considered
	// when designing the network topology and configuring workload cluster HA via
	// AvailabilityZones.
	//
	// The selection and usage of the actual availability zones for the created
	// workload cluster is randomized. In case there are 4 availability zones
	// provided in the used region and the user selects 2 availability zones, the
	// actually used availability zones in which workload cluster workload is put
	// into will tend to be different across workload cluster creations. This is
	// done in order to provide more HA during single availability zone failures.
	// In case a specific availability zone fails, not all workload clusters will be
	// affected due to the described selection process.
	AvailabilityZones int                  `json:"availabilityZones"`
	CredentialSecret  CredentialSecret     `json:"credentialSecret"`
	Etcd              AWSConfigSpecAWSEtcd `json:"etcd"`

	// HostedZones is AWS hosted zones names in the host cluster account.
	// For each zone there will be "CLUSTER_ID.k8s" NS record created in
	// the host cluster account. Then for each created NS record there will
	// be a zone created in the guest account. After that component
	// specific records under those zones:
	//	- api.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.API.Name }}
	//	- etcd.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Etcd.Name }}
	//	- ingress.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Ingress.Name }}
	//	- *.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Ingress.Name }}
	HostedZones AWSConfigSpecAWSHostedZones `json:"hostedZones"`

	Ingress AWSConfigSpecAWSIngress `json:"ingress"`
	Masters []AWSConfigSpecAWSNode  `json:"masters"`
	Region  string                  `json:"region"`
	VPC     AWSConfigSpecAWSVPC     `json:"vpc"`
	Workers []AWSConfigSpecAWSNode  `json:"workers"`
}

// AWSConfigSpecAWSAPI deprecated since aws-operator v12 resources.
// +k8s:openapi-gen=true
type AWSConfigSpecAWSAPI struct {
	HostedZones string                 `json:"hostedZones"`
	ELB         AWSConfigSpecAWSAPIELB `json:"elb"`
}

// AWSConfigSpecAWSAPIELB deprecated since aws-operator v12 resources.
// +k8s:openapi-gen=true
type AWSConfigSpecAWSAPIELB struct {
	IdleTimeoutSeconds int `json:"idleTimeoutSeconds"`
}

// AWSConfigSpecAWSEtcd deprecated since aws-operator v12 resources.
// +k8s:openapi-gen=true
type AWSConfigSpecAWSEtcd struct {
	HostedZones string                  `json:"hostedZones"`
	ELB         AWSConfigSpecAWSEtcdELB `json:"elb"`
}

// AWSConfigSpecAWSEtcdELB deprecated since aws-operator v12 resources.
// +k8s:openapi-gen=true
type AWSConfigSpecAWSEtcdELB struct {
	IdleTimeoutSeconds int `json:"idleTimeoutSeconds"`
}

// +k8s:openapi-gen=true
type AWSConfigSpecAWSHostedZones struct {
	API     AWSConfigSpecAWSHostedZonesZone `json:"api"`
	Etcd    AWSConfigSpecAWSHostedZonesZone `json:"etcd"`
	Ingress AWSConfigSpecAWSHostedZonesZone `json:"ingress"`
}

// +k8s:openapi-gen=true
type AWSConfigSpecAWSHostedZonesZone struct {
	Name string `json:"name"`
}

// AWSConfigSpecAWSIngress deprecated since aws-operator v12 resources.
// +k8s:openapi-gen=true
type AWSConfigSpecAWSIngress struct {
	HostedZones string                     `json:"hostedZones"`
	ELB         AWSConfigSpecAWSIngressELB `json:"elb"`
}

// AWSConfigSpecAWSIngressELB deprecated since aws-operator v12 resources.
// +k8s:openapi-gen=true
type AWSConfigSpecAWSIngressELB struct {
	IdleTimeoutSeconds int `json:"idleTimeoutSeconds"`
}

// +k8s:openapi-gen=true
type AWSConfigSpecAWSNode struct {
	ImageID            string `json:"imageID"`
	InstanceType       string `json:"instanceType"`
	DockerVolumeSizeGB int    `json:"dockerVolumeSizeGB"`
}

// +k8s:openapi-gen=true
type AWSConfigSpecAWSVPC struct {
	CIDR              string   `json:"cidr"`
	PrivateSubnetCIDR string   `json:"privateSubnetCidr"`
	PublicSubnetCIDR  string   `json:"publicSubnetCidr"`
	RouteTableNames   []string `json:"routeTableNames"`
	PeerID            string   `json:"peerId"`
}

// +k8s:openapi-gen=true
type AWSConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type AWSConfigStatus struct {
	AWS     AWSConfigStatusAWS `json:"aws"`
	Cluster StatusCluster      `json:"cluster"`
}

// +k8s:openapi-gen=true
type AWSConfigStatusAWS struct {
	AvailabilityZones []AWSConfigStatusAWSAvailabilityZone `json:"availabilityZones"`
	AutoScalingGroup  AWSConfigStatusAWSAutoScalingGroup   `json:"autoScalingGroup"`
}

// +k8s:openapi-gen=true
type AWSConfigStatusAWSAutoScalingGroup struct {
	Name string `json:"name"`
}

// +k8s:openapi-gen=true
type AWSConfigStatusAWSAvailabilityZone struct {
	Name   string                                   `json:"name"`
	Subnet AWSConfigStatusAWSAvailabilityZoneSubnet `json:"subnet"`
}

// +k8s:openapi-gen=true
type AWSConfigStatusAWSAvailabilityZoneSubnet struct {
	Private AWSConfigStatusAWSAvailabilityZoneSubnetPrivate `json:"private"`
	Public  AWSConfigStatusAWSAvailabilityZoneSubnetPublic  `json:"public"`
}

// +k8s:openapi-gen=true
type AWSConfigStatusAWSAvailabilityZoneSubnetPrivate struct {
	CIDR string `json:"cidr"`
}

// +k8s:openapi-gen=true
type AWSConfigStatusAWSAvailabilityZoneSubnetPublic struct {
	CIDR string `json:"cidr"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AWSConfig `json:"items"`
}
