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
)
