package v1alpha1

const (
	ExternalTrafficPolicyCluster = "Cluster"
	ExternalTrafficPolicyLocal   = "Local"

	LoadBalancerTypeInternal = "internal"
	LoadBalancerTypePublic   = "public"
)

type Cluster struct {
	Calico     ClusterCalico     `json:"calico"`
	Customer   ClusterCustomer   `json:"customer"`
	Docker     ClusterDocker     `json:"docker"`
	Etcd       ClusterEtcd       `json:"etcd"`
	ID         string            `json:"id"`
	Kubernetes ClusterKubernetes `json:"kubernetes"`
	Masters    []ClusterNode     `json:"masters"`
	Scaling    ClusterScaling    `json:"scaling"`

	// Version is DEPRECATED and should just be dropped.
	Version string `json:"version"`

	Workers []ClusterNode `json:"workers,omitempty"`
}

type ClusterCalico struct {
	CIDR   int    `json:"cidr"`
	MTU    int    `json:"mtu"`
	Subnet string `json:"subnet"`
}

type ClusterCustomer struct {
	ID string `json:"id"`
}

type ClusterDocker struct {
	Daemon ClusterDockerDaemon `json:"daemon"`
}

type ClusterDockerDaemon struct {
	CIDR string `json:"cidr"`
}

type ClusterEtcd struct {
	AltNames string `json:"altNames"`
	Domain   string `json:"domain"`
	Port     int    `json:"port"`
	Prefix   string `json:"prefix"`
}

type ClusterKubernetes struct {
	API               ClusterKubernetesAPI               `json:"api"`
	CloudProvider     string                             `json:"cloudProvider"`
	DNS               ClusterKubernetesDNS               `json:"dns"`
	Domain            string                             `json:"domain"`
	IngressController ClusterKubernetesIngressController `json:"ingressController"`
	Kubelet           ClusterKubernetesKubelet           `json:"kubelet"`
	NetworkSetup      ClusterKubernetesNetworkSetup      `json:"networkSetup"`
	SSH               ClusterKubernetesSSH               `json:"ssh"`
}

type ClusterKubernetesAPI struct {
	ClusterIPRange string `json:"clusterIPRange"`
	Domain         string `json:"domain"`
	SecurePort     int    `json:"securePort"`
}

type ClusterKubernetesDNS struct {
	IP string `json:"ip"`
}

type ClusterKubernetesIngressController struct {
	Docker         ClusterKubernetesIngressControllerDocker `json:"docker"`
	Domain         string                                   `json:"domain"`
	WildcardDomain string                                   `json:"wildcardDomain"`
	InsecurePort   int                                      `json:"insecurePort"`
	SecurePort     int                                      `json:"securePort"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=public
	// +kubebuilder:validation:Enum=internal;public
	// The LoadBalancerType property allows to choose the type of the LoadBalancer to be deployed.
	// Can be either "internal" or "public" and it is supported on Azure tenant clusters only.
	LoadBalancerType string `json:"loadBalancerType"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=Local
	// +kubebuilder:validation:Enum=Cluster;Local
	// The ExternalTrafficPolicy property is used to set the homonym property of the Kubernetes Service
	// used by the Nginx Ingress Controller.
	ExternalTrafficPolicy string `json:"externalTrafficPolicy"`
}

type ClusterKubernetesIngressControllerDocker struct {
	Image string `json:"image"`
}

type ClusterKubernetesKubelet struct {
	AltNames string `json:"altNames"`
	Domain   string `json:"domain"`
	Labels   string `json:"labels"`
	Port     int    `json:"port"`
}

type ClusterKubernetesNetworkSetup struct {
	Docker    ClusterKubernetesNetworkSetupDocker    `json:"docker"`
	KubeProxy ClusterKubernetesNetworkSetupKubeProxy `json:"kubeProxy"`
}

// ClusterKubernetesNetworkSetupKubeProxy describes values passed to the kube-proxy running in a tenant cluster.
type ClusterKubernetesNetworkSetupKubeProxy struct {
	// Maximum number of NAT connections to track per CPU core (0 to leave the limit as-is and ignore conntrack-min).
	// Passed to kube-proxy as --conntrack-max-per-core.
	ConntrackMaxPerCore int `json:"conntrackMaxPerCore"`
}

type ClusterKubernetesNetworkSetupDocker struct {
	Image string `json:"image"`
}

type ClusterKubernetesSSH struct {
	UserList []ClusterKubernetesSSHUser `json:"userList"`
}

type ClusterKubernetesSSHUser struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

type ClusterNode struct {
	ID string `json:"id"`
}

type ClusterScaling struct {
	// Max defines maximum number of worker nodes guest cluster is allowed to have.
	Max int `json:"max"`
	// Min defines minimum number of worker nodes required to be present in guest cluster.
	Min int `json:"min"`
}
