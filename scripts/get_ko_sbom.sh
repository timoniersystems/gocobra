#!/bin/bash

TAG=2.5.0
IMAGE=ghcr.io/timoniersystems/gocobra
IMAGE_DIGEST=$(crane digest $IMAGE:$TAG)

cosign download sbom $IMAGE@${IMAGE_DIGEST} | tee ko_sbom_spdx.json

