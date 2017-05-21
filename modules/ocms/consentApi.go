package ocms

import (
	"net/http"
	"encoding/json"
	"time"
	"fmt"
	"errors"
	"github.com/pascallimeux/his/modules/utils"
	"github.com/gorilla/mux"
)

// version response
// swagger:response versionResponse
type VersionResponse struct {
	Version string `json:"version"`
}

type IsConsentResponse struct {
	Consent string `json:"isconsent"`
}

type ConsentStatusResponse struct {
	Status string `json:"consentstatus"`
}

// getVersion swagger:route GET /ocms/v3/api/version orders getVersion
//
// Gets the version of the chaincode consent.
//
// Responses:
//    default: genericError
//        200: versionResponse
func (a *OCMSContext) getVersion(w http.ResponseWriter, r *http.Request) {
	log.Debug("getVersion() : calling method -")
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	var err error
	err = utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	version := VersionResponse{}
	version.Version, err = consentHelper.GetVersion(a.ChainCodeID)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
	}
	utils.SendStruct(version, w)
}

//HTTP Post - /his/v0/api/consent
func (a *OCMSContext) createConsent(w http.ResponseWriter, r *http.Request) {
	log.Debug("createConsent() : calling method -")
	var consent Consent
	err := json.NewDecoder(r.Body).Decode(&consent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	err = utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	err = check_args(&consent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorBadArgsRequest, -1)
		return
	}
	consentID, err := consentHelper.CreateConsent(a.ChainCodeID, consent.AppID, consent.OwnerID, consent.ConsumerID, consent.DataType, consent.DataAccess, consent.Dt_begin, consent.Dt_end)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
		return
	}
	consent.ConsentID = consentID
	utils.SendStruct(consent, w)
}

//HTTP Get - /his/v0/api/consents/{appid}
func (a *OCMSContext) listConsents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	message := fmt.Sprintf("listConsents(appid=%s) : calling method -", appid)
	log.Debug(message)
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	err := utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	consents, err := consentHelper.GetConsents(a.ChainCodeID, appid)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
		return
	}
	utils.SendStruct(consents, w)
}

//HTTP Get - /his/v0/api/consent/{appid, consentid}
func (a *OCMSContext) getConsent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	consentid := vars["consentid"]
	message := fmt.Sprintf("getConsent(appid=%s, consentid=%s) : calling method -", appid, consentid)
	log.Debug(message)
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	err := utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	consent, err := consentHelper.GetConsent(a.ChainCodeID, appid, consentid)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
		return
	}
	utils.SendStruct(consent, w)
}


//HTTP Delete - /his/v0/api/consent/{appid, consentid}
func (a *OCMSContext) deleteConsent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	consentid := vars["consentid"]
	message := fmt.Sprintf("deleteConsent(appid=%s, consentid=%s) : calling method -", appid, consentid)
	log.Debug(message)
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	err := utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	_, err = consentHelper.RemoveConsent(a.ChainCodeID, appid, consentid)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
		return
	}
	response := ConsentStatusResponse{}
	response.Status = "Inactivated"
	utils.SendStruct(response, w)
}

//HTTP Post - /his/v0/api/isconsent
func (a *OCMSContext) isConsent(w http.ResponseWriter, r *http.Request) {
	log.Debug("isConsent() : calling method -")
	var consent Consent
	err := json.NewDecoder(r.Body).Decode(&consent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorBadArgsRequest, -1)
		return
	}
	message := fmt.Sprintf("isConsent(consent=%s) : calling method -", consent.Print())
	log.Info(message)
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	err = utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	isconsent, err := consentHelper.IsConsentExist(a.ChainCodeID, consent.AppID, consent.OwnerID, consent.ConsumerID, consent.DataType, consent.DataAccess)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
		return
	}
	response := IsConsentResponse{}
	if isconsent {
		response.Consent = "True"
	} else {
		response.Consent = "False"
	}
	utils.SendStruct(response, w)
}

//HTTP Get - /his/v0/api/consent/owner/{appid, ownerid}
func (a *OCMSContext) getConsents4Owner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	ownerid := vars["ownerid"]
	message := fmt.Sprintf("getConsents4Owner(appid=%s, ownerid=%s) : calling method -", appid, ownerid)
	log.Debug(message)
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	err := utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	consents, err := consentHelper.GetOwnerConsents(a.ChainCodeID, appid, ownerid)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
		return
	}
	utils.SendStruct(consents, w)
}

//HTTP Get - /his/v0/api/consent/consumer/{appid, consumerid}
func (a *OCMSContext) getConsents4Consumer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	consumerid := vars["consumerid"]
	message := fmt.Sprintf("getConsents4Consumer(appid=%s, consumerid=%s) : calling method -", appid, consumerid)
	log.Debug(message)
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	err := utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	consents, err := consentHelper.GetConsumerConsents(a.ChainCodeID, appid, consumerid)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
		return
	}
	utils.SendStruct(consents, w)
}


//HTTP Get - /his/v0/api/consent/consumerowner/{appid, consumerid, ownerid}
func (a *OCMSContext) getConsents4ConsumerOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	consumerid := vars["consumerid"]
	ownerid := vars["ownerid"]
	message := fmt.Sprintf("getConsents4ConsumerOwner(appid=%s, consumerid=%s, ownerid=%s) : calling method -", appid, consumerid, ownerid)
	log.Debug(message)
	consentHelper := &ConsentHelper{ChainID:a.ChainID, StatStorePath:a.StatStorePath}
	err := utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	consents, err := consentHelper.GetConsumerOwnerConsents(a.ChainCodeID, appid, consumerid, ownerid)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
		return
	}
	utils.SendStruct(consents, w)
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
