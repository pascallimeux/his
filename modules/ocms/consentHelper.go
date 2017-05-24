package ocms

import (
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/hyperledger/fabric-sdk-go/fabric-client/events"
	"github.com/pascallimeux/his/modules/utils"
	"fmt"
	"strings"
	"encoding/json"
)

type consentHelper struct {
	ChainID         string
	StatStorePath   string
	Chain 	        fabricClient.Chain
	EventHub        events.EventHub
	Initialized	bool
}


func NewConsentHelper(chainID, statStorePath string) ConsentHelper {
	c := &consentHelper{ChainID: chainID, StatStorePath: statStorePath}
	return c
}

func (ch *consentHelper) Init(userCredentials utils.UserCredentials) error{
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

type ConsentHelper interface {
	Init(utils.UserCredentials) error
	GetVersion(chainCodeID string) (string, error)
	GetConsent(chainCodeID, appID, consentID string) (Consent, error)
	GetConsents(chainCodeID, appID string) ([]Consent, error)
	GetOwnerConsents(chainCodeID, appID, ownerID string) ([]Consent, error)
	GetConsumerConsents(chainCodeID, appID, consumerID string) ([]Consent, error)
	GetConsumerOwnerConsents(chainCodeID, appID, consumerID, ownerID string) ([]Consent, error)
	CreateConsent(chainCodeID, appID, ownerID, consumerID, datatype, dataaccess, st_date, end_date string) (string, error)
	DeleteConsents4Application(chainCodeID, appID string) (string, error)
	DeleteConsent(chainCodeID, appID, consentID string) (string, error)
	IsConsentExist(chainCodeID, appID, ownerID, consumerID, dataType, dataAccess string) (bool, error)
	CreateConsentWithRegistration(chainCodeID, appID, ownerID, consumerID, datatype, dataaccess, st_date, end_date string) (string, error)
}


func (ch *consentHelper) GetVersion(chainCodeID string) (string, error) {
	log.Debug("GetVersion(chainCodeID:"+ chainCodeID+") : calling method -")
	var args []string
	args = append(args, "getversion")
	return utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
}

func (ch *consentHelper) GetConsent(chainCodeID, appID, consentID string) (Consent, error) {
	var args []string
	args = append(args, "getconsent")
	args = append(args, appID)
	args = append(args, consentID)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractConsent(strResp, err)
}

func (ch *consentHelper) GetConsents(chainCodeID, appID string) ([]Consent, error) {
	var args []string
	args = append(args, "getconsents")
	args = append(args, appID)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractConsents(strResp, err)
}

func (ch *consentHelper) GetOwnerConsents(chainCodeID, appID, ownerID string) ([]Consent, error) {
	var args []string
	args = append(args, "getownerconsents")
	args = append(args, appID)
	args = append(args, ownerID)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractConsents(strResp, err)
}

func (ch *consentHelper) GetConsumerConsents(chainCodeID, appID, consumerID string) ([]Consent, error) {
	var args []string
	args = append(args, "getconsumerconsents")
	args = append(args, appID)
	args = append(args, consumerID)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractConsents(strResp, err)
}

func (ch *consentHelper) GetConsumerOwnerConsents(chainCodeID, appID, consumerID, ownerID string) ([]Consent, error) {
	var args []string
	args = append(args, "getconsumerownerconsents")
	args = append(args, appID)
	args = append(args, consumerID)
	args = append(args, ownerID)
	strResp, err := utils.Query(ch.ChainID, chainCodeID, args, ch.Chain)
	return extractConsents(strResp, err)
}

func (ch *consentHelper) CreateConsent(chainCodeID, appID, ownerID, consumerID, datatype, dataaccess, st_date, end_date string) (string, error) {
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

func (ch *consentHelper) DeleteConsents4Application(chainCodeID, appID string) (string, error) {
	var args []string
	args = append(args, "deleteconsents")
	args = append(args, appID)
	txID, err := utils.CreateTransaction(ch.ChainID, chainCodeID, args, ch.Chain)
	return txID, err
}

func (ch *consentHelper) DeleteConsent(chainCodeID, appID, consentID string) (string, error) {
	var args []string
	args = append(args, "deleteconsent")
	args = append(args, appID)
	args = append(args, consentID)
	txID, err := utils.CreateTransaction(ch.ChainID, chainCodeID, args, ch.Chain)
	return txID, err
}

func (ch *consentHelper) IsConsentExist(chainCodeID, appID, ownerID, consumerID, dataType, dataAccess string) (bool, error) {
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

func (ch *consentHelper) CreateConsentWithRegistration(chainCodeID, appID, ownerID, consumerID, datatype, dataaccess, st_date, end_date string) (string, error) {
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

func extractConsent(stringresp string, err error) (Consent, error) {
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