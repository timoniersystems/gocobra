#!/bin/bash

VERSION=0.5.1
BINARY=bom-amd64-linux

rm -rf $BINARY
wget https://github.com/kubernetes-sigs/bom/releases/download/v${VERSION}/$BINARY

chmod +x $BINARY
mv $BINARY ~/.local/bin/bom
