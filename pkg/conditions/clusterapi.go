package conditions

import (
	"encoding/json"

	"github.com/giantswarm/microerror"
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

// Upgrading reasons
const (
	UpgradeCompletedReason  = "UpgradeCompleted"
	UpgradeNotStartedReason = "UpgradeNotStarted"
)

type UpgradingConditionMessage struct {
	Message        string `json:"message"`
	ReleaseVersion string `json:"release_version"`
}

// SerializeUpgradingConditionMessage converts specified message object into a
// JSON string.
func SerializeUpgradingConditionMessage(message UpgradingConditionMessage) (string, error) {
	messageJson, err := json.Marshal(message)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return string(messageJson), nil
}

// DeserializeUpgradingConditionMessage parses specified JSON string and
// returns message UpgradingConditionMessage struct.
func DeserializeUpgradingConditionMessage(messageJson string) (UpgradingConditionMessage, error) {
	var message UpgradingConditionMessage

	err := json.Unmarshal([]byte(messageJson), &message)
	if err != nil {
		return UpgradingConditionMessage{}, microerror.Mask(err)
	}

	return message, nil
}
