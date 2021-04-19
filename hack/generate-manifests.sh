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
done
