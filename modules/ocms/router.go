package ocms

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("his.ocms")


const (
	VERSIONURI       = "/ocms/v3/api/version"
	CONSENTAPI       = "/ocms/v3/api/consent"
)


type OCMSContext struct {
	HttpServer     	*http.Server
	ChainCodeID   	string
	ChainID         string
	Repo            string
	StatStorePath   string
}

func (oc *OCMSContext) CreateOCMSRoutes(router *mux.Router) {
	log.Debug("CreateOCMSRoutes() : calling method -")
	router.HandleFunc(VERSIONURI, oc.getVersion).Methods("GET")
	router.HandleFunc(CONSENTAPI, oc.processConsent).Methods("POST")
}