package v1alpha1

import (
	"net"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AWSSpec `json:"spec"`
}

type AWSSpec struct {
	Cluster       AWSSpecCluster       `json:"cluster" yaml:"cluster"`
	AWS           AWSSpecAWS           `json:"aws" yaml:"aws"`
	VersionBundle AWSSpecVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type AWSSpecCluster struct {
	Calico     AWSSpecClusterCalico   `json:"calico" yaml:"calico"`
	Customer   AWSSpecClusterCustomer `json:"customer" yaml:"customer"`
	Docker     AWSSpecClusterDocker   `json:"docker" yaml:"docker"`
	Etcd       AWSSpecClusterEtcd     `json:"etcd" yaml:"etcd"`
	ID         string                 `json:"id" yaml:"id"`
	Kubernetes spec.Kubernetes        `json:"kubernetes" yaml:"kubernetes"`
	Masters    []AWSSpecClusterNode   `json:"masters" yaml:"masters"`
	Vault      AWSSpecClusterVault    `json:"vault" yaml:"vault"`
	Workers    []AWSSpecClusterNode   `json:"workers" yaml:"workers"`
}

type AWSSpecClusterCalico struct {
	CIDR   int    `json:"cidr" yaml:"cidr"`
	Domain string `json:"domain" yaml:"domain"`
	MTU    int    `json:"mtu" yaml:"mtu"`
	Subnet string `json:"subnet" yaml:"subnet"`
}

type AWSSpecClusterCustomer struct {
	ID string `json:"id" yaml:"id"`
}

type AWSSpecClusterDocker struct {
	Daemon AWSSpecClusterDockerDaemon `json:"daemon" yaml:"daemon"`
}

type AWSSpecClusterDockerDaemon struct {
	CIDR      string `json:"cidr" yaml:"cidr"`
	ExtraArgs string `json:"extraArgs" yaml:"extraArgs"`
}

type AWSSpecClusterEtcd struct {
	AltNames string `json:"altNames" yaml:"altNames"`
	Domain   string `json:"domain" yaml:"domain"`
	Port     int    `json:"port" yaml:"port"`
	Prefix   string `json:"prefix" yaml:"prefix"`
}

type AWSSpecClusterKubernetes struct {
	API               AWSSpecClusterKubernetesAPI               `json:"api" yaml:"api"`
	DNS               AWSSpecClusterKubernetesDNS               `json:"dns" yaml:"dns"`
	Domain            string                                    `json:"domain" yaml:"domain"`
	Hyperkube         AWSSpecClusterKubernetesHyperkube         `json:"hyperkube" yaml:"hyperkube"`
	IngressController AWSSpecClusterKubernetesIngressController `json:"ingressController" yaml:"ingressController"`
	Kubelet           AWSSpecClusterKubernetesKubelet           `json:"kubelet" yaml:"kubelet"`
	NetworkSetup      AWSSpecClusterKubernetesNetworkSetup      `json:"networkSetup" yaml:"networkSetup"`
	SSH               AWSSpecClusterKubernetesSSH               `json:"ssh" yaml:"ssh"`
}

type AWSSpecClusterKubernetesAPI struct {
	AltNames       string `json:"altNames" yaml:"altNames"`
	ClusterIPRange string `json:"clusterIPRange" yaml:"clusterIPRange"`
	Domain         string `json:"domain" yaml:"domain"`
	IP             net.IP `json:"ip" yaml:"ip"`
	InsecurePort   int    `json:"insecurePort" yaml:"insecurePort"`
	SecurePort     int    `json:"securePort" yaml:"securePort"`
}

type AWSSpecClusterKubernetesDNS struct {
	IP net.IP `json:"ip" yaml:"ip"`
}

type AWSSpecClusterKubernetesHyperkube struct {
	Docker AWSSpecClusterKubernetesHyperkubeDocker `json:"docker" yaml:"docker"`
}

type AWSSpecClusterKubernetesHyperkubeDocker struct {
	Image string `json:"image" yaml:"image"`
}

type AWSSpecClusterKubernetesIngressController struct {
	Docker         AWSSpecClusterKubernetesIngressControllerDocker `json:"docker" yaml:"docker"`
	Domain         string                                          `json:"domain" yaml:"domain"`
	WildcardDomain string                                          `json:"wildcardDomain" yaml:"wildcardDomain"`
	InsecurePort   int                                             `json:"insecurePort" yaml:"insecurePort"`
	SecurePort     int                                             `json:"securePort" yaml:"securePort"`
}

type AWSSpecClusterKubernetesIngressControllerDocker struct {
	Image string `json:"image" yaml:"image"`
}

type AWSSpecClusterKubernetesKubelet struct {
	AltNames string `json:"altNames" yaml:"altNames"`
	Domain   string `json:"domain" yaml:"domain"`
	Labels   string `json:"labels" yaml:"labels"`
	Port     int    `json:"port" yaml:"port"`
}

type AWSSpecClusterKubernetesNetworkSetup struct {
	Docker AWSSpecClusterKubernetesNetworkSetupDocker `json:"docker" yaml:"docker"`
}

type AWSSpecClusterKubernetesNetworkSetupDocker struct {
	Image string `json:"image" yaml:"image"`
}

type AWSSpecClusterKubernetesSSH struct {
	UserList []AWSSpecClusterKubernetesSSHUser `json:"userList" yaml:"userList"`
}

type AWSSpecClusterKubernetesSSHUser struct {
	Name      string `json:"name" yaml:"name"`
	PublicKey string `json:"publicKey" yaml:"publicKey"`
}

type AWSSpecClusterNode struct {
	ID string `json:"id" yaml:"id"`
}

type AWSSpecClusterVault struct {
	Address string `json:"address" yaml:"address"`
	Token   string `json:"token" yaml:"token"`
}

type AWSSpecAWS struct {
	API     AWSSpecAWSAPI     `json:"api" yaml:"api"`
	AZ      string            `json:"az" yaml:"az"`
	Etcd    AWSSpecAWSEtcd    `json:"etcd" yaml:"etcd"`
	Ingress AWSSpecAWSIngress `json:"ingress" yaml:"ingress"`
	Masters []AWSSpecAWSNode  `json:"masters" yaml:"masters"`
	Region  string            `json:"region" yaml:"region"`
	VPC     AWSSpecAWSVPC     `json:"vpc" yaml:"vpc"`
	Workers []AWSSpecAWSNode  `json:"workers" yaml:"workers"`
}

type AWSSpecAWSAPI struct {
	HostedZones string           `json:"hostedZones" yaml:"hostedZones"`
	ELB         AWSSpecAWSAPIELB `json:"elb" yaml:"elb"`
}

type AWSSpecAWSAPIELB struct {
	IdleTimeoutSeconds int `json:"idleTimeoutSeconds" yaml:"idleTimeoutSeconds"`
}

type AWSSpecAWSEtcd struct {
	HostedZones string            `json:"hostedZones" yaml:"hostedZones"`
	ELB         AWSSpecAWSEtcdELB `json:"elb" yaml:"elb"`
}

type AWSSpecAWSEtcdELB struct {
	IdleTimeoutSeconds int `json:"idleTimeoutSeconds" yaml:"idleTimeoutSeconds"`
}

type AWSSpecAWSIngress struct {
	HostedZones string               `json:"hostedZones" yaml:"hostedZones"`
	ELB         AWSSpecAWSIngressELB `json:"elb" yaml:"elb"`
}

type AWSSpecAWSIngressELB struct {
	IdleTimeoutSeconds int `json:"idleTimeoutSeconds" yaml:"idleTimeoutSeconds"`
}

type AWSSpecAWSNode struct {
	ImageID      string `json:"imageID" yaml:"imageID"`
	InstanceType string `json:"instanceType" yaml:"instanceType"`
}

type AWSSpecAWSVPC struct {
	CIDR              string   `json:"cidr" yaml:"cidr"`
	PrivateSubnetCIDR string   `json:"privateSubnetCidr" yaml:"privateSubnetCidr"`
	PublicSubnetCIDR  string   `json:"publicSubnetCidr" yaml:"publicSubnetCidr"`
	RouteTableNames   []string `json:"routeTableNames" yaml:"routeTableNames"`
	PeerID            string   `json:"peerId" yaml:"peerId"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AWS `json:"items"`
}
