# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Add `cilium.giantswarm.io/pod-cidr` annotation.

## [6.3.1] - 2022-07-07

### Changed

- Bump apiextensions-application to v0.5.1

## [6.3.0] - 2022-07-01

### Changed

- Bump apiextensions-application to v0.5.0

## [6.2.0] - 2022-06-15

### Changed

- Upgrade `apiextensions-application` to v0.4.0, which includes changes to Catalog CR.

## [6.1.0] - 2022-06-14

### Added

- Add `ClustersRegex` property to etcd backup type in order to allow select which clusters etcd backup will run.

### Changed

- Replace all github.com/gogo/protobuf versions with v1.3.2
- Update Silence CRD with changes from silence-operator v0.6.1

### Removed

- Remove `Organization` API.

## [6.0.0] - 2022-03-23

### Changed

- Upgrade CAPI to v1.0.4.
- Upgrade CAPZ to v1.0.1.
- Bump go module major version to v6.

### Removed

- Remove v1alpha2 legacy GiantSwarm APIs.

## [5.2.0] - 2022-03-21

### Added

- Add `alpha.aws.giantswarm.io/ebs-volume-throughput` annotation.
- Add `alpha.aws.giantswarm.io/ebs-volume-iops` annotation.

### Changed

- Updated CircleCI Ubuntu image to 20.04.

### Removed

- Don't ensure CAPI and CAPA CRDs on AWS, we will deploy them from the apps.
- Don't ensure CAPI and CAPZ CRDs on Azure, we will deploy them from the CAPI app.
- Don't ensure CAPI CRDs on OpenStack and vSphere.
- Remove templating functions and tests for old AWS CRDs.

### Fixed

- Actually remove all Cluster API CRDs from KVM installations.

## [5.1.0] - 2022-02-11

### Removed

- Remove `Cluster` CRD from KVM installations.

## [5.0.1] - 2022-02-10

### Fixed

- Bump go module major version to v5.

### Added

- Upgrade `etcdbackup` CRD.

## [5.0.0] - 2022-02-07

### Changed

- Upgrade CAPI / CAPZ CRDs to `v1beta1` on Azure.

## [4.0.1] - 2022-01-31

### Fixed

- Re-added the `preview` release state

## [4.0.0] - 2022-01-25

### Removed

- Move backup.giantswarm.io CRDs to own repo.

## [3.40.0] - 2022-01-21

### Added

- Add `ClusterNames` field to `ETCDBackup.spec` to allow backing up specific clusters within an MC.
- Add a few new fields in `ETCDBackup.status` to give better visibility on the state of the backup.

### Changed

- Replaced gopkg.in/yaml.v2 versions below v2.2.8 with v2.2.8 to mitigate CVE-2019-11254

### Removed

- Remove generated typed clients.
- Move Release API into release-operator.
- Move Config API into config-controller.
- Move Silence API into silence-operator.
- Move application.giantswarm.io API group into apiextensions-application.

### Fixed

- Fix color output during makefile execution.

## [3.39.0] - 2021-11-26

### Added

- Add AWSCNIPrefixDelegation annotation.

## [3.38.0] - 2021-11-12

### Added

- Add OpenStack Cluster API provider CRDs.
- Add support for aggregating CRDs from other repositories.
- `Notice` property for Release CRDs.

## [3.37.0] - 2021-11-12

### Added

- Add a 'preview' release state.

## [3.36.0] - 2021-11-08

### Removed

- Remove Ignition, AzureTool, and MemcachedConfig APIs and tooling and example groups.

### Added

- Add metadata about `clusterclasses.cluster.x-k8s.io` CRD for documentation.
- Add `Created At` printer column for `App` CRD.
- Add `Installed Version` printer column for `App` CRD for `-o wide` output.

### Changed

- Align AWS CAPI CRD webhooks name with `cluster-api-app`.

## [3.35.0] - 2021-10-20

### Fixed

- Use capz `v0.4.x` on Azure CRDs so that we have the experimental CRDs on their old api group.

### Changed

- CRDs ownership for Phoenix

## [3.34.0] - 2021-10-13

### Added

- Add metadata about additional vSphere CRDs for documentation.

### Fixed

- Adjust name of Cluster API certificate and service for `v1alpha4`. The annotation `cert-manager.io/inject-ca-from` value changes from `giantswarm/cluster-api-core-webhook` to `giantswarm/cluster-api-core-cert` and the service name from `cluster-api-core-webhook` to `cluster-api-core`.

### Changed

- `KVMClusterConfig`: make worker node labels optional.
- Rename provider `VMWare` to `vSphere`.
- Split Cluster API core CRDs by provider, to enable independent versioning (e.g. `v1alpha2` for AWS and `v1alpha4` for vSphere).
- Configure webhook patch for the `ClusterClass` CRD.
- Update Cluster API core CRDs to `v0.4.4` for improved defaulting and printer columns.
- In the `AppCatalogEntry` CRD, rename the printer column `APP VERSION` to `UPSTREAM VERSION` and switch the order of `VERSION` and `UPSTREAM VERSION`. This affects the output of `kubectl get appcatalogentries`.
- Update repository to use go v1.17.
- Remove reference to deprecated `AppCatalog` CRD from the `AppCatalogEntry` CRD.

## [3.33.0] - 2021-09-10

### Fixed

- Restore missing category for NetworkPool which was causing non-deterministic generation.

### Added

- Add `alpha.giantswarm.io/update-schedule-target-release` and `alpha.giantswarm.io/update-schedule-target-time` annotations.
- Add example CRs for `clusters.v1alpha3.cluster.x-k8s.io`, `machinepools.v1alpha3.exp.cluster.x-k8s.io`, `machinepools.v1alpha3.cluster.x-k8s.io` and `azuremachinepools.v1alpha3.infrastructure.cluster.x-k8s.io`.
- Add shortnames `ace` and `aces` for CRD `appcatalogentries.application.giantswarm.io`.

### Changed

- Updated URLs to CRD docs and release notes.
- Remove referencing `unique` infix from any CRDs.
- Remove App CR version label as its always defaulted.
- Update CAPV CRDs to v1alpha4 (from upstream release v0.8.1).

## [3.32.0] - 2021-08-10

### Added

- Add support for new matcher types in `silence.monitoring.giantswarm.io/v1alpha1` CRD

## [3.31.0] - 2021-08-06

### Added

- Add required field `clusterName` and label `cluster.x-k8s.io/cluster-name` for `v1alpha3` CAPI CR's.

## [3.30.0] - 2021-07-29

### Added

- Add `Chart.Description`, `Chart.Keywords` and `Chart.UpstreamChartVersion` metadata to `AppCatalogEntry` CRD.
- Add documentation of customer facing Azure annotations.

## [3.29.0] - 2021-07-27

### Added

- Add networkpool CR into Helm chart.

## [3.28.0] - 2021-07-27

### Added

- Add `v1alpha3` version for Giant Swarm AWS CRDs.

### Changed

- Add documentation for the silence.monitoring.giantswarm.io/v1alpha1 CRD.
- Conversion webhook is removed from upstream CAPZ CRDs.

### Fixed

- Add `kvm` as a valid provider in docs metadata.

## [3.27.3] - 2021-07-20

### Added

- Add deprecation info to CRD docs metadata.

### Fixed

- Set ownership of Silence CRD to Atlas.

## [3.27.2] - 2021-07-19

### Added

- Move CRD metadata for [docs.giantswarm.io](https://docs.giantswarm.io/ui-api/management-api/crd/) into this repository.

### Fixed

- Typo in certconfigs.core.giantswarm.io/v1alpha1

## [3.27.1] - 2021-07-05

### Changed

- Add documentation for the certconfigs.core.giantswarm.io/v1alpha1 CRD.

### Removed

- Drop CRD v1beta1 support.

## [3.27.0] - 2021-06-16

### Added

- Add validation for node host volumes definition in KVMConfig CRD.

### Changed

- Update `aad-pod-identity` upstream CRDs to v1.8.0.

## [3.26.0] - 2021-05-19

### Added

- Add `CatalogNamespace` spec to `App` CRD.
- Add `KVMConfigSpecKVMNodeHostVolumes` spec to `KVMConfig` CRD.

## [3.25.0] - 2021-05-18

### Changed

- Remove `alpha` prefix from `NodeTerminateUnhealthy` annotation.

## [3.24.0] - 2021-05-18

### Added

- Add `Catalog` CRD.

## [3.23.0] - 2021-05-13

### Added

- Add CAPI CRDs.
- Added the `ui.giantswarm.io/original-organization-name` annotation.
- Added status field to `Organization` CRD to hold the created namespace.

### Changed
- Updated the documentation for the `alpha.aws.giantswarm.io/aws-subnet-size` annotation to explain the current behaviour

## [3.22.0] - 2021-03-17

### Added

- Add `NamespaceConfig` spec to `Chart` CRD.

## [3.21.0] - 2021-03-16

### Added

- Add `NamespaceConfig` spec to `App` CRD.

## [3.20.0] - 2021-03-15


- Add label `ui.giantswarm.io/display`.
- Add shortnames `org` and `orgs` for CRD `organizations.security.giantswarm.io`.
- Disallow generated IDs to start with digits.

### Changed

- Terminology update to use 'workload cluster release' consistently.

## [3.19.0] - 2021-02-23

### Changed

- Register `AppCatalogEntry` CRD as a known type.

## [3.18.2] - 2021-02-18

### Changed

- Allow to add the same version again in the `common_cluster_status` to support rollbacks.

## [3.18.1] - 2021-02-15

### Added

- Add sample CRs for Azure Cluster API types.

### Changed

- Update comments in `App`, `AppCatalog` and `AppCatalogEntry` CRs.

## [3.18.0] - 2021-02-08

### Added

- Add configs.core.giantswarm.io CRD.

## [3.17.0] - 2021-02-03

### Changed

- Update CAPZ fork to v0.4.12-alpha1

## [3.16.1] - 2021-01-28

### Added

- Update CAPZ fork to v0.4.11

## [3.16.0] - 2021-01-28

### Added

- Add `Restrictions.CompatibleProviders`, `Chart.ApiVersion` metadata to `AppCatalogEntry` CRD.
- Enable rendering documentation for annotations in public documentation.
- Update CAPI fork to v0.3.13

## [3.15.1] - 2021-01-26

### Fixed

- Set a `LastHeartbeatTime` when creating new `Condition`s for `DrainerConfig`s.

## [3.15.0] - 2021-01-21

### Added

- Add `SkipCRDs` to `ChartSpecConfig` and `AppSpecConfig`.
- Add AWS CNI annotation to configure `WARM_IP_TARGET` and `MINIMUM_IP_TARGET`.

### Fixed

- Fix name for `AppCatalogEntry` example.

## [3.14.1] - 2021-01-07

### Changed

- Changed terminology from `tenant cluster` to `workload cluster`.
- Changed terminology from `control plane` to `management cluster` (where appropriate).

## [3.14.0] - 2021-01-05

### Added

- Add `app-operator.giantswarm.io/paused` annotation.

### Changed

- Update description and example annotation for the Spark CRD.

## [3.13.0] - 2020-12-07

### Changed

- Make `credentialSecret` attribute in `AWSCluster` optional. In case this attribute is not set it will be defaulted
  by `aws-admission-controller` to the credential-secret for the organization the cluster is created in.

## [3.12.0] - 2020-12-03

### Added

- Add `app-operator.giantswarm.io/latest-configmap-version` annotation.
- Add `app-operator.giantswarm.io/latest-secret-version` annotation.

## [3.11.0] - 2020-12-02

### Changed

- Change (unused yet) `config.giantswarm.io/major-version` annotation to `config.giantswarm.io/version`.

## [3.10.0] - 2020-11-30

### Changed

- Make `availabilityZones` attribute in `AWSMachineDeployment` optional.

### Added

- Add `config-controller.giantswarm.io/version` label.
- Add `config.giantswarm.io/major-version` annotation.

## [3.9.0] - 2020-11-24

### Added

- Add `catalog` field to `apps` in `release`.
- Add printer columns for Release `Ready` and `InUse` fields.
- Add printer columns for App, Chart `Version`, `Last Deployed` and `Status`.

### Changed

- Make Release Status fields `Ready` and `InUse` optional.

## [3.8.0] - 2020-11-13

### Added

- Add `Restrictions` metadata to `AppCatalogEntry` CRD.
- Add legacy `app` label.

## [3.7.0] - 2020-11-04

### Added

- Add `Silence` CRD.

## [3.6.0] - 2020-11-03

### Added

- Add catalog and kubernetes labels and notes annotation.
- Add `release.giantswarm.io/last-deployed-version` `Cluster` CR annotation

## [3.5.0] - 2020-11-03

### Added

- Add 'AWSMetadataV2' annotation to configure the metadata endpoint.
- Add 'AWSSubnetSize' annotation to configure the subnet size of Control Plane and Machinedeployments.

## [3.4.1] - 2020-10-29

### Added

- Add 'AWSUpdateMaxBatchSize' annotation to configure max batch size for AWS ASG update.
- Add 'AWSUpdatePauseTime' annotation to configure pause time between rolling a single batch in ASG.

## [3.4.0] - 2020-10-26

### Added

- Add annotation to enable feature to terminate unhealthy nodes on a cluster.
- `Cluster` condition `ProviderInfrastructureReady`: `True` when `AzureCluster` is ready
- `Cluster` condition `NodePoolsReady`: `True` when all node pools are ready
- `Cluster` `Upgrading` condition `UpgradeCompleted` reason: used when `Upgrading` is set to `False` because the upgrade has been completed
- `Cluster` `Upgrading` condition `UpgradeNotStarted` reason: used when `Upgrading` is set to `False` because the upgrade has not been started
- `Cluster` `Creating` condition `CreationCompleted` reason: used when `Creating` is set to `False` because the creation has been completed
- `Cluster` `Creating` condition `ExistingCluster` reason: used when `Creating` is set to `False` because an older cluster (created without Conditions support) is upgraded to newer release that has conditions.

## [3.3.0] - 2020-10-23

### Added

- Add display columns to `AppCatalogEntry` CRD.

## [3.2.0] - 2020-10-15

### Added

- Add `organization` CR description.

### Changed

- Hide LastDeployed, revision if they are not presented.

## [3.1.0] - 2020-10-09

### Changed

- Update CAPZ dependencies.

### Fixed

- Bump go module major version to v3.

## [3.0.0] - 2020-10-08

### Changed

- Consumers of this library need to explicitly replace CAPI/CAPZ dependencies with GiantSwarm forks on their `go.mod` files.
- Update microerror.

## [2.6.2] - 2020-10-09

### Fixed

- Revert changes in release `v2.6.1`.

## [2.6.1] - 2020-10-07

### Changed

- Update k8s related dependencies.

## [2.6.0] - 2020-10-05

- Add `AppCatalogEntry` CRD.

## [2.5.3] - 2020-10-02

### Added

- Add NetworkPool option for ClusterCRsConfig

## [2.5.2] - 2020-10-01

### Added

- Functions for generating NetworkPools CRs

## [2.5.1] - 2020-09-23

### Removed

- Removed CRDs related to Azure managed AKS clusters.

## [2.5.0] - 2020-09-22

### Added

- Add `NetworkPool` CRD.

### Changed

- Marked nullable fields in ETCDBackup types.

## [2.4.0] - 2020-09-17

### Added

- Added constants for oidc annotation names.

## [2.3.0] - 2020-09-16

### Changed

- Replace `sigs.k8s.io/cluster-api` with our fork.
- Replace `sigs.k8s.io/cluster-api-provider-azure` with our fork.

## [2.2.0] - 2020-09-15

### Added

- Add `InUse` to the `ReleaseStatus`.
- Add `KVMOperatorVersion` label.
- Added constants for node pools autoscaling annotation names.

## [2.1.0] - 2020-08-17

### Added

- Added labels for configuring scraping of services by Prometheus
  (`Monitoring`, `MonitoringPath`, `MonitoringPort`) to `pkg/label`.
- Add managed-by label.
- Add version labels for app-operator and chart-operator.

## [2.0.1] - 2020-08-13

### Changed

- Make optional fields not required for `Chart` CRD to avoid
needing to enter empty strings.
- Make optional fields `omitempty` for `Chart` CRD.

## [2.0.0] - 2020-08-10

### Changed

- Update Kubernetes dependencies to v1.18.5
- Update `sigs.k8s.io/cluster-api` to v0.3.7
- Update `sigs.k8s.io/cluster-api-provider-azure` to v0.4.6

## [0.4.20] - 2020-07-31

### Changed

- Graduate StorageConfig CRD to `v1`.

## [0.4.19] - 2020-07-29

### Added

- Add `ClusterDescription` to `pkg/annotation`
- Add `Spark` CRD for Azure Cluster API migration.

## [0.4.18] - 2020-07-27

- Add `EndOfLifeDate` to `Release` CRD.

## [0.4.17] - 2020-07-23

### Added

- Add `AzureOperatorVersion` to `pkg/label`.
- Add `MachinePool` to `pkg/label`
- Add `MachinePoolName` to `pkg/annotation`
- Add `ReleaseNotesURL` to `pkg/annotation`.
- Add descriptions to `App`, `AppCatalog` and `Chart` CRDs.
- Add deprecation notice to `ChartConfig` CRD.

### Changed

- Graduate AppCatalog CRDs to `v1`.
- Graduate App CRDs to `v1`.
- Graduate Chart CRDs to `v1`.

## [0.4.16] - 2020-07-20

### Added

- Add CR templating for external use.



## [0.4.15] - 2020-07-15

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



[Unreleased]: https://github.com/giantswarm/apiextensions/compare/v6.3.1...HEAD
[6.3.1]: https://github.com/giantswarm/apiextensions/compare/v6.3.0...v6.3.1
[6.3.0]: https://github.com/giantswarm/apiextensions/compare/v6.2.0...v6.3.0
[6.2.0]: https://github.com/giantswarm/apiextensions/compare/v6.1.0...v6.2.0
[6.1.0]: https://github.com/giantswarm/apiextensions/compare/v6.0.0...v6.1.0
[6.0.0]: https://github.com/giantswarm/apiextensions/compare/v5.2.0...v6.0.0
[5.2.0]: https://github.com/giantswarm/apiextensions/compare/v5.1.0...v5.2.0
[5.1.0]: https://github.com/giantswarm/giantswarm/compare/v5.0.1...v5.1.0
[5.0.1]: https://github.com/giantswarm/giantswarm/compare/v5.0.0...v5.0.1
[5.0.0]: https://github.com/giantswarm/giantswarm/compare/v4.0.1...v5.0.0
[4.0.1]: https://github.com/giantswarm/giantswarm/compare/v4.0.0...v4.0.1
[4.0.0]: https://github.com/giantswarm/apiextensions/compare/v3.40.0...v4.0.0
[3.40.0]: https://github.com/giantswarm/apiextensions/compare/v3.39.0...v3.40.0
[3.39.0]: https://github.com/giantswarm/apiextensions/compare/v3.38.0...v3.39.0
[3.38.0]: https://github.com/giantswarm/apiextensions/compare/v3.37.0...v3.38.0
[3.37.0]: https://github.com/giantswarm/apiextensions/compare/v3.36.0...v3.37.0
[3.36.0]: https://github.com/giantswarm/apiextensions/compare/v3.35.0...v3.36.0
[3.35.0]: https://github.com/giantswarm/apiextensions/compare/v3.34.0...v3.35.0
[3.34.0]: https://github.com/giantswarm/apiextensions/compare/v3.33.0...v3.34.0
[3.33.0]: https://github.com/giantswarm/apiextensions/compare/v3.32.0...v3.33.0
[3.32.0]: https://github.com/giantswarm/apiextensions/compare/v3.31.0...v3.32.0
[3.31.0]: https://github.com/giantswarm/apiextensions/compare/v3.30.0...v3.31.0
[3.30.0]: https://github.com/giantswarm/apiextensions/compare/v3.29.0...v3.30.0
[3.29.0]: https://github.com/giantswarm/apiextensions/compare/v3.28.0...v3.29.0
[3.28.0]: https://github.com/giantswarm/apiextensions/compare/v3.27.3...v3.28.0
[3.27.3]: https://github.com/giantswarm/apiextensions/compare/v3.27.2...v3.27.3
[3.27.2]: https://github.com/giantswarm/apiextensions/compare/v3.27.1...v3.27.2
[3.27.1]: https://github.com/giantswarm/apiextensions/compare/v3.27.0...v3.27.1
[3.27.0]: https://github.com/giantswarm/apiextensions/compare/v3.26.0...v3.27.0
[3.26.0]: https://github.com/giantswarm/apiextensions/compare/v3.25.0...v3.26.0
[3.25.0]: https://github.com/giantswarm/apiextensions/compare/v3.24.0...v3.25.0
[3.24.0]: https://github.com/giantswarm/apiextensions/compare/v3.23.0...v3.24.0
[3.23.0]: https://github.com/giantswarm/apiextensions/compare/v3.22.0...v3.23.0
[3.22.0]: https://github.com/giantswarm/apiextensions/compare/v3.21.0...v3.22.0
[3.21.0]: https://github.com/giantswarm/apiextensions/compare/v3.20.0...v3.21.0
[3.20.0]: https://github.com/giantswarm/apiextensions/compare/v3.19.0...v3.20.0
[3.19.0]: https://github.com/giantswarm/apiextensions/compare/v3.18.2...v3.19.0
[3.18.2]: https://github.com/giantswarm/apiextensions/compare/v3.18.1...v3.18.2
[3.18.1]: https://github.com/giantswarm/apiextensions/compare/v3.18.0...v3.18.1
[3.18.0]: https://github.com/giantswarm/apiextensions/compare/v3.17.0...v3.18.0
[3.17.0]: https://github.com/giantswarm/apiextensions/compare/v3.16.1...v3.17.0
[3.16.1]: https://github.com/giantswarm/apiextensions/compare/v3.16.0...v3.16.1
[3.16.0]: https://github.com/giantswarm/apiextensions/compare/v3.15.1...v3.16.0
[3.15.1]: https://github.com/giantswarm/apiextensions/compare/v3.15.0...v3.15.1
[3.15.0]: https://github.com/giantswarm/apiextensions/compare/v3.14.1...v3.15.0
[3.14.1]: https://github.com/giantswarm/apiextensions/compare/v3.14.0...v3.14.1
[3.14.0]: https://github.com/giantswarm/apiextensions/compare/v3.13.0...v3.14.0
[3.13.0]: https://github.com/giantswarm/apiextensions/compare/v3.12.0...v3.13.0
[3.12.0]: https://github.com/giantswarm/apiextensions/compare/v3.11.0...v3.12.0
[3.11.0]: https://github.com/giantswarm/apiextensions/compare/v3.10.0...v3.11.0
[3.10.0]: https://github.com/giantswarm/apiextensions/compare/v3.9.0...v3.10.0
[3.9.0]: https://github.com/giantswarm/apiextensions/compare/v3.8.0...v3.9.0
[3.8.0]: https://github.com/giantswarm/apiextensions/compare/v3.7.0...v3.8.0
[3.7.0]: https://github.com/giantswarm/apiextensions/compare/v3.6.0...v3.7.0
[3.6.0]: https://github.com/giantswarm/apiextensions/compare/v3.5.0...v3.6.0
[3.5.0]: https://github.com/giantswarm/apiextensions/compare/v3.4.1...v3.5.0
[3.4.1]: https://github.com/giantswarm/apiextensions/compare/v3.4.0...v3.4.1
[3.4.0]: https://github.com/giantswarm/apiextensions/compare/v3.3.0...v3.4.0
[3.3.0]: https://github.com/giantswarm/apiextensions/compare/v3.2.0...v3.3.0
[3.2.0]: https://github.com/giantswarm/apiextensions/compare/v3.1.0...v3.2.0
[3.1.0]: https://github.com/giantswarm/apiextensions/compare/v3.0.0...v3.1.0
[3.0.0]: https://github.com/giantswarm/apiextensions/compare/v2.6.1...v3.0.0
[2.6.1]: https://github.com/giantswarm/apiextensions/compare/v2.6.0...v2.6.1
[2.6.0]: https://github.com/giantswarm/apiextensions/compare/v2.5.3...v2.6.0
[2.5.3]: https://github.com/giantswarm/apiextensions/compare/v2.5.2...v2.5.3
[2.5.2]: https://github.com/giantswarm/apiextensions/compare/v2.5.1...v2.5.2
[2.5.1]: https://github.com/giantswarm/apiextensions/compare/v2.5.0...v2.5.1
[2.5.0]: https://github.com/giantswarm/apiextensions/compare/v2.4.0...v2.5.0
[2.4.0]: https://github.com/giantswarm/apiextensions/compare/v2.3.0...v2.4.0
[2.3.0]: https://github.com/giantswarm/apiextensions/compare/v2.2.0...v2.3.0
[2.2.0]: https://github.com/giantswarm/apiextensions/compare/v2.1.0...v2.2.0
[2.1.0]: https://github.com/giantswarm/apiextensions/compare/v2.0.1...v2.1.0
[2.0.1]: https://github.com/giantswarm/apiextensions/compare/v2.0.0...v2.0.1
[2.0.0]: https://github.com/giantswarm/apiextensions/compare/v0.4.20...v2.0.0
[0.4.20]: https://github.com/giantswarm/apiextensions/compare/v0.4.19...v0.4.20
[0.4.19]: https://github.com/giantswarm/apiextensions/compare/v0.4.18...v0.4.19
[0.4.18]: https://github.com/giantswarm/apiextensions/compare/v0.4.17...v0.4.18
[0.4.17]: https://github.com/giantswarm/apiextensions/compare/v0.4.16...v0.4.17
[0.4.16]: https://github.com/giantswarm/apiextensions/compare/v0.4.15...v0.4.16
[0.4.15]: https://github.com/giantswarm/apiextensions/compare/v0.4.14...v0.4.15
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
