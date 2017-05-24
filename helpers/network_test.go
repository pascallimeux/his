package helpers

import (
	"testing"
	"time"
	//"strings"
	sdkConfig "github.com/hyperledger/fabric-sdk-go/config"
	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric/protos/common"
	//hppeer "github.com/hyperledger/fabric/protos/peer"
	protosutils "github.com/hyperledger/fabric/protos/utils"
	//"fmt"
	"encoding/json"
	//"github.com/hyperledger/fabric/protos/orderer"
)

func TestPeerconfig(t *testing.T) {
	peersConfig, _ := sdkConfig.GetPeersConfig()
	tt, _ := json.MarshalIndent(peersConfig, "", "  ")
	t.Log(string(tt))
}


func TestGetPeers(t *testing.T) {
	peers := netHelper.GetPeers()
	for _, peer := range peers {
		t.Log("peer: name=",peer.GetName(), " url=", peer.GetURL(), "\n")
	}
}

func TestGetOrderer(t *testing.T) {
	orderers := netHelper.GetOrderers()
	for _, orderer := range orderers {
		t.Log("orderer: url=", orderer.GetURL(), "\n")
	}
}
/*
func TestMSPManager(t *testing.T) {
	mspmanager := netHelper.Chain.GetMSPManager()
	mspManagers, _ := mspmanager.GetMSPs()
	t.Log("nb msp: ", len(mspManagers))
	for _, msp := range mspManagers {
		t.Log(msp.GetIntermediateCerts())
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
}*/

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
	_, err :=netHelper.QueryTransaction(txID)
	if err != nil {
		t.Error("QueryTransaction return error: ", err)
	}
}

func TestQueryBlockByNumber(t *testing.T) {
	block, err := netHelper.QueryBlockByNumber("1")
	if err != nil {
		t.Error("QueryBlockByNumber return error: ", err)
	}
	//dis,_ := json.Marshal(block)
	t.Log("block header hash: ", block.Header.DataHash)
	t.Log("block data : ", block.Data.Data)
	t.Log("block metadata : ", block.Metadata.Metadata)
	data := &cb.BlockData{}
	if err = proto.Unmarshal(block.Data.Data[0], data); err != nil {
		t.Error("QueryBlockByNumber return error: ", err)
	}
	id, _ :=protosutils.GetChainIDFromBlock(block)

	t.Log("ID:", id)
	t.Log("dataStr: ",data.String())
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