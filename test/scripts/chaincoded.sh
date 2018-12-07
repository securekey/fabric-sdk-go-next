#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e

GO_CMD="${GO_CMD:-go}"
GOPATH="${GOPATH:-$HOME/go}"

mkdir -p ${GOPATH}/src/github.com/hyperledger
ln -s ${GOPATH}/src/chaincoded/vendor/github.com/hyperledger/fabric ${GOPATH}/src/github.com/hyperledger/fabric

echo "Installing chaincodes ..."
go install github.com/example_cc
go install github.com/example_pvt_cc
go install chaincoded/cmd/chaincoded

PEERS=(
    peer0.org1.example.com:7052
    peer1.org1.example.com:7152
    peer0.org2.example.com:8052
    peer1.org2.example.com:9052
)

# Chaincode logs are controlled via the CORE_CHAINCODE_LOGGING_LEVEL environment variable.
#export CORE_CHAINCODE_LOGGING_LEVEL=debug
# You can set CHAINCODED_VERBOSE environment variable to see additional chaincoded logs.
#export CHAINCODED_VERBOSE=true

echo "Running chaincoded ..."
chaincoded ":9375" ${PEERS[@]}