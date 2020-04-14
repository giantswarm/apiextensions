#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOPATH=$(go env GOPATH)
export GOROOT=$(go env GOROOT)

tools="sigs.k8s.io/controller-tools/cmd/controller-gen
k8s.io/code-generator/cmd/deepcopy-gen
k8s.io/code-generator/cmd/client-gen
golang.org/x/tools/cmd/goimports
github.com/markbates/pkger/cmd/pkger
sigs.k8s.io/kustomize/kustomize/v3"

mkdir -p "$dir/bin"
cd "$dir"
for tool in $tools; do
  base=$(echo "$tool" | sed "s/\/cmd\/.*//")
  version=$(go list -m -f '{{.Version}}' "$base")
  echo "Rebuilding $tool@$version"
  go build -o "$dir/bin" "$tool"
done
cd ..

module="github.com/giantswarm/apiextensions"
input_dirs=$(find ./pkg/apis -maxdepth 2 -mindepth 2 | tr '\r\n' ',')
input_dirs=${input_dirs%?}
groups=${input_dirs//.\/pkg\/apis\//}
header="${dir}/boilerplate.go.txt"

echo "Generating deepcopy funcs"
"$dir/bin/deepcopy-gen" \
  --input-dirs "$input_dirs" \
  --output-file-base zz_generated.deepcopy \
  --output-base "$dir/.." \
  --go-header-file "${dir}/boilerplate.go.txt"

echo "Generating clientset"
"$dir/bin/client-gen" \
  --clientset-name versioned \
  --input "$groups" \
  --input-base "$module/pkg/apis" \
  --output-package "$module/pkg/clientset" \
  --output-base "$dir" \
  --go-header-file "$header"

echo "Moving generated files to expected location"
cp -R "$dir/$module/pkg"/* "$dir/../pkg"
rm -rf "$dir/github.com"

# code-generator doesn't group local imports separately from third-party
# imports, so run goimports after generating new code.
cd "$dir/.." || exit
echo "Fixing imports in-place with goimports"
"$dir/bin/goimports" -local $module -w ./pkg

echo "Generating all CRDs as v1beta1"
"$dir/bin/controller-gen" \
  crd \
  paths=./pkg/apis/... \
  output:dir=docs/crd \
  crd:crdVersions=v1beta1

echo "Generating infrastructure.giantswarm.io CRDs as v1"
"$dir/bin/controller-gen" \
  crd \
  paths=./pkg/apis/infrastructure/v1alpha2 \
  output:dir=docs/crd \
  crd:crdVersions=v1

echo "Kustomizing CRDs"
"$dir/bin/kustomize" build \
  config/crd \
  -o config/crd/bases/release.giantswarm.io_releases.yaml

echo "Using pkger to package CRDs into go source virtual file system"
"$dir/bin/pkger"
