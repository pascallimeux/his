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

// Version of smart contract consent
// swagger:response versionResponse
type VersionResponse struct {
	Version string `json:"version"`
}
// is the consent exist
// swagger:response isConsentResponse
type IsConsentResponse struct {
	Consent string `json:"isconsent"`
}

// no content
// swagger:response nocontent
type NoContent struct {
	NoContent string `json:"nocontent"`
}

// A ConsentBodyParams model.
//
// This is used for operations that want an Order as body of the request
// swagger:parameters createConsent
type ConsentBodyParams struct {
	// The consent to submit.
	//
	// in: body
	// required: true
	Consent *Consent `json:"consent"`
}

// A AppID parameter model.
//
// This is used for operations that want a ID of an application in the path
// swagger:parameters createConsent getConsents getConsent isConsent getConsents4Owner getConsents4Consumer getConsents4ConsumerOwner deleteConsent deleteConsents
type AppID struct {
	// The application ID to submit.
	//
	// in: path
	// required: true
	AppID string `json:"appid"`
}

// A ConsumerID paramter model.
//
// This is used for operations that want a ID of a consumer in the path
// swagger:parameters getConsents4Consumer getConsents4ConsumerOwner
type ConsumerID struct {
	// The consumer ID to submit.
	//
	// in: path
	// required: true
	ConsumerID string `json:"consumerid"`
}

// A OwnerID parameter model.
//
// This is used for operations that want a ID of a owner in the path
// swagger:parameters getConsents4Owner getConsents4ConsumerOwner
type OwnerID struct {
	// The owner ID to submit.
	//
	// in: path
	// required: true
	OwnerIP string `json:"ownerid"`
}

// A ConsentID parameter model.
//
// This is used for operations that want a ID of a consent in the path
// swagger:parameters   getConsent deleteConsent
type ConsentID struct {
	// The consent ID to submit.
	//
	// in: path
	// required: true
	ConsentID string `json:"consentid"`
}

// getVersion swagger:route GET /ocms/v3/api/version version getVersion
//
// Responses:
//    default: genericError
//        200: versionResponse
func (a *OCMSContext) GetVersion(w http.ResponseWriter, r *http.Request) {
	log.Debug("GetVersion() : calling method -")
	consentHelper := NewConsentHelper(a.ChainID, a.StatStorePath)
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

// CreateConsent swagger:route POST /his/v0/api/app/{appid}/consents consents createConsent
//
// Responses:
//    default: genericError
//        200: consent
func (a *OCMSContext) CreateConsent(w http.ResponseWriter, r *http.Request) {
	log.Debug("CreateConsent() : calling method -")
	var consent Consent
	err := json.NewDecoder(r.Body).Decode(&consent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	vars := mux.Vars(r)
	appid := vars["appid"]
	consent.AppID = appid
	consentHelper := NewConsentHelper(a.ChainID, a.StatStorePath)
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

// DeleteConsents swagger:route DELETE /his/v0/api/app/{appid}/consents consents deleteConsents
//
// Responses:
//    default: genericError
//        204: nocontent
func (a *OCMSContext) DeleteConsents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	message := fmt.Sprintf("DeleteConsents(appid=%s) : calling method -", appid)
	log.Debug(message)
	consentHelper := NewConsentHelper(a.ChainID, a.StatStorePath)
	err := utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	_, err = consentHelper.DeleteConsents4Application(a.ChainCodeID, appid)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}


// GetConsents swagger:route GET /his/v0/api/app/{appid}/consents consents getConsents
//
// Responses:
//    default: genericError
// 	  200: []consent
func (a *OCMSContext) GetConsents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	message := fmt.Sprintf("GetConsents(appid=%s) : calling method -", appid)
	log.Debug(message)
	consentHelper := NewConsentHelper(a.ChainID, a.StatStorePath)
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

// getConsent swagger:route GET /his/v0/api/app/{appid}/consents/{consentid} consents getConsent
//
// Responses:
//    default: genericError
//        200: consent
func (a *OCMSContext) GetConsent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	consentid := vars["consentid"]
	message := fmt.Sprintf("GetConsent(appid=%s, consentid=%s) : calling method -", appid, consentid)
	log.Debug(message)
	consentHelper := NewConsentHelper(a.ChainID, a.StatStorePath)
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

// DeleteConsent swagger:route DELETE /his/v0/api/app/{appid}/consents/{consentid} consents deleteConsent
//
// Responses:
//    default: genericError
//        204: nocontent
func (a *OCMSContext) DeleteConsent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	consentid := vars["consentid"]
	message := fmt.Sprintf("DeleteConsent(appid=%s, consentid=%s) : calling method -", appid, consentid)
	log.Debug(message)
	consentHelper := NewConsentHelper(a.ChainID, a.StatStorePath)
	err := utils.InitHelper(r, consentHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	_, err = consentHelper.DeleteConsent(a.ChainCodeID, appid, consentid)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorChainCode, -1)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// IsConsent swagger:route POST /his/v0/api/app/{appid}/isconsent consents isConsent
//
// Responses:
//    default: genericError
//        200: isConsentResponse
//HTTP Post - /his/v0/api/app/{appid}/isconsent
func (a *OCMSContext) IsConsent(w http.ResponseWriter, r *http.Request) {
	log.Debug("IsConsent() : calling method -")
	var consent Consent
	err := json.NewDecoder(r.Body).Decode(&consent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorBadArgsRequest, -1)
		return
	}
	vars := mux.Vars(r)
	appid := vars["appid"]
	consent.AppID = appid
	message := fmt.Sprintf("isConsent(consent=%s) : calling method -", consent.ToString())
	log.Info(message)
	consentHelper := NewConsentHelper(a.ChainID, a.StatStorePath)
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

// GetConsents4Owner swagger:route GET /his/v0/api/app/{appid}/owner/{ownerid}/consents consents getConsents4Owner
//
// Responses:
//    default: genericError
// 	  200: []consent
func (a *OCMSContext) GetConsents4Owner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	ownerid := vars["ownerid"]
	message := fmt.Sprintf("GetConsents4Owner(appid=%s, ownerid=%s) : calling method -", appid, ownerid)
	log.Debug(message)
	consentHelper := NewConsentHelper(a.ChainID, a.StatStorePath)
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

// GetConsents4Consumer swagger:route GET /his/v0/api/app/{appid}/consumer/{consumerid}/consents consents getConsents4Consumer
//
// Responses:
//    default: genericError
// 	  200: []consent
func (a *OCMSContext) GetConsents4Consumer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	consumerid := vars["consumerid"]
	message := fmt.Sprintf("GetConsents4Consumer(appid=%s, consumerid=%s) : calling method -", appid, consumerid)
	log.Debug(message)
	consentHelper := NewConsentHelper(a.ChainID, a.StatStorePath)
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


// GetConsents4ConsumerOwner swagger:route GET /his/v0/api/app/{appid}/consumer/{consumerid}/owner/{ownerid}/consents consents getConsents4ConsumerOwner
//
// Responses:
//    default: genericError
// 	  200: []consent
func (a *OCMSContext) GetConsents4ConsumerOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appid := vars["appid"]
	consumerid := vars["consumerid"]
	ownerid := vars["ownerid"]
	message := fmt.Sprintf("GetConsents4ConsumerOwner(appid=%s, consumerid=%s, ownerid=%s) : calling method -", appid, consumerid, ownerid)
	log.Debug(message)
	consentHelper := NewConsentHelper(a.ChainID, a.StatStorePath)
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
