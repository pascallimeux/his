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
		utils.SendError(w, err)
		return
	}
	userHelper := &helpers.UserHelper{StatStorePath:a.StatStorePath}
	err = helpers.InitHelper(r, userHelper)
	if err != nil {
		utils.SendError(w, err)
		return
	}
	enrollmentSecret := EnrollmentSecret{}
	enrollmentSecret.Secret, err = userHelper.RegisterUser(user)
	if err != nil {
		utils.SendError(w, err)
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
		utils.SendError(w, err)
		return
	}
	userHelper := &helpers.UserHelper{StatStorePath:a.StatStorePath}
	err = helpers.InitHelper(r, userHelper)
	if err != nil {
		utils.SendError(w, err)
		return
	}
	err = userHelper.EnrollUser(credentials)
	if err != nil {
		utils.SendError(w, err)
		return
	}
	content := []byte("")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

//HTTP Post - /his/v0/admin/user/revoke
func (a *AppContext) revokeUser(w http.ResponseWriter, r *http.Request) {
	log.Debug("enrollUser() : calling method -")
	var credentials utils.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		utils.SendError(w, err)
		return
	}
	userHelper := &helpers.UserHelper{StatStorePath:a.StatStorePath}
	err = helpers.InitHelper(r, userHelper)
	if err != nil {
		utils.SendError(w, err)
		return
	}
	err = userHelper.RevokeUser(credentials)
	if err != nil {
		utils.SendError(w, err)
		return
	}
	content := []byte("")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}