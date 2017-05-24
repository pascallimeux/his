package api

import (
	"net/http"
	"encoding/json"
	"github.com/pascallimeux/his/helpers"
	"github.com/pascallimeux/his/modules/utils"
)


type EnrollmentSecret struct {
	Secret string
}

//HTTP Post - /his/v0/admin/user/register
func (a *AppContext) registerUser(w http.ResponseWriter, r *http.Request) {
	log.Debug("registerUser() : calling method -")
	var user helpers.UserRegistrer
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorUserCredentials, -1)
		return
	}
	userHelper := helpers.NewUserHelper(a.StatStorePath)
	err = utils.InitHelper(r, userHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	enrollmentSecret := EnrollmentSecret{}
	enrollmentSecret.Secret, err = userHelper.RegisterUser(user)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorUserCredentials, -1)
	}
	content, _ := json.Marshal(enrollmentSecret)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)

}

//HTTP Post - /his/v0/admin/user/enroll
func (a *AppContext) enrollUser(w http.ResponseWriter, r *http.Request) {
	log.Debug("enrollUser() : calling method -")
	var credentials utils.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorUserCredentials, -1)
		return
	}
	userHelper := helpers.NewUserHelper(a.StatStorePath)
	err = utils.InitHelper(r, userHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	err = userHelper.EnrollUser(credentials)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorUserCredentials, -1)
		return
	}
	content := []byte(`{"Content":"User enrolled"}`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Post - /his/v0/admin/user/revoke
func (a *AppContext) revokeUser(w http.ResponseWriter, r *http.Request) {
	log.Debug("revokeUser() : calling method -")
	var credentials utils.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorUserCredentials, -1)
		return
	}
	userHelper := helpers.NewUserHelper(a.StatStorePath)
	err = utils.InitHelper(r, userHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	err = userHelper.RevokeUser(credentials)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorUserCredentials, -1)
		return
	}
	content := []byte(`{"Content":"User revoked"}`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Post - /his/v0/admin/deploycc
func (a *AppContext) deployCC(w http.ResponseWriter, r *http.Request) {
	log.Debug("deployCC() : calling method -")
	var chaincode helpers.ChainCode
	err := json.NewDecoder(r.Body).Decode(&chaincode)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorUserCredentials, -1)
		return
	}
	netHelper := helpers.NewNetworkHelper(a.Repo, a.StatStorePath, a.ChainID)
	err = utils.InitHelper(r, netHelper, a.AdmCrendentials, a.Authent)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorInializeHelper, -1)
		return
	}
	err = netHelper.DeployCC(chaincode)
	if err != nil {
		log.Error(err)
		utils.SendError(w, utils.ErrorDeployCC, -1)
		return
	}
	content := []byte(`{"Content":"Chaincode deployed"}`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Post - /his/v0/admin/orderer
func (a *AppContext) addOrderer(w http.ResponseWriter, r *http.Request) {
	log.Debug("addOrderer() : calling method -")
	//TODO
}


//HTTP Post - /his/v0/admin/peer
func (a *AppContext) addPeer(w http.ResponseWriter, r *http.Request) {
	log.Debug("addPeer() : calling method -")
	//TODO
}

//HTTP Post - /his/v0/admin/channel
func (a *AppContext) createChannel(w http.ResponseWriter, r *http.Request) {
	log.Debug("createChannel() : calling method -")
	//TODO
}

//HTTP Post - /his/v0/admin/orderer/delete
func (a *AppContext) removeOrderer(w http.ResponseWriter, r *http.Request) {
	log.Debug("removeOrderer() : calling method -")
	//TODO
}


//HTTP Post - /his/v0/admin/peer/remove
func (a *AppContext) removePeer(w http.ResponseWriter, r *http.Request) {
	log.Debug("removePeer() : calling method -")
	//TODO
}