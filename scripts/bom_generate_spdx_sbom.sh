#!/bin/bash

IMAGE_NAME=ghcr.io/timoniersystems/gocobra
IMAGE_TAG=2.5.0

bom generate --image ${IMAGE_NAME}:${IMAGE_TAG} --format json > bom_spdx_sbom.json

