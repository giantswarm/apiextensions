package annotation

// support:
//   - crd: machinepools.exp.cluster.x-k8s.io
//     apiversion: v1alpha3
//     release: Since Azure 13.1.0
// documentation:
//   This annotation allows setting the min size of a node pool for autoscaling purposes.
//   See [Node pools](https://docs.giantswarm.io/advanced/node-pools)
const NodePoolMinSize = "cluster.k8s.io/cluster-api-autoscaler-node-group-min-size"

// support:
//   - crd: machinepools.exp.cluster.x-k8s.io
//     apiversion: v1alpha3
//     release: Since Azure 13.1.0
// documentation:
//   This annotation allows setting the max size of a node pool for autoscaling purposes.
//   See [Node pools](https://docs.giantswarm.io/advanced/node-pools)
const NodePoolMaxSize = "cluster.k8s.io/cluster-api-autoscaler-node-group-max-size"
