#!/bin/bash

SOURCE_TAG=2.5.0
SOURCE_URI=github.com/timoniersystems/gocobra
COSIGN_PUBLIC_KEY=../.github/cosign.pub

rm -rf checksums.txt checksums.txt.sig

wget https://${SOURCE_URI}/releases/download/${SOURCE_TAG}/checksums.txt
wget https://${SOURCE_URI}/releases/download/${SOURCE_TAG}/checksums.txt.sig

cosign verify-blob \
  --key $COSIGN_PUBLIC_KEY \
  --signature checksums.txt.sig \
  --insecure-ignore-tlog \
  checksums.txt

# now download release files and verify the checksum hashes

RELEASE_FILE_TYPES=$(cat checksums.txt | cut -f 3,4 -d '_')
for t in $RELEASE_FILE_TYPES; do
  rm -rf gocobra_${SOURCE_TAG}_$t
  wget https://${SOURCE_URI}/releases/download/${SOURCE_TAG}/gocobra_${SOURCE_TAG}_$t
done

sha256sum --ignore-missing -c checksums.txt
