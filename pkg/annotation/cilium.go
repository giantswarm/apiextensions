package annotation

//Keep the YAML documentation format as it is used to render in the CRD public documentation. You can use Markdown in the documentation field.

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha3
//     release: Since 18.0.0
//
// documentation:
//
//	This annotation allows specifying a CIDR to be used by cilium during cluster upgrades from v17 to v18.
const CiliumPodCidr = "cilium.giantswarm.io/pod-cidr"
