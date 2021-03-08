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

  # Add .metadata.name validation to Release CRD using kustomize since
  # kubebuilder comments can't modify metav1.ObjectMeta
  for crd in "../config/crd/patches/$version"/*; do
    ./tools/bin/kustomize --load_restrictor LoadRestrictionsNone build \
      "$crd" \
      -o "../config/crd/$version/$(basename "$crd").yaml"
  done
done
