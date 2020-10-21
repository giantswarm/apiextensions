package conditions

import (
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

// Cluster conditions
const (
	// CreatingCondition tells if the cluster is being created.
	CreatingCondition capi.ConditionType = "Creating"

	// UpgradingCondition tells if the cluster is being upgraded to a newer
	// release version.
	UpgradingCondition capi.ConditionType = "Upgrading"

	// ProviderInfrastructureReadyCondition tells if the provider
	// infrastructure defined by the related provider-specific CR is Ready. For
	// example, in Cluster CR it tells if Azure infrastructure defined by the
	// AzureCluster is ready.
	ProviderInfrastructureReadyCondition capi.ConditionType = "ProviderInfrastructureReady"

	// NodePoolsReadyCondition tells if all node pools are ready.
	NodePoolsReadyCondition capi.ConditionType = "NodePoolsReady"
)
