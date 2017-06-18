package utils

import (
	sdkUtil "github.com/hyperledger/fabric-sdk-go/fabric-client/helpers"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/hyperledger/fabric-sdk-go/fabric-client/events"
	"fmt"
	"time"
	"strings"
)

func Query(chainID, chainCodeID string, args []string, chain fabricClient.Chain) (string, error) {
	log.Debug("query(chainCodeID:"+ chainCodeID+" args:"+ strings.Join(args," ") +") : calling method -")
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("TODO change...")
	transactionProposalResponses, _, err := sdkUtil.CreateAndSendTransactionProposal(chain, chainCodeID, chainID, args, []fabricClient.Peer{chain.GetPrimaryPeer()}, transientDataMap)
	if err != nil {
		log.Error("CreateAndSendTransactionProposal return error: %v", err)
		return "", fmt.Errorf("Query CC return error")
	}
	response := string(transactionProposalResponses[0].GetResponsePayload())
	return response, nil
}

func CreateTransaction(chainID, chainCodeID string, args []string, chain fabricClient.Chain) (string, error) {
	log.Debug("createTransaction(chainCodeID:"+ chainCodeID+" args:"+ strings.Join(args," ") +") : calling method -")
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("TODO change...")
	transactionProposalResponse, txID, err := sdkUtil.CreateAndSendTransactionProposal(chain, chainCodeID, chainID, args, []fabricClient.Peer{chain.GetPrimaryPeer()}, transientDataMap)
	if err != nil {
		log.Error("CreateAndSendTransactionProposal return error: %v", err)
		return "", fmt.Errorf("CreateTransactionProposal for CC return error")
	}
	_, err = sdkUtil.CreateAndSendTransaction(chain, transactionProposalResponse)
	if err != nil {
		log.Error("CreateAndSendTransaction return error: %v", err)
		return "", fmt.Errorf("CreateTransaction for CC return error")
	}
	return txID, nil
}

func CreateTransactionWithRegistration(chainID, chainCodeID string, args []string, chain fabricClient.Chain, eventHub events.EventHub) (string, error) {
	log.Debug("createTransactionWithRegistration(chainCodeID:"+ chainCodeID+" args:"+ strings.Join(args," ") +") : calling method -")
	eventID := "test([a-zA-Z]+)"
	// Register callback for chaincode event
	done1, rce := sdkUtil.RegisterCCEvent(chainCodeID, eventID, eventHub)
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("TODO change...")
	transactionProposalResponse, txID, err := sdkUtil.CreateAndSendTransactionProposal(chain, chainCodeID, chainID, args, []fabricClient.Peer{chain.GetPrimaryPeer()}, transientDataMap)
	if err != nil {
		return "", fmt.Errorf("CreateAndSendTransactionProposal return error: %v", err)
	}
	// Register for commit event
	done, fail := sdkUtil.RegisterTxEvent(txID, eventHub)

	_, err = sdkUtil.CreateAndSendTransaction(chain, transactionProposalResponse)
	if err != nil {
		return "", fmt.Errorf("CreateAndSendTransaction return error: %v", err)
	}

	select {
	case <-done:
	case <-fail:
		return txID, fmt.Errorf("invoke Error received from eventhub for txid(%s) error(%v)", txID, fail)
	case <-time.After(time.Second * 30):
		return txID, fmt.Errorf("invoke Didn't receive block event for txid(%s)", txID)
	}

	select {
	case <-done1:
	case <-time.After(time.Second * 20):
		return txID, fmt.Errorf("Did NOT receive CC for eventId(%s)\n", eventID)
	}
	eventHub.UnregisterChaincodeEvent(rce)

	return txID, nil
}

