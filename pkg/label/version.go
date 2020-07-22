package label

// AWSOperatorVersion is the version label put into AWS specific CRs to define
// which aws-operator version should reconcile the given CR. Versions are
// defined as semver version without the "v" prefix, e.g. 8.7.0, which means
// that there is an aws-operator release v8.7.0.
const AWSOperatorVersion = "aws-operator.giantswarm.io/version"

// AzureOperatorVersion is the version label put into Azure specific CRs to define
// which azure-operator version should reconcile the given CR. Versions are
// defined as semver version without the "v" prefix, e.g. 8.7.0, which means
// that there is an azure-operator release v8.7.0.
const AzureOperatorVersion = "azure-operator.giantswarm.io/version"

// ClusterOperatorVersion is the version label put into provider independent CRs
// to define which cluster-operator version should reconcile the given CR.
// Versions are defined as semver version without the "v" prefix, e.g. 2.3.0,
// which means that there is a cluster-operator release v2.3.0.
const ClusterOperatorVersion = "cluster-operator.giantswarm.io/version"

// ReleaseVersion is the version label put into all CRs to define which Giant
// Swarm release the given CR is related to. Versions are defined as semver
// version without the "v" prefix, e.g. 11.4.0, which means that there is a
// Giant Swarm release v11.4.0.
const ReleaseVersion = "release.giantswarm.io/version"
