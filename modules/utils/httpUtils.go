package utils

import (
	"errors"
	"net/http"
	"github.com/op/go-logging"
	"strings"
)

var log = logging.MustGetLogger("his.utils")

type UserCredentials struct {
	UserName 	 string	   `json:"username"`
	Password	 string	   `json:"password"`
}

func GetUserCredentials(r *http.Request)(UserCredentials, error){
	userCredentials := UserCredentials{}
	username, password, ok :=r.BasicAuth()
	if ok{
		log.Debug("GetUserCredentials(user:" + username + ") : calling method -")
		userCredentials.UserName = username
		userCredentials.Password = password
		return userCredentials, nil
	}
	return userCredentials, errors.New("no credential in request")
}

func SendError(w http.ResponseWriter, err error) {
	log.Debug("sendError() : calling method -")
	libelle := err.Error()
	libelle = strings.Replace(libelle, "\"", "'", -1)
	log.Error("send http error: ", libelle)
	message := "{\"content\":\"" + libelle + "\"} "
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
