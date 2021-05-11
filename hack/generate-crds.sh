#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

# Using kubebuilder comments, create new CRDs from CR definitions in source files
pushd .. > /dev/null
./hack/tools/bin/controller-gen \
  crd \
  paths=./pkg/apis/... \
  output:dir="./config/crd" \
  crd:crdVersions="v1"
popd > /dev/null
