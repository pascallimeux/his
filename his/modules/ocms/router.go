package ocms

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/op/go-logging"
	"github.com/pascallimeux/his/his/modules/utils"
)

var log = logging.MustGetLogger("his.ocms")

const (
	OWNER            = "/owner"
	CONSUMER         = "/consumer"
	CONSENTS         = "/consents"
	ISCONSENT        = "/isconsent"
	API              = "/ocms/v3/api"
	APP              = "/app"
	APIURI           = API+APP+"/{appid}"
	APICONSENTSURI   = APIURI+CONSENTS
	VERSIONURI       = API+"/version"
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
	router.HandleFunc(VERSIONURI, oc.GetVersion).Methods("GET")
	router.HandleFunc(APICONSENTSURI, oc.GetConsents).Methods("GET")
	router.HandleFunc(APICONSENTSURI, oc.CreateConsent).Methods("POST")
	router.HandleFunc(APICONSENTSURI, oc.DeleteConsents).Methods("DELETE")
	router.HandleFunc(APICONSENTSURI+"/{consentid}", oc.GetConsent).Methods("GET")
	router.HandleFunc(APICONSENTSURI+"/{consentid}", oc.DeleteConsent).Methods("DELETE")
	router.HandleFunc(APIURI+ISCONSENT, oc.IsConsent).Methods("POST")
	router.HandleFunc(APIURI+OWNER+   "/{ownerid}"+CONSENTS, oc.GetConsents4Owner).Methods("GET")
	router.HandleFunc(APIURI+CONSUMER+"/{consumerid}"+CONSENTS, oc.GetConsents4Consumer).Methods("GET")
	router.HandleFunc(APIURI+CONSUMER+"/{consumerid}"+OWNER+"/{ownerid}"+CONSENTS, oc.GetConsents4ConsumerOwner).Methods("GET")
}
