# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]

### Added

- Added CODEOWNERS file so that teams can more easily watch files that are relevant to them

### Changed

- Deprecated AWSConfig and StorageConfig
- Update `sigs.k8s.io/cluster-api` to v0.3.7-rc.1
- Update `sigs.k8s.io/cluster-api-provider-azure` to v0.4.5

## [0.4.14] - 2020-07-14

### Added

- Add `pkg/label` and `pkg/annotation` as strategic single source of truth.
- Add `catalog`, `reference`, and `releaseOperatorDeploy` fields to `release` CRDs, and expose a `Ready` status.

## [0.4.13] - 2020-07-13

- `AWSMachineDeployment`: Made `OnDemandBaseCapacity` and `OnDemandPercentageAboveBaseCapacity` optional attributes, removed default value for `OnDemandPercentageAboveBaseCapacity`.
- `AWSMachineDeployment`: Made `OnDemandPercentageAboveBaseCapacity` an int pointer instead of an int.

## [0.4.12] - 2020-07-10

- Change `type` of `age` column to `date` in `additionalPrinterColumn` of release CRD

## [0.4.11] - 2020-07-09

### Changed

- Update architect-orb to 0.10.0
- Add release notes URL to additionalPrinterColumns for Release CRD

## [0.4.10] - 2020-07-08

### Changed

- Allow `AzureConfig.Spec.Azure.Workers` to be null when moving towards node
  pools.

## [0.4.9] 2020-07-07

### Changed

- Allow suffixes in release names

## [0.4.8] 2020-06-22

### Changed

- Make optional fields not required for `App` and `AppCatalog` CRDs to avoid
needing to enter empty strings.
- Make optional fields `omitempty` for `App` and `AppCatalog` CRDs.



## [0.4.7] 2020-06-11

### Changed

- Update the `status` comment in `App` CRs.
- Make more fields `omitempty` in `AWSMachineDeployment` CRs.



## [0.4.6] 2020-06-01

### Added

- Add Cluster-scoped Organization CRD



## [0.4.5] 2020-06-01

### Changed

- Import latest upstream Cluster API & Azure Cluster API CRDs.



## [0.4.4] 2020-05-29

### Changed

- Make LastDeployed a nullable field.



## [0.4.3] 2020-05-25

### Changed

- Fixing AppCatalog CRD as Cluster-scoped resource.



## [0.4.2] 2020-05-25

### Changed

- Make .status.release.lastDeployed of app and chart CRs optional.



## [0.4.1] 2020-05-22

### Added

- Categories for all CRDs.

### Changed

- Make `.status.kvm.nodeIndexes` of `KVMConfig` optional.
- Update example application group CRDs to include version labels.

### Fixed

- Serialization of KVM fields `MemorySizeGB`, `StorageSizeGB`, and `Disk` broken during migration to `kubebuilder`.
- Code generation from within `$GOPATH`.
- Loading of `AWSMachineDeployment` CRD.



## [0.4.0] 2020-05-20

### Added

- Add external SNAT configuration for AWS CNI.



## [0.3.11] 2020-05-20

### Changed

- Make `altNames`,`ipSans` and `organizations` in CertConfigs optional.



## [0.3.10] 2020-05-18

### Changed

- Fix cluster scope for ETCDBackup CRD.
- Graduate CertConfig CRDs to `v1`.
- Graduate ETCDBackup CRDs to `v1`.
- Graduate Ignition CRDs to `v1`.
- Graduate Release CRDs to `v1`.
- Drop ReleaseCycle CRD.
- `Master` field in `AWSCluster` is being deprecated and made optional
- `InstanceType` in `AWSControlplane` is made optional
- Update AWSCluster docs.



## [0.3.9] 2020-05-12

### Added

- Add code generation directive (`+kubebuilder:storageversion`) to set CRD
  storage version when multiple versions for given type are present.

### Changed

- Graduated DrainerConfig CRDs to `v1`.
- Set docs URLs to our detail pages in https://docs.giantswarm.io/reference/cp-k8s-api/



## [0.3.8] 2020-05-08

- No changes.



## [0.3.7] 2020-05-08

### Added

- Add kube-proxy configuration to AWSCluster.

## Changed

- AvailabilityZones field is optional in `AzureConfig`.



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



[Unreleased]: https://github.com/giantswarm/apiextensions/compare/v0.4.14...HEAD
[0.4.14]: https://github.com/giantswarm/apiextensions/compare/v0.4.13...v0.4.14
[0.4.13]: https://github.com/giantswarm/apiextensions/compare/v0.4.12...v0.4.13
[0.4.12]: https://github.com/giantswarm/apiextensions/compare/v0.4.11...v0.4.12
[0.4.11]: https://github.com/giantswarm/apiextensions/compare/v0.4.10...v0.4.11
[0.4.10]: https://github.com/giantswarm/apiextensions/compare/v0.4.9...v0.4.10
[0.4.9]: https://github.com/giantswarm/apiextensions/compare/v0.4.8...v0.4.9
[0.4.8]: https://github.com/giantswarm/apiextensions/compare/v0.4.7...v0.4.8
[0.4.7]: https://github.com/giantswarm/apiextensions/compare/v0.4.6...v0.4.7
[0.4.6]: https://github.com/giantswarm/apiextensions/compare/v0.4.5...v0.4.6
[0.4.5]: https://github.com/giantswarm/apiextensions/compare/v0.4.4...v0.4.5
[0.4.4]: https://github.com/giantswarm/apiextensions/compare/v0.4.3...v0.4.4
[0.4.3]: https://github.com/giantswarm/apiextensions/compare/v0.4.2...v0.4.3
[0.4.2]: https://github.com/giantswarm/apiextensions/compare/v0.4.1...v0.4.2
[0.4.1]: https://github.com/giantswarm/apiextensions/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/giantswarm/apiextensions/compare/v0.3.11...v0.4.0
[0.3.11]: https://github.com/giantswarm/apiextensions/compare/v0.3.10...v0.3.11
[0.3.10]: https://github.com/giantswarm/apiextensions/compare/v0.3.9...v0.3.10
[0.3.9]: https://github.com/giantswarm/apiextensions/compare/v0.3.8...v0.3.9
[0.3.8]: https://github.com/giantswarm/apiextensions/compare/v0.3.7...v0.3.8
[0.3.7]: https://github.com/giantswarm/apiextensions/compare/v0.3.6...v0.3.7
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
