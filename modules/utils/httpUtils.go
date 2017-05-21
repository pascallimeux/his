package utils

import (
	"errors"
	"net/http"
	"github.com/op/go-logging"
	"strings"
	"encoding/json"
)

var ErrorBadArgsRequest  = errors.New("Bad arguments request")
var ErrorUserCredentials = errors.New("User credentials error")
var ErrorInializeHelper  = errors.New("Initialze helper error")
var ErrorDeployCC        = errors.New("Chaincode deploy error")
var ErrorHyperledger     = errors.New("Hyperledger error")
var ErrorChainCode       = errors.New("ChainCode error")
var ErrorMarshalStruct   = errors.New("return structure error")

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

func SendErrorOld(w http.ResponseWriter, err error) {
	log.Debug("sendError() : calling method -")
	libelle := err.Error()
	libelle = strings.Replace(libelle, "\"", "'", -1)
	log.Error("send http error: ", libelle)
	message := "{\"content\":\"" + libelle + "\"} "
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func SendStruct(str interface{}, w http.ResponseWriter) {
	log.Debug("SendStruct() : calling method -")
	bytes, err := json.Marshal(str)
	if err != nil {
		log.Error(err)
		SendError(w, ErrorMarshalStruct, -1)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func SendError(w http.ResponseWriter, err error, code int) {
	log.Debug("sendError() : calling method -")
	genericError := &GenericError{
			Code: code,
			Message: err,
		}

	errorResponse, _ := json.Marshal(genericError)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(errorResponse)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: path
	Code    int   `json:"code"`
     	Message error `json:"message"`
}