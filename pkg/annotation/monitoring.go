package annotation

// Monitoring is used to activate/deactivate monitoring of a Service resource.
// The value "true" activates it, "false" deactivates it.
const Monitoring = "giantswarm.io/monitoring"

// MonitoringPath is the path component of a metrics endpoint, e. g. "/metrics".
const MonitoringPath = "giantswarm.io/monitoring-path"

// MonitoringPort is the TCP port number to use for reaching the monitoring
// endpoint of a Service. E. g. "8080".
const MonitoringPort = "giantswarm.io/monitoring-port"
