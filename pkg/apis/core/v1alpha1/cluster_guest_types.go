package v1alpha1

// +k8s:openapi-gen=true
type ClusterGuestConfig struct {
	AvailabilityZones int `json:"availabilityZones,omitempty"`
	// DNSZone for guest cluster is supplemented with host prefixes for
	// specific services such as Kubernetes API or Etcd. In general this DNS
	// Zone should start with "k8s" like for example
	// "k8s.cluster.example.com.".
	DNSZone        string                            `json:"dnsZone"`
	ID             string                            `json:"id"`
	Name           string                            `json:"name,omitempty"`
	Owner          string                            `json:"owner,omitempty"`
	ReleaseVersion string                            `json:"releaseVersion,omitempty"`
	VersionBundles []ClusterGuestConfigVersionBundle `json:"versionBundles,omitempty"`
}

// +k8s:openapi-gen=true
type ClusterGuestConfigVersionBundle struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
