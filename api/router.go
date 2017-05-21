package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/op/go-logging"
	"github.com/pascallimeux/his/modules/ocms"
	utils "github.com/pascallimeux/his/modules/utils"
)

var log = logging.MustGetLogger("his.api")

const (

	BCINFO           = "/his/v0/dashboard/chain"
	QUERYTRANSACTION = "/his/v0/dashboard/transaction"
	BLOCKBYNB        = "/his/v0/dashboard/blocks/nb"
	BLOCKBYHASH      = "/his/v0/dashboard/blocks/hash"
	GETCHANNELS      = "/his/v0/dashboard/channels"
	INSTALLEDCC      = "/his/v0/dashboard/cc/installed"
	QUERYBYCC        = "/his/v0/dashboard/cc/query"
	GETPEERS         = "/his/v0/dashboard/peers"
	GETORDERS	 = "/his/v0/dashboard/orderers"
	INSTANCIATEDCC   = "/his/v0/dashboard/cc/instanciated"

	REGISTER         = "/his/v0/admin/user/register"
	ENROLL           = "/his/v0/admin/user/enroll"
	REVOKE           = "/his/v0/admin/user/revoke"
	DEPLOYCC         = "/his/v0/admin/deploycc"
	ADDORDERER	 = "/his/v0/admin/orderer"
	ADDPEER		 = "/his/v0/admin/peer"
	CREATECHANNEL	 = "/his/v0/admin/channel"
)

type AppContext struct {
	HttpServer     	*http.Server
	OcmsContext     ocms.OCMSContext
	ChainCodeID   	string
	ChainID         string
	Repo            string
	StatStorePath   string
	Authent         bool
	AdmCrendentials  utils.UserCredentials
}

func (a *AppContext) CreateHISRoutes(router *mux.Router) {
	log.Debug("CreateHISRoutes() : calling method -")
	router.HandleFunc(BCINFO, a.blockchainInfo).Methods("GET")
	router.HandleFunc(GETCHANNELS, a.getChannels).Methods("GET")
	router.HandleFunc(GETPEERS, a.getPeers).Methods("GET")
	router.HandleFunc(GETORDERS, a.getOrderers).Methods("GET")
	router.HandleFunc(INSTALLEDCC, a.getInstalledCC).Methods("GET")
	router.HandleFunc(INSTANCIATEDCC, a.getInstantiatedCC).Methods("GET")
	router.HandleFunc(QUERYTRANSACTION+"/{truuid}", a.transactionDetails).Methods("GET")
	router.HandleFunc(BLOCKBYNB+"/{blocknb}", a.blockByNumber).Methods("GET")
	router.HandleFunc(BLOCKBYHASH+"/{blockhash}", a.blockByHash).Methods("GET")
	router.HandleFunc(QUERYBYCC+"/{ccname}", a.queryByCC).Methods("GET")

	router.HandleFunc(REGISTER, a.registerUser).Methods("POST")
	router.HandleFunc(ENROLL, a.enrollUser).Methods("POST")
	router.HandleFunc(REVOKE, a.revokeUser).Methods("POST")
	router.HandleFunc(DEPLOYCC, a.deployCC).Methods("POST")
	router.HandleFunc(ADDORDERER, a.addOrderer).Methods("POST")
	router.HandleFunc(ADDPEER, a.addPeer).Methods("POST")
	router.HandleFunc(CREATECHANNEL, a.createChannel).Methods("POST")
}
