#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
toolpath="$dir/bin"

export GOPATH=$(go env GOPATH)
export GOROOT=$(go env GOROOT)

# We need specific versions of code generation tools so we install them locally
# in scripts/bin using versions from scripts/go.mod
install_tool() {
  local module=$1
  local package=$2
  local bin=$3
  pushd "$dir" > /dev/null
  version=$(go list -m -f '{{.Version}}' "$module")
  echo "Rebuilding $bin@$version"
  mkdir -p "$toolpath"
  go build -o "$toolpath/$bin" "$module""$package"
  popd > /dev/null
}

install_tool sigs.k8s.io/controller-tools /cmd/controller-gen controller-gen
install_tool k8s.io/code-generator /cmd/deepcopy-gen deepcopy-gen
install_tool k8s.io/code-generator /cmd/client-gen client-gen
install_tool golang.org/x/tools /cmd/goimports goimports
install_tool github.com/markbates/pkger /cmd/pkger pkger
install_tool sigs.k8s.io/kustomize/kustomize/v3 "" kustomize

# Set up variables for deepcopy-gen and client-gen
module="github.com/giantswarm/apiextensions"
input_dirs=$(find ./pkg/apis -maxdepth 2 -mindepth 2 | tr '\r\n' ',')
input_dirs=${input_dirs%?}
groups=${input_dirs//.\/pkg\/apis\//}
header="${dir}/boilerplate.go.txt"

# deepcopy-gen creates DeepCopy functions for each custom resource
echo "Generating deepcopy funcs"
"$toolpath/deepcopy-gen" \
  --input-dirs "$input_dirs" \
  --output-file-base zz_generated.deepcopy \
  --go-header-file "${dir}/boilerplate.go.txt"

# client-gen creates typed go clients for CRUD operations for each custom resource
echo "Generating clientset"
"$toolpath/client-gen" \
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
"$toolpath/goimports" -local $module -w ./pkg

# Ensure that we have fresh CRDs in case any have been deleted
rm -rf "$dir/../config/crd/v1"
rm -rf "$dir/../config/crd/v1beta1"

# Using kubebuilder comments, create new CRDs from CR definitions in source files
echo "Generating all CRDs"
"$toolpath/controller-gen" \
  crd \
  paths=./pkg/apis/... \
  output:dir=config/crd/v1 \
  crd:crdVersions=v1
"$toolpath/controller-gen" \
  crd \
  paths=./pkg/apis/... \
  output:dir=config/crd/v1beta1 \
  crd:crdVersions=v1beta1

# Add .metadata.name validation to Release CRD using kustomize since
# kubebuilder comments can't modify metav1.ObjectMeta
echo "Kustomizing CRDs"
for version in v1 v1beta1; do
  for crd in "config/crd/patches/$version"/*; do
    kustomize --load_restrictor LoadRestrictionsNone build \
      "$crd" \
      -o "config/crd/$version/$(basename "$crd").yaml"
  done
done

# Package CRD YAMLs into a virtual filesystem so they can be accessed by New*CRD() functions
hash=$(find config/crd -type f -print0 | xargs -0 sha1sum | sort -df | sha1sum)
prevhash=$(cat "$dir/.hash" 2> /dev/null || echo "")
if [ "$hash" != "$prevhash" ]; then
  echo "Detected changes in CRD YAMLs"
  echo "Using pkger to package CRDs into go source virtual file system"
  "$toolpath/pkger" -include /config/crd/v1 -include /config/crd/v1beta1 -o pkg/crd
  echo "$hash" > "$dir/.hash"
fi

echo "Applying linter patch to generated files"
git apply "$dir/generated.patch"
