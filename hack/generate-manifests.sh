#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

for version in v1 v1beta1; do
  # Using kubebuilder comments, create new CRDs from CR definitions in source files
  pushd .. > /dev/null
  ./hack/tools/bin/controller-gen \
    crd \
    paths=./pkg/apis/... \
    output:dir="./config/crd/$version" \
    crd:crdVersions="$version"
  popd > /dev/null

  # Building Cluster API CRDs.
  # At the time of writing, `paths=sigs.k8s.io/cluster-api/api/...` would
  # result in both `v1alpha2` and `v1alpha3` being included.
  ./tools/bin/controller-gen \
    crd \
    paths="sigs.k8s.io/cluster-api/api/..." \
    output:dir="../config/crd/$version" \
    crd:crdVersions="$version"

  # Also build experimental Cluster API CRDs. Most importantly MachinePool CRD.
  ./tools/bin/controller-gen \
    crd \
    paths="sigs.k8s.io/cluster-api/exp/api/..." \
    output:dir="../config/crd/$version" \
    crd:crdVersions="$version"

  # We also want Provider specific CRDs for Azure.
  ./tools/bin/controller-gen \
    crd \
    paths=sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3 \
    output:dir="../config/crd/$version" \
    crd:crdVersions="$version"

  # With MachinePool related types.
  ./tools/bin/controller-gen \
    crd \
    paths=sigs.k8s.io/cluster-api-provider-azure/exp/api/v1alpha3 \
    output:dir="../config/crd/$version" \
    crd:crdVersions="$version"

  # Delete unused upstream CRDs.
  rm ../config/crd/$version/infrastructure.cluster.x-k8s.io_azuremachinetemplates.yaml
  rm ../config/crd/$version/exp.infrastructure.cluster.x-k8s.io_azuremanagedclusters.yaml
  rm ../config/crd/$version/exp.infrastructure.cluster.x-k8s.io_azuremanagedcontrolplanes.yaml
  rm ../config/crd/$version/exp.infrastructure.cluster.x-k8s.io_azuremanagedmachinepools.yaml


  # Add .metadata.name validation to Release CRD using kustomize since
  # kubebuilder comments can't modify metav1.ObjectMeta
  for crd in "../config/crd/patches/$version"/*; do
    ./tools/bin/kustomize --load_restrictor LoadRestrictionsNone build \
      "$crd" \
      -o "../config/crd/$version/$(basename "$crd").yaml"
  done
done
