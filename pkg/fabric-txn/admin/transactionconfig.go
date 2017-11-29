/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package admin

import (
	"time"

	fab "github.com/hyperledger/fabric-sdk-go/api/apifabclient"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"

	"github.com/hyperledger/fabric-sdk-go/pkg/errors"
	internal "github.com/hyperledger/fabric-sdk-go/pkg/fabric-txn/internal"
	"github.com/hyperledger/fabric-sdk-go/pkg/logging"
)

var logger = logging.NewLogger("fabric_sdk_go")

// SendInstantiateCC Sends instantiate CC proposal to one or more endorsing peers
func SendInstantiateCC(channel fab.Channel, chainCodeID string, args [][]byte,
	chaincodePath string, chaincodeVersion string, chaincodePolicy *common.SignaturePolicyEnvelope, targets []apitxn.ProposalProcessor, eventHub fab.EventHub) error {

	transactionProposalResponse, txID, err := channel.SendInstantiateProposal(chainCodeID,
		args, chaincodePath, chaincodeVersion, chaincodePolicy, targets)
	if err != nil {
		return errors.WithMessage(err, "SendInstantiateProposal failed")
	}

	for _, v := range transactionProposalResponse {
		if v.Err != nil {
			logger.Debugf("SendInstantiateProposal endorser %s returned error", v.Endorser)
			return errors.WithMessage(v.Err, "SendInstantiateProposal endorser failed")
		}
		logger.Debug("SendInstantiateProposal endorser '%s' returned ProposalResponse status:%v", v.Endorser, v.Status)
	}

	// Register for commit event
	chcode := internal.RegisterTxEvent(txID, eventHub)

	if _, err = internal.CreateAndSendTransaction(channel, transactionProposalResponse); err != nil {
		return errors.WithMessage(err, "CreateAndSendTransaction failed")
	}

	select {
	case code := <-chcode:
		if code == peer.TxValidationCode_VALID {
			return nil
		}
		logger.Debugf("instantiateCC error received from eventhub for txid(%s), code(%s)", txID, code)
		return errors.Errorf("instantiateCC with code %s", code)
	case <-time.After(time.Second * 30):
		logger.Debugf("instantiateCC didn't receive block event for txid(%s)", txID)
		return errors.New("instantiateCC timeout")
	}
}

// SendUpgradeCC Sends upgrade CC proposal to one or more endorsing peers
func SendUpgradeCC(channel fab.Channel, chainCodeID string, args [][]byte,
	chaincodePath string, chaincodeVersion string, chaincodePolicy *common.SignaturePolicyEnvelope, targets []apitxn.ProposalProcessor, eventHub fab.EventHub) error {

	transactionProposalResponse, txID, err := channel.SendUpgradeProposal(chainCodeID,
		args, chaincodePath, chaincodeVersion, chaincodePolicy, targets)
	if err != nil {
		return errors.WithMessage(err, "SendUpgradeProposal failed")
	}

	for _, v := range transactionProposalResponse {
		if v.Err != nil {
			logger.Debugf("SendUpgradeProposal endorser %s failed", v.Endorser)
			return errors.WithMessage(v.Err, "SendUpgradeProposal endorser failed")
		}
		logger.Debug("SendUpgradeProposal Endorser '%s' returned ProposalResponse status:%v\n", v.Endorser, v.Status)
	}

	// Register for commit event
	chcode := internal.RegisterTxEvent(txID, eventHub)

	if _, err = internal.CreateAndSendTransaction(channel, transactionProposalResponse); err != nil {
		return errors.WithMessage(err, "CreateAndSendTransaction failed")
	}

	select {
	case code := <-chcode:
		if code == peer.TxValidationCode_VALID {
			return nil
		}
		logger.Debugf("upgradeCC Error received from eventhub for txid(%s) code(%s)", txID, code)
		return errors.Errorf("upgradeCC failed with code %s", code)
	case <-time.After(time.Second * 30):
		logger.Debugf("instantiateCC didn't receive block event for txid(%s)", txID)
		return errors.New("upgradeCC timeout")
	}
}
