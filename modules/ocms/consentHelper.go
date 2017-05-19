package ocms

import (
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/hyperledger/fabric-sdk-go/fabric-client/events"
	"github.com/pascallimeux/his/modules/utils"
	"fmt"
	"strings"
	"encoding/json"
)

type ConsentHelper struct {
	ChainID         string
	StatStorePath   string
	Chain 	        fabricClient.Chain
	EventHub        events.EventHub
	Initialized	bool
}

type Consent struct {
	Action		string     `json:"action"`
	AppID 		string     `json:"appid"`
	State       	string     `json:"state"`
	ConsentID      	string     `json:"consentid"`
	OwnerID       	string     `json:"ownerid"`
	ConsumerID      string     `json:"consumerid"`
	DataType      	string     `json:"datatype"`
	DataAccess      string     `json:"dataaccess"`
	Dt_begin      	string     `json:"dtbegin"`
	Dt_end       	string     `json:"dtend"`
}


func (ch *ConsentHelper) Init(userCredentials utils.UserCredentials) error{
	chain, err := utils.GetChain(userCredentials, ch.StatStorePath, ch.ChainID)
	if err != nil {
		return err
	}
	eventHub, err := utils.GetEventHub()
	if err != nil {
		return err
	}
	if err := eventHub.Connect(); err != nil {
		return err
	}
	ch.Chain    = chain
	ch.EventHub = eventHub
	ch.Initialized = true
	return nil
}

func (ch *Consent) Print() string {
	consentStr := fmt.Sprintf("ConsentID:%s ConsumerID:%s OwnerID:%s Datatype:%s Dataaccess:%s Dt_begin:%s Dt_end:%s", ch.ConsentID, ch.ConsumerID, ch.OwnerID, ch.DataType, ch.DataAccess, ch.Dt_begin, ch.Dt_end)
	return consentStr
}


func (ch *ConsentHelper) GetVersion(chainCodeID string) (string, error) {
	log.Debug("GetVersion(chainCodeID:"+ chainCodeID+") : calling method -")
	var args []string
	args = append(args, "getversion")
	return utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
}

func (ch *ConsentHelper) GetConsent(chainCodeID, appID, consentID string) (Consent, error) {
	var args []string
	args = append(args, "getconsent")
	args = append(args, appID)
	args = append(args, consentID)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractConsent(consentID, strResp, err)
}

func (ch *ConsentHelper) GetConsents(chainCodeID, appID string) ([]Consent, error) {
	var args []string
	args = append(args, "getconsents")
	args = append(args, appID)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractConsents(strResp, err)
}

func (ch *ConsentHelper) GetOwnerConsents(chainCodeID, appID, ownerID string) ([]Consent, error) {
	var args []string
	args = append(args, "getownerconsents")
	args = append(args, appID)
	args = append(args, ownerID)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractConsents(strResp, err)
}

func (ch *ConsentHelper) GetConsumerConsents(chainCodeID, appID, consumerID string) ([]Consent, error) {
	var args []string
	args = append(args, "getconsumerconsents")
	args = append(args, appID)
	args = append(args, consumerID)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractConsents(strResp, err)
}

func (ch *ConsentHelper) GetConsumerOwnerConsents(chainCodeID, appID, consumerID, ownerID string) ([]Consent, error) {
	var args []string
	args = append(args, "getconsumerownerconsents")
	args = append(args, appID)
	args = append(args, ownerID)
	args = append(args, consumerID)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractConsents(strResp, err)
}

func (ch *ConsentHelper) CreateConsent(chainCodeID, appID, ownerID, consumerID, datatype, dataaccess, st_date, end_date string) (string, error) {
	var args []string
	args = append(args, "postconsent")
	args = append(args, appID)
	args = append(args, ownerID)
	args = append(args, consumerID)
	args = append(args, datatype)
	args = append(args, dataaccess)
	args = append(args, st_date)
	args = append(args, end_date)
	txID, err := utils.CreateTransaction(ch.ChainID, chainCodeID, args, ch.Chain)
	return txID, err
}

func (ch *ConsentHelper) DeleteConsents4Application(chainCodeID, appID string) (string, error) {
	var args []string
	args = append(args, "resetconsents")
	args = append(args, appID)
	txID, err := utils.CreateTransaction(ch.ChainID, chainCodeID, args, ch.Chain)
	return txID, err
}

func (ch *ConsentHelper) RemoveConsent(chainCodeID, appID, consentID string) (string, error) {
	var args []string
	args = append(args, "removeconsent")
	args = append(args, appID)
	args = append(args, consentID)
	return utils.CreateTransaction(ch.ChainID, chainCodeID, args, ch.Chain)
}

func (ch *ConsentHelper) IsConsentExist(chainCodeID, appID, ownerID, consumerID, dataType, dataAccess string) (bool, error) {
	var args []string
	args = append(args, "isconsent")
	args = append(args, appID)
	args = append(args, ownerID)
	args = append(args, consumerID)
	args = append(args, dataType)
	args = append(args, dataAccess)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractIsConsent(strResp, err)
}

func (ch *ConsentHelper) CreateConsentWithRegistration(chainCodeID, appID, ownerID, consumerID, datatype, dataaccess, st_date, end_date string) (string, error) {
	var args []string
	args = append(args, "postconsent")
	args = append(args, appID)
	args = append(args, ownerID)
	args = append(args, consumerID)
	args = append(args, datatype)
	args = append(args, dataaccess)
	args = append(args, st_date)
	args = append(args, end_date)
	return utils.CreateTransactionWithRegistration(ch.ChainID, chainCodeID, args, ch.Chain, ch.EventHub)
}


func extractConsents(stringresp string, err error) ([]Consent, error) {
	var consents []Consent
	if err != nil {
		return consents, err
	}
	dec := json.NewDecoder(strings.NewReader(stringresp))
	err = dec.Decode(&consents)
	if err != nil {
		log.Error(err)
		err = fmt.Errorf("Extract consents return error")
	}
	return consents, err
}

func extractConsent(consentID, stringresp string, err error) (Consent, error) {
	var consent Consent
	if err != nil {
		return consent, err
	}
	dec := json.NewDecoder(strings.NewReader(stringresp))
	err = dec.Decode(&consent)
	if err != nil {
		log.Error(err)
		err = fmt.Errorf("Extract a consent return error")
	}
	return consent, err
}

func extractIsConsent(stringresp string, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	if stringresp == "True" {
		return true, nil
	} else {
		return false, nil
	}
}