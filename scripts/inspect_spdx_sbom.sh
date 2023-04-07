#!/bin/bash

SBOM_FILE=$1

echo Inspecting SPDX SBOM ${SBOM_FILE}

#cat ${SBOM_FILE} | jq 
cat ${SBOM_FILE} | jq -r '.packages[]|.name,.externalRefs[0].referenceLocator'
