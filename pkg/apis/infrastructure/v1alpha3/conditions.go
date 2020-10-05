package v1alpha3

import (
	capiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
)

// AzureMachinePool conditions
const (
	DeploymentSucceededCondition capiv1alpha3.ConditionType = "DeploymentSucceeded"
	VMSSReadyCondition capiv1alpha3.ConditionType = "VMSSReady"
	SubnetReadyCondition capiv1alpha3.ConditionType = "SubnetReady"
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
