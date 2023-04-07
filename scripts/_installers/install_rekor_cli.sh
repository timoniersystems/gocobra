#!/bin/bash

VERSION=1.1.0
BINARY=rekor-cli-linux-amd64

rm -rf $BINARY
wget https://github.com/sigstore/rekor/releases/download/v${VERSION}/$BINARY
chmod +x $BINARY
mv $BINARY ~/.local/bin/rekor-cli
