package annotation

// ConfigMajorVersion is the annotation put on App CRs and consumed by
// config-controller. This indicates what major version of the configuration
// should be used for this application. Major versions are defined as single
// number, e.g. 3, which means that the latest v3.x.x tag of configuration
// should be used to generate ConfigMap and Secret for this App CR.
const ConfigMajorVersion = "config.giantswarm.io/major-version"

// LastDeployedReleaseVersion is the version annotation put into Cluster CR to
// define which Giant Swarm release version was last successfully deployed
// during cluster creation or upgrade. Versions are defined as semver version
// without the "v" prefix, e.g. 14.1.0, which means that cluster was created
// with or upgraded to Giant Swarm release v14.1.0.
const LastDeployedReleaseVersion = "release.giantswarm.io/last-deployed-version"
