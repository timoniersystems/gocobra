#!/bin/bash

IMAGE_TAG=2.5.0

COSIGN_PUBLIC_KEY=../.github/cosign.pub
PROJECT_NAME=gocobra
IMAGE_NAME=ghcr.io/timoniersystems/${PROJECT_NAME}

cosign verify --key $COSIGN_PUBLIC_KEY ${IMAGE_NAME}:${IMAGE_TAG}

### Triangulation

# show signature artifact
cosign triangulate ${IMAGE_NAME}:${IMAGE_TAG}

# show signature manifest
echo Cosign signature manifest:
crane manifest $(cosign triangulate ${IMAGE_NAME}:${IMAGE_TAG}) | jq -r . | tee ${PROJECT_NAME}:${IMAGE_TAG}-manifest.json

# get transparency log (rekor) index from manifest
cat ${PROJECT_NAME}:${IMAGE_TAG}-manifest.json | jq '.layers[0].annotations."dev.sigstore.cosign/bundle"' |sed 's/\\//g' > ${PROJECT_NAME}:${IMAGE_TAG}-bundle.json
LOG_INDEX=$(cat ${PROJECT_NAME}:${IMAGE_TAG}-bundle.json | cut -d , -f 4 | cut -d : -f 2)

# get signature from rekor record
echo Rekor record for index $LOG_INDEX:
rekor-cli get --log-index $LOG_INDEX --format json
SIG_FROM_REKOR=$(rekor-cli get --log-index $LOG_INDEX --format json | jq -r '.Body.HashedRekordObj.signature.content')
echo -e "SIG_FROM_REKOR:\t\t$SIG_FROM_REKOR"

# get signature from Cosign manifest
SIG_FROM_COSIGN=$(cat ${PROJECT_NAME}:${IMAGE_TAG}-manifest.json | jq -r '.layers[0].annotations."dev.cosignproject.cosign/signature"')
echo -e "SIG_FROM_COSIGN:\t$SIG_FROM_COSIGN"
