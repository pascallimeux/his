package ocms

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/op/go-logging"
	"github.com/pascallimeux/his/modules/utils"
)

var log = logging.MustGetLogger("his.ocms")


const (
	VERSIONURI       = "/ocms/v3/api/version"
	CONSENT          = "/ocms/v3/api/consent"
	CONSENTS         = "/ocms/v3/api/consents"
	ISCONSENT        = "/ocms/v3/api/isconsent"
	CONSENTSOWNER    = "/ocms/v3/api/ownerconsents"
	CONSENTSCONSUMER = "/ocms/v3/api/consumerconsents"
	CONSENTSCONSOWN  = "/ocms/v3/api/consumerownerconsents"
)


type OCMSContext struct {
	HttpServer     	*http.Server
	ChainCodeID   	string
	ChainID         string
	Repo            string
	StatStorePath   string
	Authent         bool
	AdmCrendentials  utils.UserCredentials
}

func (oc *OCMSContext) CreateOCMSRoutes(router *mux.Router) {
	log.Debug("CreateOCMSRoutes() : calling method -")
	router.HandleFunc(VERSIONURI, oc.getVersion).Methods("GET")
	router.HandleFunc(CONSENT, oc.createConsent).Methods("POST")
	router.HandleFunc(CONSENT+"/{appid, consentid}", oc.getConsent).Methods("GET")
	router.HandleFunc(CONSENT+"/{appid, consentid}", oc.deleteConsent).Methods("DELETE")
	router.HandleFunc(CONSENTS+"/{appid}", oc.listConsents).Methods("GET")
	router.HandleFunc(ISCONSENT, oc.isConsent).Methods("POST")
	router.HandleFunc(CONSENTSOWNER+"/{appid, ownerid}", oc.getConsents4Owner).Methods("GET")
	router.HandleFunc(CONSENTSCONSUMER+"/{appid, consumerid}", oc.getConsents4Consumer).Methods("GET")
	router.HandleFunc(CONSENTSCONSOWN+"/{appid, consumerid, ownerid}", oc.getConsents4ConsumerOwner).Methods("GET")
}
