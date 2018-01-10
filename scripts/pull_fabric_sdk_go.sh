#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#


export FABRIC_CA_FIXTURE_TAG=1.0.3
export FABRIC_ORDERER_FIXTURE_TAG=1.1.0-0.0.3-snapshot-2f6440d
export FABRIC_PEER_FIXTURE_TAG=1.1.0-0.0.3-snapshot-2f6440d
export FABRIC_COUCHDB_FIXTURE_TAG=1.1.0-0.0.3-snapshot-2f6440d
export FABRIC_BUILDER_FIXTURE_TAG=1.1.0-0.0.3-snapshot-2f6440d
export FABRIC_BASEOS_FIXTURE_TAG=0.4.2

export FABRIC_CA_FIXTURE_IMAGE=hyperledger/fabric-ca
export FABRIC_ORDERER_FIXTURE_IMAGE=next/hyperledger/fabric-orderer
export FABRIC_PEER_FIXTURE_IMAGE=next/hyperledger/fabric-peer
export FABRIC_COUCHDB_FIXTURE_IMAGE=next/hyperledger/fabric-couchdb
export FABRIC_BUILDER_FIXTURE_IMAGE=next/hyperledger/fabric-ccenv
export FABRIC_BASEOS_FIXTURE_IMAGE=hyperledger/fabric-baseos
export FABRIC_BASEIMAGE_FIXTURE_IMAGE=hyperledger/fabric-baseimage
export FABRIC_RELEASE_REGISTRY=repo.onetap.ca:8443
export FABRIC_DEV_REGISTRY=repo.onetap.ca:8443
export FABRIC_DEV_REGISTRY_PRE_CMD=
export FABRIC_STABLE_PKCS11_INTTEST=true
export FABRIC_STABLE_VERSION=1.1.0
export FABRIC_STABLE_VERSION_MINOR=1.1
export FABRIC_STABLE_VERSION_MAJOR=1
export FABRIC_SDK_EXTRA_GO_TAGS="devstable prerelease"


MY_PATH="`dirname \"$0\"`"              # relative
MY_PATH="`( cd \"$MY_PATH\" && pwd )`"  # absolutized and normalized
if [ -z "$MY_PATH" ] ; then
  # error; for some reason, the path is not accessible
  # to the script (e.g. permissions re-evaled after suid)
  exit 1  # fail
fi

TMP=`mktemp -d 2>/dev/null || mktemp -d -t 'mytmpdir'`

fabricSdkGoPath=$GOPATH/src/github.com/hyperledger/fabric-sdk-go

GOPATH=$TMP

mkdir -p $GOPATH/src/github.com/hyperledger/fabric-sdk-go
cd $GOPATH/src/github.com/hyperledger/
git clone https://gerrit.hyperledger.org/r/fabric-sdk-go
cd fabric-sdk-go
git checkout 188110a6d67e7298950bbdc01769008d2bb0e920

## logger fix to override deflogger for private-logger changes
git fetch https://gerrit.hyperledger.org/r/fabric-sdk-go refs/changes/81/16681/1 && git cherry-pick FETCH_HEAD

#export FABRIC_SDK_EXTRA_GO_TAGS=channelevents
make all

if [ -d "$fabricSdkGoPath" ]; then
echo "can not copy fabric-sdk-go already exist in GOPATH"
exit 1
fi
cp -r $GOPATH/src/github.com/hyperledger/fabric-sdk-go $fabricSdkGoPath

rm -Rf $TMP
