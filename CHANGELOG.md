# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

### Changed
- ClusterKubernetesIngressController added LoadBalancerType.
- ClusterKubernetesIngressController added ExternalTrafficPolicy.

## [0.3.6] 2020-05-07

### Changed

- Load StorageConfig from VFS as expected.
- AzureClusterConfig allow empty labels for guest cluster worker nodes.

## [0.3.5] - 2020-05-06

### Added

- Add Azure Tools CRDs.

## [0.3.4] - 2020-04-30

### Changed

- All CRDs are now available as both `v1.CustomResourcDefinition` and `v1beta1.CustomResourceDefinition` through 
  `crd.LoadV1` and `crd.LoadV1Beta1`. Type-specific `New*CRD()` functions are unchanged.
- Graduated Azure CRDs to `v1`.

## [0.3.3] - 2020-04-28

### Added

- Generate Cluster API CRDs from upstream module.


## [0.3.2] - 2020-04-27

### Changed

- Relax `AzureConfig` CRD validation.


## [0.3.1] - 2020-04-22

### Added

- Modified docs for the G8sControlPlane CRD
- Add property descriptions to AWSMachineDeployment in infrastructure.giantswarm.io/v1alpha2.
- Add property descriptions to AWSControlPlane in infrastructure.giantswarm.io/v1alpha2.

### Changed

- Change `release` CR back to be cluster scoped.
- Make more CR status fields `omitempty`.
- Make CR status fields optional.


### Fixed

- Fix mistake in the main description of `G8sControlPlane` in `infrastructure.giantswarm.io`.


## [0.3.0] - 2020-04-16

### Added

- Add `.spec.provider.pods` field to AWSCluster in core.giantswarm.io/v1alpha1.

### Changed

- Replace custom `time.Time` wrapper `DeepCopyTime` with Kubernetes built-in `metav1.Time`.
- Generate CRDs via `kubebuilder` tools based on CRs.


## [0.2.6] - 2020-04-15

### Added

- Document G8sControlPlane in infrastructure.giantswarm.io [#405](https://github.com/giantswarm/apiextensions/pull/405)
- Document Chart in core.giantswarm.io/v1alpha1 [#406](https://github.com/giantswarm/apiextensions/pull/406)


## [0.2.5] - 2020-04-09

### Changed

- Fix docs for MachineDeployment machinedeployments.cluster.x-k8s.io [#404](https://github.com/giantswarm/apiextensions/pull/404)


## [0.2.4] - 2020-04-08

### Added

- Add schema documentation for CertConfig in core.giantswarm.io/v1alpha1 [#401](https://github.com/giantswarm/apiextensions/pull/401)

### Changed

- Fix path of CR and CRD yaml files for Cluster and MachineDeployment in infrastructure.giantswarm.io/v1alpha2 [#403](https://github.com/giantswarm/apiextensions/pull/403)


## [0.2.3] - 2020-04-08

### Added

- Add Helm revision number to chart CR status.
- Extend documentation Chart CR documentation.


## [0.2.2] - 2020-04-07

### Added

- Extend App and AppCatalog CR documentation.


## [0.2.1] - 2020-04-06

### Added

- Add CRD and CR documentation.
- Add Spot Instances configuration.
- Add VPC ID to be exposed with the AWSCluster CR.

### Changed

- Make OIDC `omitempty`.


## [0.2.0] - 2020-03-20

### Changed

- Switch from dep to Go modules.

### Fixed

- Fix CRD OpenAPISchemas for:
  - App
  - AppCatalog
  - Chart


## [0.1.2] - 2020-03-20

### Added

- Add kube-proxy configuration to Cluster type in provider.giantswarm.io/v1alpha1.


## [0.1.1] - 2020-03-12

### Fixed

- Fix CRD OpenAPISchemas for:
  - AWSCluster
  - AWSMachineDeployment
  - AWSControlPlane
  - G8SControlPlane



## [0.1.0] - 2020-03-05

### Added

- First release.


[Unreleased]: https://github.com/giantswarm/apiextensions/compare/v0.3.6...HEAD

[0.3.6]: https://github.com/giantswarm/apiextensions/compare/v0.3.5...v0.3.6
[0.3.5]: https://github.com/giantswarm/apiextensions/compare/v0.3.4...v0.3.5
[0.3.4]: https://github.com/giantswarm/apiextensions/compare/v0.3.3...v0.3.4
[0.3.3]: https://github.com/giantswarm/apiextensions/compare/v0.3.2...v0.3.3
[0.3.2]: https://github.com/giantswarm/apiextensions/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/giantswarm/apiextensions/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/giantswarm/apiextensions/compare/v0.2.6...v0.3.0
[0.2.6]: https://github.com/giantswarm/apiextensions/compare/v0.2.5...v0.2.6
[0.2.5]: https://github.com/giantswarm/apiextensions/compare/v0.2.4...v0.2.5
[0.2.4]: https://github.com/giantswarm/apiextensions/compare/v0.2.3...v0.2.4
[0.2.3]: https://github.com/giantswarm/apiextensions/compare/v0.2.2...v0.2.3
[0.2.2]: https://github.com/giantswarm/apiextensions/compare/v0.2.1...v0.2.2
[0.2.1]: https://github.com/giantswarm/apiextensions/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/giantswarm/apiextensions/compare/v0.1.2...v0.2.0
[0.1.2]: https://github.com/giantswarm/apiextensions/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/giantswarm/apiextensions/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/giantswarm/apiextensions/releases/tag/v0.1.0
