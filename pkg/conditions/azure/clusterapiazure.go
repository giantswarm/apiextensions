package azure

import (
	capi "sigs.k8s.io/cluster-api/api/v1beta1"
)

// Common AzureCluster and AzureMachinePool conditions
const (
	// DeploymentSucceededCondition is true when deployments of all Azure
	// resources, required by a CR, are in succeeded provisioning state.
	DeploymentSucceededCondition capi.ConditionType = "DeploymentSucceeded"
)

// AzureCluster conditions
const (
	ResourceGroupReadyCondition  capi.ConditionType = "ResourceGroupReady"
	StorageAccountReadyCondition capi.ConditionType = "StorageAccountReady"
	VirtualNetworkReadyCondition capi.ConditionType = "VirtualNetworkReady"
	VPNGatewayReadyCondition     capi.ConditionType = "VPNGatewayReady"
)

// AzureMachinePool and AzureMachine conditions
const (
	VMSSReadyCondition   capi.ConditionType = "VMSSReady"
	SubnetReadyCondition capi.ConditionType = "SubnetReady"
)

// Azure VMSS instance statuses used as condition reasons
const (
	// VMProvisioningStateSucceededReason: the user-initiated actions have completed,
	// ConditionSeverity is Info.
	VMProvisioningStateSucceededReason = "VMProvisioningStateSucceeded"

	// VMProvisioningStateCreatingReason: the user-initiated VM (or VMSS
	// instance) creation, ConditionSeverity is Info.
	VMProvisioningStateCreatingReason = "VMProvisioningStateCreating"

	// VMProvisioningStateCreatingOSProvisioningInProgressReason: the user-initiated VM (or VMSS
	// instance) creation, the VM is running, and installation of guest OS is in progress,
	// ConditionSeverity is Info.
	VMProvisioningStateCreatingOSProvisioningInProgressReason = "VMProvisioningStateCreatingOSProvisioningInProgress"

	// VMProvisioningStateCreatingOSProvisioningCompleteReason: the user-initiated VM (or VMSS
	// instance) creation, ConditionSeverity is Info.
	VMProvisioningStateCreatingOSProvisioningCompleteReason = "VMProvisioningStateCreatingOSProvisioningComplete"

	// VMProvisioningStateUpdatingReason: the user-initiated VM (or VMSS instance) update,
	// ConditionSeverity is Info.
	VMProvisioningStateUpdatingReason = "VMProvisioningStateUpdating"

	// VMProvisioningStateDeletingReason: the user-initiated VM (or VMSS instance) deletion,
	// ConditionSeverity is Info.
	VMProvisioningStateDeletingReason = "VMProvisioningStateDeleting"

	// VMProvisioningStateFailedReason: failed operation. Refer to the error codes to get more
	// information and possible solutions, ConditionSeverity is Error for a node pool, and Warning
	// for a cluster.
	VMProvisioningStateFailedReason = "VMProvisioningStateFailed"

	// VMPowerStateStartingReason: the VM is starting up. ConditionSeverity is Info.
	VMPowerStateStartingReason = "VMPowerStateStarting"

	// VMPowerStateRunningReason represents a normal working state for a VM, ConditionSeverity is
	// Info.
	VMPowerStateRunningReason = "VMPowerStateRunning"

	// VMPowerStateStoppingReason signals a transitional state, when completed, it will show as
	// Stopped. ConditionSeverity is Info.
	VMPowerStateStoppingReason = "VMPowerStateStopping"

	// VMPowerStateStoppedReason signals that the VM has been shut down from within the guest OS or
	// using the PowerOff APIs. Hardware is still allocated to the VM and it remains on the host.
	// ConditionSeverity is Info.
	VMPowerStateStoppedReason = "VMPowerStateStopped"

	// VMPowerStateDeallocatingReason: ransitional state. When completed, the VM will show as
	// Deallocated. ConditionSeverity is Info.
	VMPowerStateDeallocatingReason = "VMPowerStateDeallocating"

	// VMPowerStateDeallocatedReason: The VM has been stopped successfully and removed from the
	// host. ConditionSeverity is Info.
	VMPowerStateDeallocatedReason = "VMPowerStateDeallocated"
)
