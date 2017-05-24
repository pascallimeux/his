package main

import(
	"github.com/pascallimeux/his/helpers"
	"github.com/pascallimeux/his/api"
	"github.com/pascallimeux/his/settings"
	"net/http"
	"time"
	"github.com/pascallimeux/his/modules/ocms"
	"github.com/op/go-logging"
	"github.com/gorilla/mux"
	"github.com/pascallimeux/his/modules/utils"
)

var log = logging.MustGetLogger("his")

//go:generate swagger generate spec
func main() {
	// Init settings
	configuration, err := settings.GetSettings(".", "his")
	if err != nil {
		panic(err.Error())
	}
	adminCredentials := utils.UserCredentials {
		UserName:configuration.Adminusername,
		Password:configuration.AdminPwd}

	// Init Hyperledger network
	networkHelper := helpers.NewNetworkHelper(configuration.Repo, configuration.StatstorePath, configuration.ChainID)

	err = networkHelper.StartNetwork(adminCredentials, configuration.ProviderName, configuration.NetworkConfigfile, configuration.ChannelConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	err = configuration.InitLogger()
	if err != nil {
		log.Fatal(err)
	}

	// Deploy the consent smartcontract if is not deployed
	networkHelper.DeployCC(helpers.ChainCode{
		ChainCodePath: configuration.ChainCodePath,
		ChainCodeVersion: configuration.ChainCodeVersion,
		ChainCodeID: configuration.ChainCodeID})

	// Init applications context
	ocmsContext := ocms.OCMSContext{
		ChainCodeID: 		configuration.ChainCodeID,
		Repo:                   configuration.Repo,
		StatStorePath:          configuration.StatstorePath,
		ChainID:         	configuration.ChainID,
		Authent:                configuration.AuthMode,
		AdmCrendentials:        adminCredentials,
	}

	appContext := api.AppContext{
		OcmsContext:            ocmsContext,
		ChainCodeID: 		configuration.ChainCodeID,
		Repo:                   configuration.Repo,
		StatStorePath:          configuration.StatstorePath,
		ChainID:         	configuration.ChainID,
		Authent:                configuration.AuthMode,
		AdmCrendentials:        adminCredentials,
	}

	// Init routes for application
	router := mux.NewRouter().StrictSlash(false)
	// Set HIS routes
	appContext.CreateHISRoutes(router)
	// Set OCMS routes
	ocmsContext.CreateOCMSRoutes(router)

	s := &http.Server{
		Addr:         configuration.HttpHostUrl,
		Handler:      router,
		ReadTimeout:  configuration.ReadTimeout * time.Nanosecond,
		WriteTimeout: configuration.WriteTimeout * time.Nanosecond,
	}
	if configuration.Tls {
		log.Debug("Start https Server")
		log.Fatal(s.ListenAndServeTLS("server.crt", "server.key"))
	}else{
		log.Debug("Start http Server")
		log.Fatal(s.ListenAndServe().Error())
	}


	defer configuration.CloseLogger()

}
