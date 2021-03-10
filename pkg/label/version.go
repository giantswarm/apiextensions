package label

// AppOperatorVersion is the version label put into app CRs to define
// which app-operator version should reconcile the given CR. Versions are
// defined as semver version without the "v" prefix, e.g. 2.1.0, which means
// that there is an app-operator release v2.1.0.
const AppOperatorVersion = "app-operator.giantswarm.io/version"

// ChartOperatorVersion is the version label put into app CRs to define
// which app-operator version should reconcile the given CR. Versions are
// defined as semver version without the "v" prefix, e.g. 1.0.0, which means
// that there is a chart-operator release v1.0.0.
const ChartOperatorVersion = "chart-operator.giantswarm.io/version"

// ConfigControllerVersion is the version label put into app CRs to define
// which config-controller version should reconcile the given CR. Versions are
// defined as semver version without the "v" prefix, e.g. 1.0.0, which means
// that there is a config-controller release v1.0.0.
const ConfigControllerVersion = "config-controller.giantswarm.io/version"

// AWSOperatorVersion is the version label put into AWS specific CRs to define
// which aws-operator version should reconcile the given CR. Versions are
// defined as semver version without the "v" prefix, e.g. 8.7.0, which means
// that there is an aws-operator release v8.7.0.
const AWSOperatorVersion = "aws-operator.giantswarm.io/version"

// AzureOperatorVersion is the version label put into Azure specific CRs to define
// which azure-operator version should reconcile the given CR. Versions are
// defined as semver version without the "v" prefix, e.g. 4.1.0, which means
// that there is an azure-operator release v4.1.0.
const AzureOperatorVersion = "azure-operator.giantswarm.io/version"

// CAPIVersion is the version label put into provider independent CRs
// to define which cluster-api-core version should reconcile the given CR.
const CAPIVersion = "cluster-api.giantswarm.io/version"

// CABPKVersion is the version label put into bootstrap CRs
// to define which cluster-api-bootstrap-provider-kubeadm version should reconcile the given CR.
const CABPKVersion = "cluster-api-bootstrap-kubeadm.giantswarm.io/version"

// CACPKVersion is the version label put into control plane CRs
// to define which cluster-api-control-plane-kubeadm version should reconcile the given CR.
const CACPKVersion = "cluster-api-control-plane-kubeadm.giantswarm.io/version"

// CAPAVersion is the version label put into AWS CRs
// to define which CAPA version should reconcile the given CR.
const CAPAVersion = "cluster-api-provider-aws.giantswarm.io/version"

// CAPZVersion is the version label put into Azure CRs
// to define which CAPZ version should reconcile the given CR.
const CAPZVersion = "cluster-api-provider-azure.giantswarm.io/version"

// ClusterOperatorVersion is the version label put into VSphere CRs
// to define which CAPV version should reconcile the given CR.
const CAPVVersion = "cluster-api-provider-vsphere.giantswarm.io/version"

// ClusterOperatorVersion is the version label put into provider independent CRs
// to define which cluster-operator version should reconcile the given CR.
// Versions are defined as semver version without the "v" prefix, e.g. 2.3.0,
// which means that there is a cluster-operator release v2.3.0.
const ClusterOperatorVersion = "cluster-operator.giantswarm.io/version"

// KVMOperatorVersion is the version label put into KVM specific CRs to define
// which kvm-operator version should reconcile the given CR. Versions are
// defined as semver version without the "v" prefix, e.g. 8.7.0, which means
// that there is a kvm-operator release v8.7.0.
const KVMOperatorVersion = "kvm-operator.giantswarm.io/version"

// ReleaseVersion is the version label put into all CRs to define which Giant
// Swarm release the given CR is related to. Versions are defined as semver
// version without the "v" prefix, e.g. 11.4.0, which means that there is a
// Giant Swarm release v11.4.0.
const ReleaseVersion = "release.giantswarm.io/version"
