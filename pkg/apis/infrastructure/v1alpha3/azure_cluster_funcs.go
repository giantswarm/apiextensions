package v1alpha3

func (c *AzureCluster) GetCommonClusterStatus() CommonClusterStatus {
	return c.Status.Cluster
}

func (c *AzureCluster) SetCommonClusterStatus(s CommonClusterStatus) {
	c.Status.Cluster = s
}
