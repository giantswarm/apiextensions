package annotation

// AnnotationNodePoolMinSize is the cluster annotation used for storing
// the minimum size of a node pool.
const NodePoolMinSize = "cluster.k8s.io/cluster-api-autoscaler-node-group-min-size"

// AnnotationNodePoolMaxSize is the cluster annotation used for storing
// the maximum size of a node pool.
const NodePoolMaxSize = "cluster.k8s.io/cluster-api-autoscaler-node-group-max-size"
