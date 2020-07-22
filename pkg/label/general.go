package label

// Certificate defines a certificate name as defined in the
// https://github.com/giantswarm/certs repository. This is used
// in certificate Secrets and CertConfig resources.
const Certificate = "giantswarm.io/certificate"

// ManagedBy informs about the operator or component managing
// a resource. The value should be a repository name, e. g.
// "kvm-operator".
const ManagedBy = "giantswarm.io/managed-by"

// Provider should indicate the (cloud) provider the
// labelled resource should be used with. Either `"aws"`,
// `"azure"`, or `"kvm"`.
const Provider = "giantswarm.io/provider"

// RandomKey should contain a randomkey name as defined in the
// https://github.com/giantswarm/randomkeys repository. This is
// used in encryption Secrets.
const RandomKey = "giantswarm.io/randomkey"

// ServiceType should be either `"system"` for core system
// services (i.e. K8s components) or `"managed"` for Giant
// Swarm managed services (i.e. networking, DNS, monitoring,
// ingress controller, etc.). Latter would be  managed by
// `chart-operator`.
const ServiceType = "giantswarm.io/service-type"
