#!/bin/bash

SOURCE_URI=github.com/timoniersystems/gocobra
SOURCE_TAG=2.0.4
ARTIFACT_FILE=gocobra-linux-amd64-slsa3
INTOTO_ATTESTATION=gocobra-linux-amd64-slsa3.intoto.jsonl

wget https://${SOURCE_URI}/releases/download/${SOURCE_TAG}/${ARTIFACT_FILE}
wget https://${SOURCE_URI}/releases/download/${SOURCE_TAG}/${INTOTO_ATTESTATION}

~/go/bin/slsa-verifier verify-artifact $ARTIFACT_FILE \
  --provenance-path $INTOTO_ATTESTATION \
  --source-uri $SOURCE_URI \
  --source-tag $SOURCE_TAG \
  --print-provenance
