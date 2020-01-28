#!/usr/bin/env bash

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

bash "$(go list -m -f '{{.Dir}}' k8s.io/code-generator)/generate-groups.sh"  \
    "deepcopy,client" \
    github.com/giantswarm/apiextensions/pkg \
    github.com/giantswarm/apiextensions/pkg/apis \
    "application:v1alpha1 core:v1alpha1 example:v1alpha1 provider:v1alpha1 release:v1alpha1 infrastructure:v1alpha2" \
    --output-base "$dir" \
    --go-header-file "${dir}/boilerplate.go.txt"

cp -R "$dir/github.com/giantswarm/apiextensions/pkg"/* "$dir/../pkg"
rm -rf "$dir/github.com"
