package annotation

// LatestConfigMapVersion is the highest resource version among the configmaps
// app CRs depends on.
const LatestConfigMapVersion = "app-operator.giantswarm.io/giantswarm.io/latest-configmap-version"

// LatestSecretVersion is the highest resource version among the secret
// app CRs depends on.
const LatestSecretVersion = "app-operator.giantswarm.io/latest-secret-version"

// PauseReconcilliation annotation stops app-operator from reconciling App CR
// using it.
const PauseReconcilliation = "app-operator.giantswarm.io/pause"
