#!/bin/bash
set -euo pipefail

OUTPUT=test

## cluster-api
REPO=kubernetes-sigs/cluster-api
FILE=cluster-api-components.yaml
TAG=v0.3.14

echo "> fetching CRDs from ${REPO}@${TAG}"
FILE_URL=$(curl "https://api.github.com/repos/${REPO}/releases/tags/${TAG}" | \
	yq -r '.assets[] | select(.name == "'"${FILE}"'") | .browser_download_url')

mkdir -p ${OUTPUT}

curl -L "${FILE_URL}" | \
	yq --yaml-output -r 'select(.kind == "CustomResourceDefinition")' > ${OUTPUT}/${REPO/\//\-}_${TAG}.yaml
