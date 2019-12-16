package v1alpha2

func (c *AWSCluster) GetCommonClusterStatus() CommonClusterStatus {
	return c.Status.CommonClusterStatus
}

func (c *AWSCluster) SetCommonClusterStatus(s CommonClusterStatus) {
	c.Status.CommonClusterStatus = s
}
