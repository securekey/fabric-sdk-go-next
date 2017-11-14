#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

cd $TMP_PROJECT_PATH

echo "Applying cherry picks (channel event client) ..."
git fetch https://gerrit.hyperledger.org/r/fabric refs/changes/83/15183/2 && git cherry-pick FETCH_HEAD
git fetch https://gerrit.hyperledger.org/r/fabric refs/changes/71/15171/4 && git cherry-pick FETCH_HEAD

echo "Applying cherry picks (Private Data - Phase 3) ..."
git fetch https://gerrit.hyperledger.org/r/fabric refs/changes/15/14515/8 && git cherry-pick FETCH_HEAD
git fetch https://gerrit.hyperledger.org/r/fabric refs/changes/17/14517/6 && git cherry-pick FETCH_HEAD
git fetch https://gerrit.hyperledger.org/r/fabric refs/changes/91/14291/24 && git cherry-pick FETCH_HEAD
git fetch https://gerrit.hyperledger.org/r/fabric refs/changes/67/14367/15 && git cherry-pick FETCH_HEAD
git fetch https://gerrit.hyperledger.org/r/fabric refs/changes/71/14371/20 && git cherry-pick FETCH_HEAD
