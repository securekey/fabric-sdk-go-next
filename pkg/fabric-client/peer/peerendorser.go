/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package peer

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn"
	"github.com/hyperledger/fabric-sdk-go/pkg/config/urlutil"
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"

	"crypto/tls"

	"github.com/hyperledger/fabric-sdk-go/pkg/errors"
)

// peerEndorser enables access to a GRPC-based endorser for running transaction proposal simulations
type peerEndorser struct {
	grpcDialOption []grpc.DialOption
	target         string
}

// TransactionProposalError represents an error condition that prevented proposal processing.
type TransactionProposalError struct {
	Endorser string
	Proposal apitxn.TransactionProposal
	Err      error
}

func newPeerEndorser(target string, certificate string, serverHostOverride string,
	dialBlocking bool, config apiconfig.Config) (
	peerEndorser, error) {
	if len(target) == 0 {
		return peerEndorser{}, errors.New("target is required")
	}

	// Construct dialer options for the connection
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTimeout(config.TimeoutOrDefault(apiconfig.Endorser)))
	if dialBlocking { // TODO: configurable?
		opts = append(opts, grpc.WithBlock())
	}

	if urlutil.IsTLSEnabled(target) {
		// get the root cert pool
		certPool, _ := config.TLSCACertPool("")
		if len(certificate) == 0 && len(certPool.Subjects()) == 0 {
			return peerEndorser{}, errors.New("certificate is required")
		}

		tlsCaCertPool, err := config.TLSCACertPool(certificate)
		if err != nil {
			return peerEndorser{}, err
		}

		clientConfig, err := config.Client()
		if err != nil {
			return peerEndorser{}, err
		}

		var certificates []tls.Certificate

		if clientConfig.TLSCerts.Client.CertPem != "" {
			clientCerts, err := tls.X509KeyPair([]byte(clientConfig.TLSCerts.Client.CertPem), []byte(clientConfig.TLSCerts.Client.KeyPem))
			if err != nil {
				return peerEndorser{}, errors.Errorf("Error loading cert/key pair as TLS client credentials: %v", err)
			}
			certificates = []tls.Certificate{clientCerts}
		} else if clientConfig.TLSCerts.Client.Certfile != "" {
			clientCerts, err := tls.LoadX509KeyPair(clientConfig.TLSCerts.Client.Certfile, clientConfig.TLSCerts.Client.Keyfile)
			if err != nil {
				return peerEndorser{}, errors.Errorf("Error loading cert/key pair as TLS client credentials: %v", err)
			}
			certificates = []tls.Certificate{clientCerts}
		}

		creds := credentials.NewTLS(&tls.Config{
			Certificates: certificates,
			RootCAs:      tlsCaCertPool,
			ServerName:   serverHostOverride,
		})

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	pc := peerEndorser{grpcDialOption: opts, target: urlutil.ToAddress(target)}

	return pc, nil
}

// ProcessTransactionProposal sends the transaction proposal to a peer and returns the response.
func (p *peerEndorser) ProcessTransactionProposal(proposal apitxn.TransactionProposal) (apitxn.TransactionProposalResult, error) {
	logger.Debugf("Processing proposal using endorser :%s", p.target)

	proposalResponse, err := p.sendProposal(proposal)
	if err != nil {
		tpe := TransactionProposalError{
			Endorser: p.target,
			Proposal: proposal,
			Err:      err,
		}
		return apitxn.TransactionProposalResult{}, &tpe
	}

	return apitxn.TransactionProposalResult{
		Proposal:         proposal,
		ProposalResponse: proposalResponse,
		Endorser:         p.target, // TODO: what format is expected for Endorser? Just target? URL?
		Status:           proposalResponse.GetResponse().Status,
	}, nil
}

func (p *peerEndorser) conn() (*grpc.ClientConn, error) {
	return grpc.Dial(p.target, p.grpcDialOption...)
}

func (p *peerEndorser) releaseConn(conn *grpc.ClientConn) {
	conn.Close()
}

func (p *peerEndorser) sendProposal(proposal apitxn.TransactionProposal) (*pb.ProposalResponse, error) {
	conn, err := p.conn()
	if err != nil {
		return nil, err
	}
	defer p.releaseConn(conn)

	endorserClient := pb.NewEndorserClient(conn)
	return endorserClient.ProcessProposal(context.Background(), proposal.SignedProposal)
}

func (tpe TransactionProposalError) Error() string {
	return fmt.Sprintf("Transaction processor (%s) returned error '%s' for txID '%s'",
		tpe.Endorser, tpe.Err.Error(), tpe.Proposal.TxnID.ID)
}
