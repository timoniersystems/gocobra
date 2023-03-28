#!/bin/bash

COSIGN_PUBLIC_KEY=$HOME/.cosign/cosign.pub
IMAGE=ghcr.io/timoniersystems/gocobra:2.0.8

cosign verify --key $COSIGN_PUBLIC_KEY "$IMAGE"

### Triangulation

# show signature artifact
cosign triangulate $IMAGE

# show signature manifest
crane manifest $(cosign triangulate $IMAGE) | jq .
