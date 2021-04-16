#!/bin/bash
set -euo pipefail

# Requirements:
#
# yq: https://github.com/kislyuk/yq

SCRIPT_DIR="$( cd "$( dirname "$0" )" && pwd )"

OUTPUT=$SCRIPT_DIR/../config/crd/
TMP_DIR=$(mktemp -d /tmp/apiextensions.XXXXXXXXXX)
cleanup() {
  rm -rf "${TMP_DIR}"
}
trap cleanup EXIT

fetch() {
	local REPO=$1
	local VERSION=$2
	local FILE=$3

	# remove "/" for output file name.
	local OUTPUT_FILE=${TMP_DIR}/${REPO/\//\-}.yaml

	echo "> fetching CRDs from ${REPO}@${VERSION} to ${OUTPUT_FILE}"
	FILE_URL=$(curl -sS "https://api.github.com/repos/${REPO}/releases/tags/${VERSION}" | \
		yq -r '.assets[] | select(.name == "'"${FILE}"'") | .browser_download_url')

	mkdir -p ${OUTPUT}

	# filter only CRD files.
	curl --progress-bar -L "${FILE_URL}" | \
		yq --yaml-output -r 'select(.kind == "CustomResourceDefinition")' > ${OUTPUT_FILE}

	# split the yaml into multiple files.
	kubernetes-split-yaml --outdir ${TMP_DIR} ${OUTPUT_FILE}
	rm ${OUTPUT_FILE}

	# rename the file with <group>_<kind>.yaml
	for file in $(ls ${TMP_DIR}); do
		GROUP=$(yq -r '.spec.group' ${TMP_DIR}/${file})
		KIND=$(yq -r '.spec.names.plural' ${TMP_DIR}/${file})
		API_VERSION=$(yq -r '.apiVersion' ${TMP_DIR}/${file} | cut -d/ -f2)
		mv ${TMP_DIR}/${file} ${OUTPUT}/${API_VERSION}/${GROUP}_${KIND}.yaml
	done
}

## cluster-api
REPO=kubernetes-sigs/cluster-api
VERSION=v0.3.14
FILE=cluster-api-components.yaml
fetch $REPO $VERSION $FILE

# cluster-api-provider-aws
REPO=kubernetes-sigs/cluster-api-provider-aws
VERSION=v0.6.5
FILE=infrastructure-components.yaml
fetch $REPO $VERSION $FILE

# cluster-api-provider-azure
REPO=kubernetes-sigs/cluster-api-provider-azure
VERSION=v0.4.12
FILE=infrastructure-components.yaml
fetch $REPO $VERSION $FILE

# cluster-api-provider-vsphere
REPO=kubernetes-sigs/cluster-api-provider-vsphere
VERSION=v0.7.6
FILE=infrastructure-components.yaml
fetch $REPO $VERSION $FILE

# aad-pod-identity
REPO=Azure/aad-pod-identity
VERSION=v1.7.4
FILE=deployment.yaml
