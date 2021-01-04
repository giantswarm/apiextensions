package annotation

// AppOperatorPaused annotation when present prevents app-operator from
// reconciling the resource.
const AppOperatorPaused = "app-operator.giantswarm.io/paused"

// LatestConfigMapVersion is the highest resource version among the configmaps
// app CRs depends on.
const LatestConfigMapVersion = "app-operator.giantswarm.io/giantswarm.io/latest-configmap-version"

// LatestSecretVersion is the highest resource version among the secret
// app CRs depends on.
const LatestSecretVersion = "app-operator.giantswarm.io/latest-secret-version"
