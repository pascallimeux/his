package helpers

import (
	"testing"
	"time"
	"strings"
	sdkConfig "github.com/hyperledger/fabric-sdk-go/config"
)

func TestPeerconfig(t *testing.T) {
	peersConfig, _ := sdkConfig.GetPeersConfig()
	t.Log("***************************************nb peers:", len(peersConfig))
	for _,peer := range peersConfig {
		t.Log("Host: ", peer.Host)
		t.Log("Port: ", peer.Port)
		t.Log("EventHost: ", peer.EventHost)
		t.Log("EventPort: ", peer.EventPort)
		t.Log("Primary: ", peer.Primary)
	}

}

func TestChain(t *testing.T) {
	chainname := netHelper.Chain.GetName()
	orderers := netHelper.Chain.GetOrderers()
	peers := netHelper.Chain.GetPeers()
	ppeers := netHelper.Chain.GetPrimaryPeer()
	orgs, _ := netHelper.Chain.GetOrganizationUnits()
	mspManager,_ := netHelper.Chain.GetMSPManager().GetMSPs()
	t.Log("chainID: ", chainname)
	for _,orderer := range orderers {
		t.Log("orderer: ", orderer.GetURL())
	}
	for _,peer := range peers {
		t.Log("peer: ", peer.GetURL())
	}
	t.Log("primary peer: ", string(ppeers.GetURL()))
	t.Log("organisation unit: ", strings.Join(orgs," "))
	t.Log("mspManager: ", mspManager)
}

func TestClient(t *testing.T) {
	value,_ := netHelper.Client.GetStateStore().GetValue("admin")
	user, _:= netHelper.Client.LoadUserFromStateStore("admin")
	t.Log("value from statstore (admin): ", string(value))
	t.Log("user name: ", user.GetName())
	t.Log("user enrollment cert: ", string(user.GetEnrollmentCertificate()))
	t.Log("user roles: ",  strings.Join(user.GetRoles()," "))
}

func TestQueryInfos(t *testing.T) {
	t.Log("TestQueryInfos ")
	bci, err := netHelper.QueryInfos()
	if err != nil {
		t.Error(err)
	}
	t.Log("bcinfo string: ", bci.String())
}


func TestQueryTransaction(t *testing.T) {
	txID := createTransaction(t)
	time.Sleep(time.Millisecond * 1500)
	processedTransaction, err :=netHelper.QueryTransaction(txID)
	if err != nil {
		t.Error("QueryTransaction return error: ", err)
	}
	t.Log("transaction: ", processedTransaction.TransactionEnvelope.Payload)
}

func TestQueryBlockByNumber(t *testing.T) {
	block, err := netHelper.QueryBlockByNumber("1")
	if err != nil {
		t.Error("QueryBlockByNumber return error: ", err)
	}
	//dis,_ := json.Marshal(block)
	//t.Log("block: ", dis)
	t.Log("block data : ", block.Data.Data)
	t.Log("block metadata : ", block.Metadata.Metadata)
}

func TestQueryBlockByHash(t *testing.T) {
	bci, err := netHelper.QueryInfos()
	if err != nil {
		t.Fatalf("QueryInfo return error: %v", err)
	}

	// Test Query Block by Hash - retrieve current block by hash
	block, err := netHelper.QueryBlockByHash(string(bci.CurrentBlockHash))
	if err != nil {
		t.Fatalf("QueryBlockByHash return error: %v", err)
	}
	t.Log("block: ", block.String())
}

func TestQueryChannels(t *testing.T) {
	channelQueryResponse, err := netHelper.QueryChannels()
	if err != nil {
		t.Fatalf("QueryChannels return error: %v", err)
	}
	for _, channel := range channelQueryResponse.Channels {
		t.Log("Channel: ",channel, "\n")
	}
}

func TestGetInstalledChainCode(t *testing.T) {
	chaincodeQueryResponse, err := netHelper.GetInstalledChainCode()
	if err != nil {
		t.Fatalf("QueryInstalledChaincodes return error: %v", err)
	}

	for _, chaincode := range chaincodeQueryResponse.Chaincodes {
		t.Log("InstalledCC: ",chaincode,"\n")
	}
}

func TestQueryByChainCode(t *testing.T) {
	chaincodeQueryResponse, err := netHelper.QueryByChainCode("lccc")
	if err != nil {
		t.Fatalf("QueryInstantiatedChaincodes return error: %v", err)
	}

	for _, chaincode := range chaincodeQueryResponse {
		t.Log("InstantiatedCC: ",chaincode,"\n")
	}
}



func createTransaction(t *testing.T)string{
	txID, err := consHelper.CreateConsent(configuration.ChainCodeID, APPID1, OWNERID1, CONSUMERID1, DATATYPE1, DATAACCESS1, getStringDateNow(0), getStringDateNow(7))
	if err != nil {
		t.Error("CreateConsent return error: ", err)
	}
	return txID
}