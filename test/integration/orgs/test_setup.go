/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package orgs

import (
	"testing"
	"time"

	"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	ca "github.com/hyperledger/fabric-sdk-go/api/apifabca"
	fab "github.com/hyperledger/fabric-sdk-go/api/apifabclient"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn"

	packager "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/ccpackager/gopackager"

	chmgmt "github.com/hyperledger/fabric-sdk-go/api/apitxn/chmgmtclient"
	resmgmt "github.com/hyperledger/fabric-sdk-go/api/apitxn/resmgmtclient"
	deffab "github.com/hyperledger/fabric-sdk-go/def/fabapi"
	"github.com/hyperledger/fabric-sdk-go/pkg/config"
	cryptosuite "github.com/hyperledger/fabric-sdk-go/pkg/cryptosuite/bccsp"
	"github.com/hyperledger/fabric-sdk-go/pkg/errors"
	client "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/events"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/orderer"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/signingmgr"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-txn/admin"
	"github.com/hyperledger/fabric-sdk-go/test/integration"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
)

var org1 = "Org1"
var org2 = "Org2"

// Client
var orgTestClient fab.FabricClient

// Channel
var orgTestChannel fab.Channel

// Orderers
var orgTestOrderer fab.Orderer

// Peers
var orgTestPeer0 fab.Peer
var orgTestPeer1 fab.Peer

// EventHubs
var peer0EventHub fab.EventHub
var peer1EventHub fab.EventHub

// Users
var org1AdminUser ca.User
var org2AdminUser ca.User
var ordererAdminUser ca.User
var org1User ca.User
var org2User ca.User

// Flag to indicate if test has run before (to skip certain steps)
var foundChannel bool

// Org1 and Org2 resource management clients
var org1ResMgmt resmgmt.ResourceMgmtClient
var org2ResMgmt resmgmt.ResourceMgmtClient

// initializeFabricClient initializes fabric-sdk-go
func initializeFabricClient(t *testing.T) {
	// Initialize configuration
	configImpl, err := config.InitConfig("../" + integration.ConfigTestFile)
	if err != nil {
		t.Fatal(err)
	}

	// Instantiate client
	fcClient := client.NewClient(configImpl)

	// Initialize crypto suite
	cryptoSuiteprovider, err := cryptosuite.GetSuiteByConfig(configImpl)
	if err != nil {
		t.Fatal(err)
	}

	fcClient.SetCryptoSuite(cryptoSuiteprovider)

	signingMgr, err := signingmgr.NewSigningManager(cryptoSuiteprovider, configImpl)
	if err != nil {
		t.Fatal(err)
	}

	fcClient.SetSigningManager(signingMgr)

	// From now on use interface only
	orgTestClient = fcClient
}

func createTestChannel(t *testing.T, sdk *deffab.FabricSDK) {
	var err error

	orgTestChannel, err = channel.NewChannel("orgchannel", orgTestClient)
	if err != nil {
		t.Fatal(err)
	}

	orgTestChannel.AddPeer(orgTestPeer0)
	orgTestChannel.AddPeer(orgTestPeer1)
	orgTestChannel.SetPrimaryPeer(orgTestPeer0)

	orgTestChannel.AddOrderer(orgTestOrderer)

	orgTestClient.SetUserContext(org1User)

	foundChannel, err = integration.HasPrimaryPeerJoinedChannel(orgTestClient, orgTestChannel)
	if err != nil {
		t.Fatal(err)
	}

	if foundChannel {
		return
	}

	// Channel management client is responsible for managing channels (create/update channel)
	chMgmtClient, err := sdk.NewChannelMgmtClientWithOpts("Admin", &deffab.ChannelMgmtClientOpts{OrgName: "ordererorg"})
	if err != nil {
		t.Fatal(err)
	}

	// Create channel (or update if it already exists)
	req := chmgmt.SaveChannelRequest{ChannelID: "orgchannel", ChannelConfig: "../../fixtures/channel/orgchannel.tx", SigningUser: org1AdminUser}
	if err = chMgmtClient.SaveChannel(req); err != nil {
		t.Fatal(err)
	}

	// Allow orderer to process channel creation
	time.Sleep(time.Second * 3)
}

func joinTestChannel(t *testing.T, sdk *deffab.FabricSDK) {
	if foundChannel {
		return
	}

	var err error

	// Org1 resource management client (Org1 is default org)
	org1ResMgmt, err = sdk.NewResourceMgmtClient("Admin")
	if err != nil {
		t.Fatalf("Failed to create new resource management client: %s", err)
	}

	// Org1 peers join channel
	if err = org1ResMgmt.JoinChannel("orgchannel"); err != nil {
		t.Fatalf("Org1 peers failed to JoinChannel: %s", err)
	}

	// Org2 resource management client
	org2ResMgmt, err = sdk.NewResourceMgmtClientWithOpts("Admin", &deffab.ResourceMgmtClientOpts{OrgName: "Org2"})
	if err != nil {
		t.Fatal(err)
	}

	// Org2 peers join channel
	if err = org2ResMgmt.JoinChannel("orgchannel"); err != nil {
		t.Fatalf("Org2 peers failed to JoinChannel: %s", err)
	}

}

func installAndInstantiate(t *testing.T) {
	if foundChannel {
		return
	}

	ccPkg, err := packager.NewCCPackage("github.com/example_cc", "../../fixtures/testdata")
	if err != nil {
		t.Fatal(err)
	}

	req := resmgmt.InstallCCRequest{Name: "exampleCC", Path: "github.com/example_cc", Version: "0", Package: ccPkg}

	// Install example cc for Org1
	_, err = org1ResMgmt.InstallCC(req)
	if err != nil {
		t.Fatal(err)
	}

	// Install example cc for Org2
	_, err = org2ResMgmt.InstallCC(req)
	if err != nil {
		t.Fatal(err)
	}

	chaincodePolicy := cauthdsl.SignedByAnyMember([]string{
		org1AdminUser.MspID(), org2AdminUser.MspID()})

	orgTestClient.SetUserContext(org2AdminUser)
	err = admin.SendInstantiateCC(orgTestChannel, "exampleCC",
		integration.ExampleCCInitArgs(), "github.com/example_cc", "0", chaincodePolicy, []apitxn.ProposalProcessor{orgTestPeer1}, peer1EventHub)
	if err != nil {
		t.Fatal(err)
	}
}

func loadOrderer(t *testing.T) {
	ordererConfig, err := orgTestClient.Config().RandomOrdererConfig()
	if err != nil {
		t.Fatal(err)
	}

	orgTestOrderer, err = orderer.NewOrdererFromConfig(ordererConfig, orgTestClient.Config())
	if err != nil {
		t.Fatal(err)
	}
}

func loadOrgPeers(t *testing.T) {
	org1Peers, err := orgTestClient.Config().PeersConfig(org1)
	if err != nil {
		t.Fatal(err)
	}

	org2Peers, err := orgTestClient.Config().PeersConfig(org2)
	if err != nil {
		t.Fatal(err)
	}

	orgTestPeer0, err = peer.NewPeerFromConfig(&apiconfig.NetworkPeer{PeerConfig: org1Peers[0]}, orgTestClient.Config())
	if err != nil {
		t.Fatal(err)
	}

	orgTestPeer1, err = peer.NewPeerFromConfig(&apiconfig.NetworkPeer{PeerConfig: org2Peers[0]}, orgTestClient.Config())
	if err != nil {
		t.Fatal(err)
	}

	peer0EventHub, err = events.NewEventHub(orgTestClient)
	if err != nil {
		t.Fatal(err)
	}

	// TODO: See if required after events merge
	serverHostOverrideOrg1 := ""
	if str, ok := org1Peers[0].GRPCOptions["ssl-target-name-override"].(string); ok {
		serverHostOverrideOrg1 = str
	}
	peer0EventHub.SetPeerAddr(org1Peers[0].EventURL, org1Peers[0].TLSCACerts.Path, serverHostOverrideOrg1)

	orgTestClient.SetUserContext(org1User)
	err = peer0EventHub.Connect()
	if err != nil {
		t.Fatal(err)
	}

	peer1EventHub, err = events.NewEventHub(orgTestClient)
	if err != nil {
		t.Fatal(err)
	}

	// TODO: See if required after events merge
	serverHostOverrideOrg2 := ""
	if str, ok := org2Peers[0].GRPCOptions["ssl-target-name-override"].(string); ok {
		serverHostOverrideOrg2 = str
	}
	peer1EventHub.SetPeerAddr(org2Peers[0].EventURL, org2Peers[0].TLSCACerts.Path, serverHostOverrideOrg2)

	orgTestClient.SetUserContext(org2User)
	err = peer1EventHub.Connect()
	if err != nil {
		t.Fatal(err)
	}
}

// loadOrgUsers Loads all the users required to perform this test
func loadOrgUsers(t *testing.T) {
	var err error

	// Create SDK setup for the integration tests
	sdkOptions := deffab.Options{
		ConfigFile: "../" + integration.ConfigTestFile,
	}

	sdk, err := deffab.NewSDK(sdkOptions)
	if err != nil {
		t.Fatal(err)
	}

	ordererAdminUser = loadOrgUser(t, sdk, "ordererorg", "Admin")

	org1AdminUser = loadOrgUser(t, sdk, org1, "Admin")
	org2AdminUser = loadOrgUser(t, sdk, org2, "Admin")

	org1User = loadOrgUser(t, sdk, org1, "User1")
	org2User = loadOrgUser(t, sdk, org2, "User1")

}

func loadOrgUser(t *testing.T, sdk *deffab.FabricSDK, orgName string, userName string) fab.User {

	user, err := sdk.NewPreEnrolledUser(orgName, userName)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "NewPreEnrolledUser failed, %s, %s", orgName, userName))
	}

	return user

}
