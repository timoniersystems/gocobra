#!/bin/bash

IMAGE_TAG=2.2.7

COSIGN_PUBLIC_KEY=$HOME/.cosign/cosign.pub
IMAGE_NAME=ghcr.io/timoniersystems/gocobra

cosign verify --key $COSIGN_PUBLIC_KEY ${IMAGE_NAME}:${IMAGE_TAG}

### Triangulation

# show signature artifact
cosign triangulate ${IMAGE_NAME}:${IMAGE_TAG}

# show signature manifest
crane manifest $(cosign triangulate ${IMAGE_NAME}:${IMAGE_TAG}) | jq .
