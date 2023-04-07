#!/bin/bash

IMAGE_NAME=ghcr.io/timoniersystems/gocobra
IMAGE_TAG=2.5.0

syft packages ${IMAGE_NAME}:${IMAGE_TAG} -o json > syft_sbom.json

