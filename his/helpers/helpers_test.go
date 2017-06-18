package helpers

import (
	"time"
	"testing"
	"os"
	"github.com/pascallimeux/his/his/settings"
	"github.com/pascallimeux/his/his/modules/utils"
	"github.com/pascallimeux/his/his/modules/ocms"
)

const(
	CCVERSION   	   = "Orange Consent Application chaincode ver 3 Dated 2017-03-09"
	APPID1      	   = "APP4TESTS1"
	APPID2     	   = "APP4TESTS2"
	APPID3     	   = "APP4TESTS3"
	APPID4     	   = "APP4TESTS4"
	APPID5     	   = "APP4TESTS5"
	APPID6     	   = "APP4TESTS6"
	OWNERID1   	   = "owner1"
	OWNERID2   	   = "owner2"
	OWNERID3   	   = "owner3"
	CONSUMERID1	   = "consumer1"
	CONSUMERID2	   = "consumer2"
	CONSUMERID3	   = "consumer3"
	DATATYPE1  	   = "type1"
	DATAACCESS1	   = "access1"
	TransactionTimeout = time.Millisecond * 1500
)

var netHelper  NetworkHelper
var consHelper ocms.ConsentHelper
var userhelper UserHelper
var configuration settings.Settings
var statStorePath string

func setup() {

	var err error
	// Init settings
	configuration, err = settings.GetSettings("..", "his_test")
	if err != nil {
		panic(err.Error())
	}
	statStorePath =  configuration.StatstorePath
	adminCredentials := utils.UserCredentials {UserName:configuration.Adminusername, Password:configuration.AdminPwd}

	// Init network helper
	netHelper = NewNetworkHelper(configuration.Repo, configuration.StatstorePath, configuration.ChainID)
	err = netHelper.StartNetwork(adminCredentials, configuration.ProviderName, configuration.NetworkConfigfile, configuration.ChannelConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	err = configuration.InitLogger()
	if err != nil {
		log.Fatal(err)
	}
	err = netHelper.Init(adminCredentials)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Init user helper
	userhelper = NewUserHelper(configuration.StatstorePath)
	userhelper.Init(adminCredentials)

	// Init consent helper
	consHelper = ocms.NewConsentHelper(configuration.ChainID, configuration.StatstorePath)
	err = consHelper.Init(adminCredentials)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Deploy the smartcontract
	netHelper.DeployCC(ChainCode{ChainCodePath: configuration.ChainCodePath,ChainCodeVersion: configuration.ChainCodeVersion, ChainCodeID: configuration.ChainCodeID})
}

func shutdown(){
	defer configuration.CloseLogger()
}

func TestMain(m *testing.M) {
	setup()
	time.Sleep(time.Millisecond * 3000)
	code := m.Run()
	shutdown()
	os.Exit(code)
}


func getStringDateNow(nbdaysafter time.Duration) string{
	t := time.Now().Add(nbdaysafter * 24 * time.Hour)
	return t.Format("2006-01-02")
}

