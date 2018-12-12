#!/usr/bin/env bash

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd ${dir}/../vendor/k8s.io/code-generator && ./generate-groups.sh \
    "deepcopy,client" \
    github.com/giantswarm/apiextensions/pkg \
    github.com/giantswarm/apiextensions/pkg/apis \
    "app:v1alpha1 core:v1alpha1 example:v1alpha1 provider:v1alpha1" \
    --go-header-file ${dir}/boilerplate.go.txt
