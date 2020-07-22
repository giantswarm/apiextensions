package annotation

// MachinePoolName is the node pool annotation where human-friendly node pool
// name set by the customer is stored.
const MachinePoolName = "machine-pool.giantswarm.io/name"

// MachinePoolSubnet is the node pool annotation which contains the name of the
// subnet where the tenant cluster node pool is deployed. This name should be
// equivalent to the tenant cluster node pool ID. E.g.
// machine-pool.giantswarm.io/subnet=de19f-h94vd.
const MachinePoolSubnet = "machine-pool.giantswarm.io/subnet"
