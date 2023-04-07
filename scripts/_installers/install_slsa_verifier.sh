#!/bin/bash

VERSION=2.1.0
BINARY=slsa-verifier-linux-amd64

rm -rf $BINARY
wget https://github.com/slsa-framework/slsa-verifier/releases/download/v${VERSION}/$BINARY
chmod +x $BINARY
mv $BINARY ~/.local/bin/slsa-verifier
