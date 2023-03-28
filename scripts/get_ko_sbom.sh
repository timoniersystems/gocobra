#!/bin/bash

IMAGE=ghcr.io/timoniersystems/gocobra
TAG=2.0.4
IMAGE_DIGEST=$(crane digest $IMAGE:$TAG)

cosign download sbom $IMAGE@${IMAGE_DIGEST}

