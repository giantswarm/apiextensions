#!/usr/bin/env bash

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOPATH=$(go env GOPATH) # ensure that GOPATH is defined for generate-groups.sh
go mod tidy # ensure that code-generator is cached in $GOPATH/pkg/mod for generate-groups.sh

bash "$(go list -m -f '{{.Dir}}' k8s.io/code-generator)/generate-groups.sh"  \
    "deepcopy,client" \
    github.com/giantswarm/apiextensions/pkg \
    github.com/giantswarm/apiextensions/pkg/apis \
    "application:v1alpha1 core:v1alpha1 example:v1alpha1 provider:v1alpha1 release:v1alpha1 infrastructure:v1alpha2 backup:v1alpha1 tooling:v1alpha1" \
    --output-base "$dir" \
    --go-header-file "${dir}/boilerplate.go.txt"

cp -R "$dir/github.com/giantswarm/apiextensions/pkg"/* "$dir/../pkg"
rm -rf "$dir/github.com"

# code-generator doesn't group local imports separately from third-party
# imports, so run goimports after generating new code.
cd "$dir/.." || exit
if [ ! -x "$(command -v goimports)" ]; then
  echo "goimports not found, downloading to $GOPATH/bin using 'go get'"
  go get golang.org/x/tools/cmd/goimports
  export PATH=$PATH:$GOPATH/bin
fi
echo "Fixing imports in-place with goimports"
goimports -local github.com/giantswarm/apiextensions -w ./pkg
