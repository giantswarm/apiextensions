package annotation

// ConfigVersion is the annotation put on App CRs and consumed by
// config-controller. This indicates what major version of the configuration
// should be used for this application. Versions are configured in a
// <major>.x.x format (e.g. 3.x.x), which means the latest v3.<minor>.<patch>
// should be used to generate ConfigMap and Secret for this App CR.
// When given version does not match the <major>.x.x format, config-controller
// assumes given version is a branch reference (e.g. "master") and the matching
// branch will be used to generate configuration instead.
const ConfigVersion = "config.giantswarm.io/version"

// LastDeployedReleaseVersion is the version annotation put into Cluster CR to
// define which Giant Swarm release version was last successfully deployed
// during cluster creation or upgrade. Versions are defined as semver version
// without the "v" prefix, e.g. 14.1.0, which means that cluster was created
// with or upgraded to Giant Swarm release v14.1.0.
const LastDeployedReleaseVersion = "release.giantswarm.io/last-deployed-version"
