package utils

import (
	"fmt"
	sdkUtil "github.com/hyperledger/fabric-sdk-go/fabric-client/helpers"
	sdkConfig "github.com/hyperledger/fabric-sdk-go/config"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/hyperledger/fabric-sdk-go/fabric-client/events"
	"errors"
	"strconv"
	"time"
	"math/rand"
)

func GetChain(userCredentials UserCredentials, statStorePath, chainID string) (fabricClient.Chain, error) {
	log.Debug("GetChain(username:"+ userCredentials.UserName+") : calling method -")
	var chain fabricClient.Chain

	client, err := GetClient(userCredentials, statStorePath)
	if err != nil {
		return chain, errors.New("getClient return error: %v" + err.Error())
	}
	chain, err = sdkUtil.GetChain(client, chainID)
	if err != nil {
		log.Error("Create chain ", chainID," failed: ", err)
		return chain, err
	}
	return chain, nil
}

func GetEventHub() (events.EventHub, error) {
	log.Debug("GetEventHub() : calling method -")
	eventHub := events.NewEventHub()
	foundEventHub := false
	peerConfig, err := sdkConfig.GetPeersConfig()
	if err != nil {
		return nil, fmt.Errorf("Error reading peer config: %v", err)
	}
	for _, p := range peerConfig {
		if p.EventHost != "" && p.EventPort != 0 {
			log.Debug("EventHub connect to peer (", p.EventHost,":", p.EventPort,")")
			eventHub.SetPeerAddr(fmt.Sprintf("%s:%d", p.EventHost, p.EventPort),
				p.TLS.Certificate, p.TLS.ServerHostOverride)
			foundEventHub = true
			break
		}
	}
	if !foundEventHub {
		return nil, fmt.Errorf("No EventHub configuration found")
	}
	return eventHub, nil
}


func GetClient(userCredentials UserCredentials, statStorePath string) (fabricClient.Client, error) {
	log.Debug("GetClient(username:"+ userCredentials.UserName+") : calling method -")
	client, err := sdkUtil.GetClient(userCredentials.UserName, userCredentials.Password, statStorePath)
	if err != nil {
		log.Debug("getClient return error: %v" + err.Error())
		return client, errors.New("getClient return error: %v" + err.Error())
	}
	return client, nil
}

func CreateRandomName() string {
	rand.Seed(time.Now().UnixNano())
	return "user" + strconv.Itoa(rand.Intn(500000))
}

func GetRandomString(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}