#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

cd $TMP_PROJECT_PATH

echo "Applying cherry picks (Private Data - Phase 3) ..."
git fetch https://gerrit.hyperledger.org/r/fabric refs/changes/15/14515/8 && git cherry-pick FETCH_HEAD
git fetch https://gerrit.hyperledger.org/r/fabric refs/changes/17/14517/6 && git cherry-pick FETCH_HEAD

