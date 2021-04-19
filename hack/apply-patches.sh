#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

for version in v1 v1beta1; do
  for crd in "../config/crd/patches/$version"/*; do
    ./tools/bin/kustomize --load_restrictor LoadRestrictionsNone build \
      "$crd" \
      -o "../config/crd/$version/$(basename "$crd").yaml"
  done
done
