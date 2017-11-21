#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

CONFIGTXGEN_CMD="${CONFIGTXGEN_CMD:-configtxgen}"

declare -a channels=("mychannel" "orgchannel" "testchannel")
declare -a orgs=("Org1MSP" "Org2MSP")

export CHANNEL_DIR="/opt/gopath/src/github.com/hyperledger/fabric-sdk-go/test/fixtures/channel"
export FABRIC_CFG_PATH=${CHANNEL_DIR}

echo "Generating Orderer Genesis block"
$CONFIGTXGEN_CMD -profile TwoOrgsOrdererGenesis -outputBlock ${CHANNEL_DIR}/twoorgs.genesis.block

for i in "${channels[@]}"
do
   echo "Generating artifacts for channel: $i"

   echo "Generating channel configuration transaction"
   $CONFIGTXGEN_CMD -profile TwoOrgsChannel -outputCreateChannelTx .${CHANNEL_DIR}/${i}.tx -channelID $i

   for j in "${orgs[@]}"
   do
     echo "Generating anchor peer update for org $j"
     $CONFIGTXGEN_CMD -profile TwoOrgsChannel -outputAnchorPeersUpdate ${CHANNEL_DIR}/${i}${j}anchors.tx -channelID $i -asOrg $j
   done
done
