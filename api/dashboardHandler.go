package api
import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pascallimeux/his/helpers"
	"github.com/pascallimeux/his/modules/utils"
	"github.com/pkg/errors"
)

var badRequest =errors.New("Bad request")

//HTTP Get - /his/v0/dashboard/chain
func (a *AppContext) blockchainInfo(w http.ResponseWriter, r *http.Request) {
	log.Debug("blockchainInfo() : calling method -")
	var err error
	netHelper := &helpers.NetworkHelper{Repo: a.Repo, StatStorePath: a.StatStorePath, ChainID: a.ChainID}
	err = utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	blockchainInfo, err := netHelper.QueryInfos()
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
	}
	content, err := json.Marshal(blockchainInfo)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Get - /his/v0/dashboard/channels
func (a *AppContext) getChannels(w http.ResponseWriter, r *http.Request) {
	log.Debug("getChannels() : calling method -")
	var err error
	netHelper := &helpers.NetworkHelper{Repo: a.Repo, StatStorePath: a.StatStorePath, ChainID: a.ChainID}
	err = utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	channels, err := netHelper.QueryChannels()
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
	}
	content, err := json.Marshal(channels)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Get - /his/v0/dashboard/peers
func (a *AppContext) getPeers(w http.ResponseWriter, r *http.Request) {
	log.Debug("getPeers() : calling method -")
	var err error
	netHelper := &helpers.NetworkHelper{Repo: a.Repo, StatStorePath: a.StatStorePath, ChainID: a.ChainID}
	err = utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	peers :=netHelper.GetPeers()
	content, err := json.Marshal(peers)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Get - /his/v0/dashboard/orderers
func (a *AppContext) getOrderers(w http.ResponseWriter, r *http.Request) {
	log.Debug("getOrderers() : calling method -")
	var err error
	netHelper := &helpers.NetworkHelper{Repo: a.Repo, StatStorePath: a.StatStorePath, ChainID: a.ChainID}
	err = utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	orderers :=netHelper.GetOrderers()
	content, err := json.Marshal(orderers)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Get - /his/v0/dashboard/cc/installed
func (a *AppContext) getInstalledCC(w http.ResponseWriter, r *http.Request) {
	log.Debug("getInstalledCC() : calling method -")
	var err error
	netHelper := &helpers.NetworkHelper{Repo: a.Repo, StatStorePath: a.StatStorePath, ChainID: a.ChainID}
	err = utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	cc, err := netHelper.GetInstalledChainCode()
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
	}
	content, err := json.Marshal(cc)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Get - /his/v0/dashboard/cc/instanciated
func (a *AppContext) getInstantiatedCC(w http.ResponseWriter, r *http.Request) {
	log.Debug("getInstantiatedCC() : calling method -")
	var err error
	netHelper := &helpers.NetworkHelper{Repo: a.Repo, StatStorePath: a.StatStorePath, ChainID: a.ChainID}
	err = utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	cc, err := netHelper.GetInstanciateChainCode()
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
	}
	content, err := json.Marshal(cc)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Get - /his/v0/dashboard/transaction/{truuid}
func (a *AppContext) transactionDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tr_uuid := vars["truuid"]
	message := fmt.Sprintf("transactionDetails(tr_uuid=%s) : calling method -", tr_uuid)
	log.Debug(message)
	netHelper := &helpers.NetworkHelper{Repo: a.Repo, StatStorePath: a.StatStorePath, ChainID: a.ChainID}
	err := utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	transaction, err := netHelper.QueryTransaction(tr_uuid)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	content, err := json.Marshal(transaction)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Get - /his/v0/dashboard/blocks/nb/{blocknb}
func (a *AppContext) blockByNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blocNb := vars["blocknb"]
	message := fmt.Sprintf("blockByNumber(blocknb=%s) : calling method -", blocNb)
	log.Debug(message)

	netHelper := &helpers.NetworkHelper{Repo: a.Repo, StatStorePath: a.StatStorePath, ChainID: a.ChainID}
	err := utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	block, err := netHelper.QueryBlockByNumber(blocNb)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	content, err := json.Marshal(block)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}


//HTTP Get - /his/v0/dashboard/blocks/hash/{blockhash}
func (a *AppContext) blockByHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockHash := vars["blockhash"]
	message := fmt.Sprintf("blockByHash(blockHash=%s) : calling method -", blockHash)
	log.Debug(message)
	netHelper := &helpers.NetworkHelper{Repo: a.Repo, StatStorePath: a.StatStorePath, ChainID: a.ChainID}
	err := utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	block, err := netHelper.QueryBlockByHash(blockHash)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	content, err := json.Marshal(block)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}


//HTTP Get - /his/v0/dashboard/blocks/hash/{ccname}
func (a *AppContext) queryByCC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chaincodeName := vars["ccname"]
	message := fmt.Sprintf("queryByCC(blockHash=%s) : calling method -", chaincodeName)
	log.Debug(message)
	netHelper := &helpers.NetworkHelper{Repo: a.Repo, StatStorePath: a.StatStorePath, ChainID: a.ChainID}
	err := utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	response, err := netHelper.QueryByChainCode(chaincodeName)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	content, err := json.Marshal(response)
	if err != nil {
		log.Error(err)
		utils.SendError(w, badRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}