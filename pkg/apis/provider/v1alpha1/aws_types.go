package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	crDocsAnnotation = "giantswarm.io/docs"
	kindAWSConfig    = "AWSConfig"
	// TODO: Change to https://docs.giantswarm.io/reference/cp-k8s-api/awsconfigs.provider.giantswarm.io/
	// after the docs have been published for the first time.
	awsConfigDocumentationLink = "https://docs.giantswarm.io/reference/cp-k8s-api/"
)

const awsConfigCRDYAML = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: awsconfigs.provider.giantswarm.io
spec:
  conversion:
    strategy: None
  group: provider.giantswarm.io
  names:
    kind: AWSConfig
    listKind: AWSConfigList
    plural: awsconfigs
    singular: awsconfig
  preserveUnknownFields: true
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    subresources:
      status: {}
    schema:
      openAPIV3Schema:
        description: |
          Together with AWSClusterConfig, defines a tenant cluster in a
          Giant Swarm installation in releases before v10.0.1.
          Reconciled by aws-operator.
        type: object
        properties:
          spec:
            type: object
            properties:
              aws:
                description: AWS-specific configuration.
                type: object
                properties:
                  api:
                    description: |
                      Configures the tenant cluster Kubernetes API.
                      Deprecated since aws-operator v12.
                    type: object
                    properties:
                      elb:
                        description: |
                          Configures the Elastic Load Balancer for the
                          tenant cluster Kubernetes API.
                        type: object
                        properties:
                          idleTimeoutSeconds:
                            description: |
                              Seconds to keep an idle connection open.
                            type: int
                      hostedZones:
                        description: TODO
                        type: string
                  availabilityZones:
                    description: |
                      Number of availability zones to use for the tenant cluster's worker nodes.
                      There are limitations on availability zone settings due to binary IP range
                      splitting and provider specific region capabilities. When for instance choosing
                      3 availability zones, the configured IP range will be split into 4 ranges and
                      thus one of it will not be able to be utilized. Such limitations have to be considered
                      when designing the network topology and configuring tenant cluster HA via availability
                      zones.

                      The selection and usage of the actual availability zones for the created
                      tenant cluster is randomized. In case there are 4 availability zones
                      provided in the used region and the user selects 2 availability zones, the
                      actually used availability zones in which tenant cluster workload is put
                      into will tend to be different across tenant cluster creations. This is
                      done in order to provide more HA during single availability zone failures.
                      In case a specific availability zone fails, not all tenant clusters will be
                      affected due to the described selection process.
                    type: int
                  az:
                    description: |
                      Deprecated way to define the availability zone to use
                      by the tenant cluster.
                    type: string
                  credentialSecret:
                    description: |
                      Defines which Secret resource to use in orde to obtain the IAM role to assume
                      for managing resources for this tenant cluster on AWS. This allows to
                      run different tenant clusters of the same installation in different
                      AWS accounts.
                    type: object
                    properties:
                      name:
                        description: Name of the Secret resource.
                        type: string
                      namespace:
                        description: Namespace of the Secret resource.
                        type: string
                  etcd:
                    description: |
                      Configures the Etcd cluster storing the tenant cluster's data.
                      Deprecated since aws-operator v12.
                    type: object
                    properties:
                      elb:
                        description: |
                          Elastic Load Balancer configuration.
                        type: object
                        properties:
                          idleTimeoutSeconds:
                            description:  |
                              Seconds to keep an idle connection open.
                            type: int
                      hostedZones:
                        description: TODO
                        type: string
                  hostedZones:
                    description: |
                      HostedZones is AWS hosted zones names in the host cluster account.
                      For each zone there will be "CLUSTER_ID.k8s" NS record created in
                      the host cluster account. Then for each created NS record there will
                      be a zone created in the guest account. After that component-specific
                      records under those zones:

                      - api.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.API.Name }}
                      - etcd.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Etcd.Name }}
                      - ingress.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Ingress.Name }}
                      - *.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Ingress.Name }}
                    type: object
                    properties:
                      api:
                        description: |
                          DNS zone configuration for the Kubernetes API.
                        type: object
                        properties:
                          name:
                            description: Zone name.
                            type: string
                      etcd:
                        description: |
                          DNS zone configuration for Etcd.
                        type: object
                        properties:
                          name:
                            description: Zone name.
                            type: string
                      ingress:
                        description: |
                          DNS zone configuration for Ingress.
                        type: object
                        properties:
                          name:
                            description: Zone name.
                            type: string
                  ingress:
                    description: |
                      Ingress configuration.
                      Deprecated since aws-operator v12.
                    type: object
                    properties:
                      elb:
                        description: |
                          Configuration of the Elastic Load Balancer for Ingress.
                        type: object
                        properties:
                          idleTimeoutSeconds:
                            description:  |
                              Seconds to keep an idle connection open.
                            type: int
                      hostedZones:
                        description: TODO
                        type: string
                  masters:
                    description: |
                      Configuration of the master nodes.
                    type: array
                    items:
                      description: |
                        Configuration for each individual master node.
                      type: object
                      properties:
                        dockerVolumeSizeGB:
                          description: TODO (appears unused)
                          type: int
                        imageID:
                          description: |
                            Amazon Machine Image (AMI) identifier to use as a base for
                            the EC2 instance of this master node.
                          type: string
                        instanceType:
                          description: |
                            Name of the AWS EC2 instance type to use for this master node.
                          type: string
                  region:
                    description: |
                      Name of the AWS region the tenant cluster is running in.
                    type: string
                  vpc:
                    description: |
                      Configuration of the Virtual Private Cloud (VPC) to use for
                      the tenant cluster.
                    type: object
                    properties:
                      cidr:
                        description: TODO
                        type: string
                      peerId:
                        description: TODO
                        type: string
                      privateSubnetCidr:
                        description: TODO
                        type: string
                      publicSubnetCidr:
                        description: TODO
                        type: string
                      routeTableNames:
                        description: TODO
                        type: array
                        items:
                          type: string
                  workers:
                    description: |
                      Configuration of worker nodes.
                    type: array
                    items:
                      description: |
                        Configuration of an individual worker node. Each worker node
                        is represented by one item.
                      type: object
                      properties:
                        dockerVolumeSizeGB:
                          description: TODO
                          type: int
                        imageID:
                          description: |
                            Amazon Machine Image (AMI) identifier to use as a base for
                            the EC2 instance of this master node.
                          type: string
                        instanceType:
                          description: |
                            Name of the AWS EC2 instance type to use for this master node.
                          type: string
              cluster:
                description: |
                  Cluster configuration parts that are not meant to be specific to AWS
                  and thus might look similar or identical on other providers.
                type: object
                properties:
                  calico:
                    description: |
                      Configuration for the Project Calico Container Network Interface
                      (CNI).
                    type: object
                    properties:
                      cidr:
                        description: |
                          Subnet size, expresses as the count of leading 1 bits in the subnet
                          mask of this subnet. In other words, in CIDR notation, the integer
                          behind the slash.
                        type: int
                      mtu:
                        description: |
                          Size of the maximum transition unit (MTU) in bytes.
                        type: int
                      subnet:
                        description: |
                          Subnet IPv4 address. In other words, in CIDR notation, the
                          part before the slash.
                        type: string
                  customer:
                    description: |
                      Information on the Giant Swarm customer owning the
                      tenant cluster.
                    type: object
                    properties:
                      id:
                        description: |
                          Unique customer identifier.
                        type: string
                  docker:
                    description: |
                      Configuration for Docker.
                    type: object
                    properties:
                      daemon:
                        description: |
                          Configuration for the Docker daemon.
                        type: object
                        properties:
                          cidr:
                            description: |
                              CIDR notation for the subnet to use for Docker
                              networking.
                            type: string
                  etcd:
                    description: |
                      Configuration for Etcd.
                    type: object
                    properties:
                      altNames:
                        description: TODO
                        type: string
                      domain:
                        description: |
                          Domain name to use for Etcd.
                        type: string
                      port:
                        description: |
                          Port number Etcd is listening on.
                        type: int
                      prefix:
                        description: |
                          Prefix to prepend to all Etcd keys.
                        type: string
                  id:
                    description: |
                      Cluster identifier, unique within the installation.
                      The identifier is expected to be exactly 5 characters
                      long, with characters from the range a-z and 0-9.
                    type: string
                  kubernetes:
                    description: |
                      Various Kubernetes configuration items.
                    type: object
                    properties:
                      api:
                        description: |
                          Configuration for the Kubernetes api-server.
                        type: object
                        properties:
                          clusterIPRange:
                            description: |
                              IP range to use for ClusterIP, in CIDR notation.
                            type: string
                          domain:
                            description: |
                              Fully qualified host name of the API server.
                            type: string
                          securePort:
                            description: |
                              Port number for HTTPS access to the API.
                            type: int
                      cloudProvider:
                        description: Name of the cloud provider. Must be 'aws'.
                        type: string
                      dns:
                        description: DNS configuration.
                        type: object
                        properties:
                          ip:
                            description: TODO
                            type: string
                      domain:
                        description: Domain name to use internally within the cluster.
                        type: string
                      ingressController:
                        description: |
                          Configuration of the Ingress Controller.
                        type: object
                        properties:
                          docker:
                            description: Ingress Controller Docker configuration.
                            type: object
                            properties:
                              image:
                                description: Docker image for the Ingress Controller.
                                type: string
                          domain:
                            description: |
                              Fully qualified host name the Ingress load balancer is listening on.
                            type: string
                          insecurePort:
                            description: TODO
                            type: int
                          securePort:
                            description: TODO
                            type: int
                          wildcardDomain:
                            description: TODO
                            type: string
                      kubelet:
                        description: TODO
                        type: object
                        properties:
                          altNames:
                            description: |
                              Alternative internal DNS names to access the API server.
                            type: string
                          domain:
                            description: |
                              DNS suffix to use for worker nodes.
                            type: string
                          labels:
                            description: TODO
                            type: string
                          port:
                            description: TODO
                            type: int
                      networkSetup:
                        description: TODO
                        type: object
                        properties:
                          docker:
                            description: TODO
                            type: object
                            properties:
                              image:
                                description: |
                                  Docker image for setting up the network environment.
                                type: string
                          kubeProxy:
                            description: |
                              kubeProxy configuration.
                            type: object
                            properties:
                              conntrackMaxPerCore:
                                description: TODO
                                type: int
                      ssh:
                        description: |
                          SSH access configuration applied to master and worker nodes
                          of the tenant cluster.
                        type: object
                        properties:
                          userList:
                            description: |
                              List of SSH users with access to the nodes.
                            type: array
                            items:
                              description: |
                                An individual SSH user with access.
                              type: object
                              properties:
                                name:
                                  description: |
                                    Posix username to assign to the user logging in with
                                    the given SSH public key.
                                  type: string
                                publicKey:
                                  description: |
                                    SSH public key in Base64 encoding.
                                  type: string
                  masters:
                    description: |
                      Configuration of master nodes.
                    type: array
                    items:
                      description: TODO
                      type: object
                      properties:
                        id:
                          description: TODO
                          type: string
                  scaling:
                    description: |
                      Configuration of worker node scaling. The cluster-autoscaler
                      sets the actual number of worker nodes in the cluster based
                      on the 'min' and 'max' value.
                    type: object
                    properties:
                      min:
                        description: |
                          Minimum amount of worker nodes. Lower limit of the
                          autoscaling range.
                        type: int
                      max:
                        description: |
                          Maximum amount of worker nodes. Upper limit of the
                          autoscaling range.
                        type: int
                  version:
                    description: TODO
                    type: string
                  workers:
                    description: |
                      List of worker nodes with configuration.
                    type: array
                    items:
                      description: |
                        Configuration of an individual worker node.
                      type: object
                      properties:
                        id:
                          description: |
                            Unique identifier of the worker node.
                          type: string
              versionBundle:
                description: TODO
                type: object
                properties:
                  version:
                    description: TODO
                    type: string
`

var awsConfigCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(awsConfigCRDYAML), &awsConfigCRD)
	if err != nil {
		panic(err)
	}
}

// NewAWSConfigCRD returns a new custom resource definition for AWSConfig. This
// might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: awsconfigs.provider.giantswarm.io
//     spec:
//       group: provider.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: AWSConfig
//         plural: awsconfigs
//         singular: awsconfig
//       subresources:
//         status: {}
//
func NewAWSConfigCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return awsConfigCRD.DeepCopy()
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
				crDocsAnnotation: awsConfigDocumentationLink,
			},
		},
		TypeMeta: NewAWSClusterTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AWSConfigSpec   `json:"spec"`
	Status            AWSConfigStatus `json:"status" yaml:"status"`
}

type AWSConfigSpec struct {
	Cluster       Cluster                    `json:"cluster" yaml:"cluster"`
	AWS           AWSConfigSpecAWS           `json:"aws" yaml:"aws"`
	VersionBundle AWSConfigSpecVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type AWSConfigSpecAWS struct {
	API AWSConfigSpecAWSAPI `json:"api" yaml:"api"`
	// TODO remove the deprecated AZ field due to AvailabilityZones.
	//
	//     https://github.com/giantswarm/giantswarm/issues/4507
	//
	AZ string `json:"az" yaml:"az"`
	// AvailabilityZones is the number of AWS availability zones used to spread
	// the tenant cluster's worker nodes across. There are limitations on
	// availability zone settings due to binary IP range splitting and provider
	// specific region capabilities. When for instance choosing 3 availability
	// zones, the configured IP range will be split into 4 ranges and thus one of
	// it will not be able to be utilized. Such limitations have to be considered
	// when designing the network topology and configuring tenant cluster HA via
	// AvailabilityZones.
	//
	// The selection and usage of the actual availability zones for the created
	// tenant cluster is randomized. In case there are 4 availability zones
	// provided in the used region and the user selects 2 availability zones, the
	// actually used availability zones in which tenant cluster workload is put
	// into will tend to be different across tenant cluster creations. This is
	// done in order to provide more HA during single availability zone failures.
	// In case a specific availability zone fails, not all tenant clusters will be
	// affected due to the described selection process.
	AvailabilityZones int                  `json:"availabilityZones" yaml:"availabilityZones"`
	CredentialSecret  CredentialSecret     `json:"credentialSecret" yaml:"credentialSecret"`
	Etcd              AWSConfigSpecAWSEtcd `json:"etcd" yaml:"etcd"`

	// HostedZones is AWS hosted zones names in the host cluster account.
	// For each zone there will be "CLUSTER_ID.k8s" NS record created in
	// the host cluster account. Then for each created NS record there will
	// be a zone created in the guest account. After that component
	// specific records under those zones:
	//	- api.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.API.Name }}
	//	- etcd.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Etcd.Name }}
	//	- ingress.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Ingress.Name }}
	//	- *.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Ingress.Name }}
	HostedZones AWSConfigSpecAWSHostedZones `json:"hostedZones" yaml:"hostedZones"`

	Ingress AWSConfigSpecAWSIngress `json:"ingress" yaml:"ingress"`
	Masters []AWSConfigSpecAWSNode  `json:"masters" yaml:"masters"`
	Region  string                  `json:"region" yaml:"region"`
	VPC     AWSConfigSpecAWSVPC     `json:"vpc" yaml:"vpc"`
	Workers []AWSConfigSpecAWSNode  `json:"workers" yaml:"workers"`
}

// AWSConfigSpecAWSAPI deprecated since aws-operator v12 resources.
type AWSConfigSpecAWSAPI struct {
	HostedZones string                 `json:"hostedZones" yaml:"hostedZones"`
	ELB         AWSConfigSpecAWSAPIELB `json:"elb" yaml:"elb"`
}

// AWSConfigSpecAWSAPIELB deprecated since aws-operator v12 resources.
type AWSConfigSpecAWSAPIELB struct {
	IdleTimeoutSeconds int `json:"idleTimeoutSeconds" yaml:"idleTimeoutSeconds"`
}

// AWSConfigSpecAWSEtcd deprecated since aws-operator v12 resources.
type AWSConfigSpecAWSEtcd struct {
	HostedZones string                  `json:"hostedZones" yaml:"hostedZones"`
	ELB         AWSConfigSpecAWSEtcdELB `json:"elb" yaml:"elb"`
}

// AWSConfigSpecAWSEtcdELB deprecated since aws-operator v12 resources.
type AWSConfigSpecAWSEtcdELB struct {
	IdleTimeoutSeconds int `json:"idleTimeoutSeconds" yaml:"idleTimeoutSeconds"`
}

type AWSConfigSpecAWSHostedZones struct {
	API     AWSConfigSpecAWSHostedZonesZone `json:"api" yaml:"api"`
	Etcd    AWSConfigSpecAWSHostedZonesZone `json:"etcd" yaml:"etcd"`
	Ingress AWSConfigSpecAWSHostedZonesZone `json:"ingress" yaml:"ingress"`
}

type AWSConfigSpecAWSHostedZonesZone struct {
	Name string `json:"name" yaml:"name"`
}

// AWSConfigSpecAWSIngress deprecated since aws-operator v12 resources.
type AWSConfigSpecAWSIngress struct {
	HostedZones string                     `json:"hostedZones" yaml:"hostedZones"`
	ELB         AWSConfigSpecAWSIngressELB `json:"elb" yaml:"elb"`
}

// AWSConfigSpecAWSIngressELB deprecated since aws-operator v12 resources.
type AWSConfigSpecAWSIngressELB struct {
	IdleTimeoutSeconds int `json:"idleTimeoutSeconds" yaml:"idleTimeoutSeconds"`
}

type AWSConfigSpecAWSNode struct {
	ImageID            string `json:"imageID" yaml:"imageID"`
	InstanceType       string `json:"instanceType" yaml:"instanceType"`
	DockerVolumeSizeGB int    `json:"dockerVolumeSizeGB" yaml:"dockerVolumeSizeGB"`
}

type AWSConfigSpecAWSVPC struct {
	CIDR              string   `json:"cidr" yaml:"cidr"`
	PrivateSubnetCIDR string   `json:"privateSubnetCidr" yaml:"privateSubnetCidr"`
	PublicSubnetCIDR  string   `json:"publicSubnetCidr" yaml:"publicSubnetCidr"`
	RouteTableNames   []string `json:"routeTableNames" yaml:"routeTableNames"`
	PeerID            string   `json:"peerId" yaml:"peerId"`
}

type AWSConfigSpecVersionBundle struct {
	Version string `json:"version" yaml:"version"`
}

type AWSConfigStatus struct {
	AWS     AWSConfigStatusAWS `json:"aws" yaml:"aws"`
	Cluster StatusCluster      `json:"cluster" yaml:"cluster"`
}

type AWSConfigStatusAWS struct {
	AvailabilityZones []AWSConfigStatusAWSAvailabilityZone `json:"availabilityZones" yaml:"availabilityZones"`
	AutoScalingGroup  AWSConfigStatusAWSAutoScalingGroup   `json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

type AWSConfigStatusAWSAutoScalingGroup struct {
	Name string `json:"name"`
}

type AWSConfigStatusAWSAvailabilityZone struct {
	Name   string                                   `json:"name" yaml:"name"`
	Subnet AWSConfigStatusAWSAvailabilityZoneSubnet `json:"subnet" yaml:"subnet"`
}

type AWSConfigStatusAWSAvailabilityZoneSubnet struct {
	Private AWSConfigStatusAWSAvailabilityZoneSubnetPrivate `json:"private" yaml:"private"`
	Public  AWSConfigStatusAWSAvailabilityZoneSubnetPublic  `json:"public" yaml:"public"`
}

type AWSConfigStatusAWSAvailabilityZoneSubnetPrivate struct {
	CIDR string `json:"cidr" yaml:"cidr"`
}

type AWSConfigStatusAWSAvailabilityZoneSubnetPublic struct {
	CIDR string `json:"cidr" yaml:"cidr"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AWSConfig `json:"items"`
}
