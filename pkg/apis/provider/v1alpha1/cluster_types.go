package v1alpha1

// +k8s:openapi-gen=true
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

// +k8s:openapi-gen=true
type ClusterCalico struct {
	CIDR   int    `json:"cidr"`
	MTU    int    `json:"mtu"`
	Subnet string `json:"subnet"`
}

// +k8s:openapi-gen=true
type ClusterCustomer struct {
	ID string `json:"id"`
}

// +k8s:openapi-gen=true
type ClusterDocker struct {
	Daemon ClusterDockerDaemon `json:"daemon"`
}

// +k8s:openapi-gen=true
type ClusterDockerDaemon struct {
	CIDR string `json:"cidr"`
}

// +k8s:openapi-gen=true
type ClusterEtcd struct {
	AltNames string `json:"altNames"`
	Domain   string `json:"domain"`
	Port     int    `json:"port"`
	Prefix   string `json:"prefix"`
}

// +k8s:openapi-gen=true
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

// +k8s:openapi-gen=true
type ClusterKubernetesAPI struct {
	ClusterIPRange string `json:"clusterIPRange"`
	Domain         string `json:"domain"`
	SecurePort     int    `json:"securePort"`
}

// +k8s:openapi-gen=true
type ClusterKubernetesDNS struct {
	IP string `json:"ip"`
}

// +k8s:openapi-gen=true
type ClusterKubernetesIngressController struct {
	Docker         ClusterKubernetesIngressControllerDocker `json:"docker"`
	Domain         string                                   `json:"domain"`
	WildcardDomain string                                   `json:"wildcardDomain"`
	InsecurePort   int                                      `json:"insecurePort"`
	SecurePort     int                                      `json:"securePort"`
}

// +k8s:openapi-gen=true
type ClusterKubernetesIngressControllerDocker struct {
	Image string `json:"image"`
}

// +k8s:openapi-gen=true
type ClusterKubernetesKubelet struct {
	AltNames string `json:"altNames"`
	Domain   string `json:"domain"`
	Labels   string `json:"labels"`
	Port     int    `json:"port"`
}

// +k8s:openapi-gen=true
type ClusterKubernetesNetworkSetup struct {
	Docker    ClusterKubernetesNetworkSetupDocker    `json:"docker"`
	KubeProxy ClusterKubernetesNetworkSetupKubeProxy `json:"kubeProxy"`
}

// ClusterKubernetesNetworkSetupKubeProxy describes values passed to the kube-proxy running in a workload cluster.
// +k8s:openapi-gen=true
type ClusterKubernetesNetworkSetupKubeProxy struct {
	// Maximum number of NAT connections to track per CPU core (0 to leave the limit as-is and ignore conntrack-min).
	// Passed to kube-proxy as --conntrack-max-per-core.
	ConntrackMaxPerCore int `json:"conntrackMaxPerCore"`
}

// +k8s:openapi-gen=true
type ClusterKubernetesNetworkSetupDocker struct {
	Image string `json:"image"`
}

// +k8s:openapi-gen=true
type ClusterKubernetesSSH struct {
	UserList []ClusterKubernetesSSHUser `json:"userList"`
}

// +k8s:openapi-gen=true
type ClusterKubernetesSSHUser struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

// +k8s:openapi-gen=true
type ClusterNode struct {
	ID string `json:"id"`
}

// +k8s:openapi-gen=true
type ClusterScaling struct {
	// Max defines maximum number of worker nodes guest cluster is allowed to have.
	Max int `json:"max"`
	// Min defines minimum number of worker nodes required to be present in guest cluster.
	Min int `json:"min"`
}
