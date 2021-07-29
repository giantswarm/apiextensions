package annotation

// support:
//   - crd: azureclusters.infrastructure.cluster.x-k8s.io
//     apiversion: v1alpha3
//     release: Since 15.1.0
// documentation:
//   This annotation allows reusing an existing public IP address for egress traffic of worker nodes.
//   See [Setting an egress IP address on Azure](https://docs.giantswarm.io/advanced/egress-ip-address-azure/)
const AzureWorkersEgressExternalPublicIP = "alpha.cni.aws.giantswarm.io/minimum-ip-target"
