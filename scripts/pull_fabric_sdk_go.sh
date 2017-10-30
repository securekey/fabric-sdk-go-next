#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

MY_PATH="`dirname \"$0\"`"              # relative
MY_PATH="`( cd \"$MY_PATH\" && pwd )`"  # absolutized and normalized
if [ -z "$MY_PATH" ] ; then
  # error; for some reason, the path is not accessible
  # to the script (e.g. permissions re-evaled after suid)
  exit 1  # fail
fi

rm -rf $GOPATH/src/github.com/hyperledger/fabric-sdk-go
mkdir -p $GOPATH/src/github.com/hyperledger/fabric-sdk-go
cd $GOPATH/src/github.com/hyperledger/
git clone https://gerrit.hyperledger.org/r/fabric-sdk-go
cd fabric-sdk-go
git checkout aafbea28037ef0f2e8c1d3bbef42671d5525e81a

##WIP - [FAB-6243] - Channel Event Client
#https://gerrit.hyperledger.org/r/#/c/13673/ - [FAB-6243] - Channel Event Client
git fetch https://gerrit.hyperledger.org/r/fabric-sdk-go refs/changes/73/13673/14 && git cherry-pick FETCH_HEAD


cp $MY_PATH/.env $GOPATH/src/github.com/hyperledger/fabric-sdk-go/test/fixtures
export FABRIC_SDK_EXTRA_GO_TAGS=channelevents
make all