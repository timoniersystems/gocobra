#!/bin/bash

SBOM=$1
REPORT=grype_report.sarif
grype -o sarif --file $REPORT sbom:$SBOM
