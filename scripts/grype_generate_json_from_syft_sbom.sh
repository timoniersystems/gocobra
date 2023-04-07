#!/bin/bash

SBOM=$1
REPORT=grype_report.json
grype -o json --file $REPORT sbom:$SBOM
