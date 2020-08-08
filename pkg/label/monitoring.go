package label

// Monitoring is a label that can be used on `Service` resources to mark them
// for scraping by Giant Swarm Prometheus instances that monitor the cluster
// they're found in. The service will be scraped only if the value is `true`.
//
// Spec: https://intranet.giantswarm.io/docs/architecture-specs-adrs/specs/configuration-of-targets-in-tc-prometheus/
const Monitoring = "giantswarm.io/monitoring"

// MonitoringPort is an annotation that tells Prometheus which port a service
// exposes metrics on. It can be used on `Service` resources and will be picked
// up by `simple-service` `ServiceMonitor` to configure Prometheus.
//
// The default is to use the main port exposed in the `Service` resource if
// this annotation isn't used.
//
// Spec: https://intranet.giantswarm.io/docs/architecture-specs-adrs/specs/configuration-of-targets-in-tc-prometheus/
const MonitoringPort = "giantswarm.io/monitoring-port"

// MonitoringPath is an annotation that tells Prometheus the URL path a service
// exposes metrics on. It can be used on `Service` resources and will be picked
// up by `simple-service` `ServiceMonitor` to configure Prometheus.
//
// The default is to use the path `/metrics`.
//
// Spec: https://intranet.giantswarm.io/docs/architecture-specs-adrs/specs/configuration-of-targets-in-tc-prometheus/
const MonitoringPath = "giantswarm.io/monitoring-path"
