# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]
- Improve documentation


## [0.3.0] - 2020-04-16

### Changed

- Add `.spec.provider.pods` field to AWSCluster.
- Replace custom `time.Time` wrapper `DeepCopyTime` with Kubernetes built-in `metav1.Time`.
- Update `architect-orb` to `v0.8.8`.
- Generate CRDs via `kubebuilder` tools based on CRs.



## [0.2.6] - 2020-04-15

### Added

- Document G8sControlPlane CRD [#405](https://github.com/giantswarm/apiextensions/pull/405)
- Document Chart CRD [#406](https://github.com/giantswarm/apiextensions/pull/406)



## [0.2.5] - 2020-04-09

### Changed

- Fix a problem in the MachineDeployment CRD YAML file [#404](https://github.com/giantswarm/apiextensions/pull/404)



## [0.2.4] - 2020-04-08

### Changed

- Fix path of CR and CRD yaml files for Cluster and MachineDeployment [#403](https://github.com/giantswarm/apiextensions/pull/403)
- Add schema documentation for CertConfig [#401](https://github.com/giantswarm/apiextensions/pull/401)



## [0.2.3] - 2020-04-08

### Added

- Add Helm revision number to chart CR status.
- Extend Chart CR documentation.



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

- Fixed CRD OpenAPISchemas.
  - App
  - AppCatalog
  - Chart



## [0.1.2] - 2020-03-20

### Added

- Add kube-proxy configuration to Cluster type in provider.giantswarm.io/v1alpha1.



## [0.1.1] - 2020-03-12

### Fixed

- Fixed CRD OpenAPISchemas.
  - AWSCluster
  - AWSMachineDeployment
  - AWSControlPlane
  - G8SControlPlane



## [0.1.0] - 2020-03-05

### Added

- First release.



[Unreleased]: https://github.com/giantswarm/apiextensions/compare/v0.3.0...HEAD

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
