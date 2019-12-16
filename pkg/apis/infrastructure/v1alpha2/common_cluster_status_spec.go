package v1alpha2

// CommonClusterStatusGetSetter provides abstract way to manipulate common
// provider independent cluster status field in provider CR's status.
type CommonClusterStatusGetSetter interface {
	GetCommonClusterStatus() CommonClusterStatus
	SetCommonClusterStatus(ccs CommonClusterStatus)
}
