package ocms

import (
	"net/http"
	"encoding/json"
	"time"
	"fmt"
	"errors"
	utils "github.com/pascallimeux/his/modules/utils"
)
var badRequest =errors.New("Bad request")

type Version struct {
	Version string
}

type IsConsent struct {
	Consent string
}

//HTTP Get - /his/v0/api/version
func (a *OCMSContext) getVersion(w http.ResponseWriter, r *http.Request) {
	log.Debug("getVersion() : calling method -")
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	err := utils.InitHelper(r, consentHelper)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	version := Version{}
	version.Version, err = consentHelper.GetVersion(a.ChainCodeID)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
	}
	content, _ := json.Marshal(version)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Post - /his/v0/api/consent
func (a *OCMSContext) processConsent(w http.ResponseWriter, r *http.Request) {
	log.Debug("processConsent() : calling method -")

	var bytes []byte
	var consent Consent
	err := json.NewDecoder(r.Body).Decode(&consent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	err = utils.InitHelper(r, consentHelper)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	switch action := consent.Action; action {
	case "create":
		bytes, err = a.createConsent(consentHelper, a.ChainCodeID, consent)
	case "list":
		bytes, err = a.listConsents(consentHelper, a.ChainCodeID, consent.AppID)
	case "get":
		bytes, err = a.getConsent(consentHelper, a.ChainCodeID, consent.AppID, consent.ConsentID)
	case "remove":
		bytes, err = a.unactivateConsent(consentHelper, a.ChainCodeID, consent.AppID, consent.ConsentID)
	case "list4owner":
		bytes, err = a.getConsents4Owner(consentHelper, a.ChainCodeID, consent.AppID, consent.OwnerID)
	case "list4consumer":
		bytes, err = a.getConsents4Consumer(consentHelper, a.ChainCodeID, consent.AppID, consent.ConsumerID)
	case "isconsent":
		bytes, err = a.isConsent(consentHelper, a.ChainCodeID, consent)
	default:
		log.Error("bad action request: ", action)
		utils.SendError(w, errors.New("Bad action request"))
		return
	}
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}


func (a *OCMSContext) createConsent(consentHelper *ConsentHelper, chainCodeID string, consent Consent) ([]byte, error) {
	log.Debug("createConsent(applicationID:"+ consent.AppID+") : calling method -")
	err := check_args(&consent)
	var message string
	if err != nil {
		message = fmt.Sprintf("createConsent(%s) : calling method -", err.Error())
	} else {
		message = fmt.Sprintf("createConsent(%s) : calling method -", consent.Print())
	}
	log.Info(message)
	if err != nil {
		return nil, err
	}
	consentID, err := consentHelper.CreateConsent(chainCodeID, consent.AppID, consent.OwnerID, consent.ConsumerID, consent.DataType, consent.DataAccess, consent.Dt_begin, consent.Dt_end)
	if err != nil {
		return nil, err
	}
	consent.ConsentID = consentID
	return consent2Bytes(consent)
}

func (a *OCMSContext) listConsents(consentHelper *ConsentHelper, chainCodeID, applicationID string) ([]byte, error) {
	message := fmt.Sprintf("listConsents(applicationID=%s) : calling method -", applicationID)
	log.Info(message)
	consents, err := consentHelper.GetConsents(chainCodeID, applicationID)
	if err != nil {
		return nil, err
	}
	return consents2Bytes(consents)
}

func (a *OCMSContext) getConsent(consentHelper *ConsentHelper, chainCodeID, applicationID, consentID string) ([]byte, error) {
	message := fmt.Sprintf("getConsent(applicationID=%s, consentID=%s) : calling method -", applicationID, consentID)
	log.Info(message)
	consent, err := consentHelper.GetConsent(chainCodeID, applicationID, consentID)
	if err != nil {
		return nil, err
	}
	return consent2Bytes(consent)
}

func (a *OCMSContext) unactivateConsent(consentHelper *ConsentHelper, chainCodeID, applicationID, consentID string) ([]byte, error) {
	message := fmt.Sprintf("unactivateConsent(applicationID=%s, consentID=%s) : calling method -", applicationID, consentID)
	log.Info(message)
	_, err := consentHelper.RemoveConsent(chainCodeID, applicationID, consentID)
	if err != nil {
		return nil, err
	}
	consent, err := consentHelper.GetConsent(chainCodeID, applicationID, consentID)
	if err != nil {
		return nil, err
	}
	return consent2Bytes(consent)
}

func (a *OCMSContext) getConsents4Consumer(consentHelper *ConsentHelper, chainCodeID, applicationID, consumerID string) ([]byte, error) {
	message := fmt.Sprintf("getConsents4Consumer(applicationID=%s, consumerID=%s) : calling method -", applicationID, consumerID)
	log.Info(message)
	consents, err := consentHelper.GetConsumerConsents(chainCodeID, applicationID, consumerID)
	if err != nil {
		return nil, err
	}
	return consents2Bytes(consents)
}

func (a *OCMSContext) getConsents4Owner(consentHelper *ConsentHelper, chainCodeID, applicationID, ownerID string) ([]byte, error) {
	message := fmt.Sprintf("getConsents4Owner(applicationID=%s, ownerID=%s) : calling method -", applicationID, ownerID)
	log.Info(message)
	consents, err := consentHelper.GetOwnerConsents(chainCodeID, applicationID, ownerID)
	if err != nil {
		return nil, err
	}
	return consents2Bytes(consents)
}

func (a *OCMSContext) isConsent(consentHelper *ConsentHelper, chainCodeID string, consent Consent) ([]byte, error) {
	message := fmt.Sprintf("isConsent(consent=%s) : calling method -", consent.Print())
	log.Info(message)
	isconsent, err := consentHelper.IsConsentExist(chainCodeID, consent.AppID, consent.OwnerID, consent.ConsumerID, consent.DataType, consent.DataAccess)
	if err != nil {
		return nil, err
	}
	response := IsConsent{}
	if isconsent {
		response.Consent = "True"
	} else {
		response.Consent = "False"
	}
	content, _ := json.Marshal(response)
	return content, nil
}

func consents2Bytes(consents []Consent) ([]byte, error) {
	log.Debug("consents2Bytes() : calling method -")
	j, err := json.Marshal(consents)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func consent2Bytes(consent Consent) ([]byte, error) {
	log.Debug("consent2Bytes() : calling method -")
	j, err := json.Marshal(consent)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func check_args(consent *Consent) error {
	log.Debug("check_args() : calling method -")
	if consent.AppID == "" {
		return errors.New("appID is mandatory!")
	}
	if consent.OwnerID == "" {
		return errors.New("ownerID is mandatory!")
	}
	if consent.ConsumerID == "" {
		return errors.New("consumerID is mandatory!")
	}
	if consent.DataAccess == "" {
		consent.DataAccess = "A"
	}
	if consent.DataType == "" {
		consent.DataType = "All"
	}
	if consent.Dt_begin == "" {
		consent.Dt_begin = time.Now().Format("2006-01-02")
	}
	if consent.Dt_end == "" {
		consent.Dt_end = "2099-01-01"
	}
	return nil
}
