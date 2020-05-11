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
  # The line `paths=sigs.k8s.io/cluster-api/api/v1alpha2`
  # indicates that the CRD should contain only `v1alpha2`.
  # At the time of writing, `paths=sigs.k8s.io/cluster-api/api` would
  # result in both `v1alpha2` and `v1alpha3` being included.
  ./tools/bin/controller-gen \
    crd \
    paths=sigs.k8s.io/cluster-api/api/v1alpha2 \
    output:dir="../config/crd/$version" \
    crd:crdVersions="$version"

  # We only want Cluster and MachineDeployment for now, so delete the other two CAPI CRDs.
  rm ../config/crd/$version/cluster.x-k8s.io_machines.yaml
  rm ../config/crd/$version/cluster.x-k8s.io_machinesets.yaml

  # Add .metadata.name validation to Release CRD using kustomize since
  # kubebuilder comments can't modify metav1.ObjectMeta
  for crd in "../config/crd/patches/$version"/*; do
    ./tools/bin/kustomize --load_restrictor LoadRestrictionsNone build \
      "$crd" \
      -o "../config/crd/$version/$(basename "$crd").yaml"
  done

  mkdir -p "../config/crd/patches/all"
  for crd in "../config/crd/$version"/*; do
    kind=$(basename "$crd" | sed 's/.*_//' | sed 's/\..*//')
    group=$(basename "$crd" | sed 's/_.*//')
    name=$kind.$group
    cat > "../config/crd/patches/all/patch.yaml" <<EOF
- op: remove
  path: /metadata/creationTimestamp
- op: remove
  path: /metadata/annotations
- op: add
  path: /metadata/annotations
  value:
    giantswarm.io/docs: https://docs.giantswarm.io/reference/cp-k8s-api/$name/
EOF
    cat > "../config/crd/patches/all/kustomization.yaml" <<EOF
resources:
  - ../../$version/$(basename "$crd")
patchesJson6902:
  - target:
      group: apiextensions.k8s.io
      version: $version
      kind: CustomResourceDefinition
      name: $name
    path: patch.yaml
EOF
    ./tools/bin/kustomize --load_restrictor LoadRestrictionsNone build \
      "../config/crd/patches/all" \
      -o "$crd"
  done
  rm -rf "../config/crd/patches/all"
done
