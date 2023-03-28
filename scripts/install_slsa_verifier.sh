#!/bin/bash

VERSION=2.1.0
rm -rf slsa-verifier-linux-amd64
wget https://github.com/slsa-framework/slsa-verifier/releases/download/v${VERSION}/slsa-verifier-linux-amd64
chmod +x slsa-verifier-linux-amd64
mv slsa-verifier-linux-amd64 ~/.local/bin/slsa-verifier
