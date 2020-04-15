#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOPATH=$(go env GOPATH)
export GOROOT=$(go env GOROOT)

# We need specific versions of above tools so we install them locally
# in scripts/bin using versions from scripts/go.mod
install_tool() {
  local tool=$1
  pushd "$dir"
  base=$(echo "$tool" | sed "s/\/cmd\/.*//")
  version=$(go list -m -f '{{.Version}}' "$base")
  echo "Rebuilding $tool@$version"
  go build -o "$dir/bin" "$tool"
  popd
}

install_tool "sigs.k8s.io/controller-tools/cmd/controller-gen"
install_tool "k8s.io/code-generator/cmd/deepcopy-gen"
install_tool "k8s.io/code-generator/cmd/client-gen"
install_tool "golang.org/x/tools/cmd/goimports"
install_tool "github.com/markbates/pkger/cmd/pkger"
install_tool "sigs.k8s.io/kustomize/kustomize/v3"

# Set up variables for deepcopy-gen and client-gen
module="github.com/giantswarm/apiextensions"
input_dirs=$(find ./pkg/apis -maxdepth 2 -mindepth 2 | tr '\r\n' ',')
input_dirs=${input_dirs%?}
groups=${input_dirs//.\/pkg\/apis\//}
header="${dir}/boilerplate.go.txt"

# deepcopy-gen creates DeepCopy functions for each custom resource
echo "Generating deepcopy funcs"
"$dir/bin/deepcopy-gen" \
  --input-dirs "$input_dirs" \
  --output-file-base zz_generated.deepcopy \
  --output-base "$dir/.." \
  --go-header-file "${dir}/boilerplate.go.txt"

# client-gen creates typed go clients for CRUD operations for each custom resource
echo "Generating clientset"
"$dir/bin/client-gen" \
  --clientset-name versioned \
  --input "$groups" \
  --input-base "$module/pkg/apis" \
  --output-package "$module/pkg/clientset" \
  --output-base "$dir" \
  --go-header-file "$header"

# client-gen expects to be run in $GOPATH so it generates files in
# ./github.com/giantswarm/apiextensions/pkg/clientset which need to
# be manually moved into pkg.
echo "Moving generated files to expected location"
cp -R "$dir/$module/pkg"/* "$dir/../pkg"
rm -rf "$dir/github.com"

# code-generator doesn't group local imports separately from third-party
# imports, so run goimports after generating new code.
cd "$dir/.." || exit
echo "Fixing imports in-place with goimports"
"$dir/bin/goimports" -local $module -w ./pkg

# Ensure that we have fresh CRDs in case any have been deleted
rm -rf "$dir/../config/crd/bases"

# Using kubebuilder comments, create new CRDs from CR definitions in source files
echo "Generating all CRDs as v1beta1"
"$dir/bin/controller-gen" \
  crd \
  paths=./pkg/apis/... \
  output:dir=config/crd/bases \
  crd:crdVersions=v1beta1

# Overwrite CRDs infrastructure.giantswarm.io as v1 until all other
# groups can be migrated to v1
echo "Generating infrastructure.giantswarm.io CRDs as v1"
"$dir/bin/controller-gen" \
  crd \
  paths=./pkg/apis/infrastructure/v1alpha2 \
  output:dir=config/crd/bases \
  crd:crdVersions=v1

# Add .metadata.name validation to Release CRD using kustomize since
# kubebuilder comments can't modify metav1.ObjectMeta
echo "Kustomizing CRDs"
"$dir/bin/kustomize" build \
  config/crd \
  -o config/crd/bases/release.giantswarm.io_releases.yaml

# Package CRD YAMLs into a virtual filesystem so they can be accessed by New*CRD() functions
echo "Using pkger to package CRDs into go source virtual file system"
"$dir/bin/pkger" -include /config/crd/bases -o pkg/crd
