package v1alpha3

func (c *AWSCluster) GetCommonClusterStatus() CommonClusterStatus {
	return c.Status.Cluster
}

func (c *AWSCluster) SetCommonClusterStatus(s CommonClusterStatus) {
	c.Status.Cluster = s
}
