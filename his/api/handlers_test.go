package api

import (
	"github.com/gorilla/mux"
	"net/http/httptest"
	"github.com/pascallimeux/his/his/settings"
	"github.com/pascallimeux/his/his/helpers"
	utils "github.com/pascallimeux/his/his/modules/utils"
	"time"
	"testing"
	"os"
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/pascallimeux/his/his/modules/ocms"
)

var configuration settings.Settings
var httpServerTest *httptest.Server
const(
	ADMINNAME          = "admin"
	ADMINPWD           = "admpw"
	APPID              = "apptest"
	TransactionTimeout = time.Millisecond * 2500
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	// Init settings
	var err error
	configuration, err = settings.GetSettings("..", "his_test")
	if err != nil {
		panic(err.Error())
	}
	networkHelper := helpers.NewNetworkHelper(configuration.Repo, configuration.StatstorePath, configuration.ChainID)

	adminCredentials := utils.UserCredentials {
		UserName:configuration.Adminusername,
		Password:configuration.AdminPwd}

	err = networkHelper.StartNetwork(adminCredentials, configuration.ProviderName, configuration.NetworkConfigfile, configuration.ChannelConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	err = configuration.InitLogger()
	if err != nil {
		log.Fatal(err)
	}

	networkHelper.DeployCC(helpers.ChainCode{ChainCodePath: configuration.ChainCodePath,ChainCodeVersion: configuration.ChainCodeVersion, ChainCodeID: configuration.ChainCodeID})

	// Init applications context
	ocmsContext := ocms.OCMSContext{
		ChainCodeID: 		configuration.ChainCodeID,
		Repo:                   configuration.Repo,
		StatStorePath:          configuration.StatstorePath,
		ChainID:         	configuration.ChainID,
		Authent:                configuration.AuthMode,
		AdmCrendentials:        adminCredentials,
	}

	appContext := AppContext{
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

	// Init http server for tests
	httpServerTest = httptest.NewServer(router)

}

func shutdown() {
	defer httpServerTest.Close()
	defer configuration.CloseLogger()
}


func buildRequestWithLoginPassword(method, uri, data, login, password string) (*http.Request, error) {
	request, err := buildRequest(method, uri, data)
	if err != nil {
		return request, err
	}
	request.SetBasicAuth(login,password)
	return request, nil
}

func buildRequest(method, uri, data string) (*http.Request, error) {
	var requestData *strings.Reader
	if data != "" {
		requestData = strings.NewReader(data)
	} else {
		requestData = strings.NewReader(" ")
		//requestData = nil
	}
	request, err := http.NewRequest(method, uri, requestData)
	if err != nil {
		return request, err
	}
	return request, nil
}


func executeRequest(request *http.Request) (int, []byte, error) {
	status := 0
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return status, nil, err
	}
	status = response.StatusCode
	body_bytes, err2 := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err2 != nil {
		return status, body_bytes, err2
	}
	return status, body_bytes, nil
}
