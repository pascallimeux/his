package api

import (
	"github.com/pascallimeux/his/helpers"
	utils "github.com/pascallimeux/his/modules/utils"
	"testing"
	"errors"
	"net/http"
	"encoding/json"
)

func TestRegisterUserAPINominal(t *testing.T) {
	username := utils.CreateRandomName()
	registerUser := helpers.UserRegistrer{Name: username, Type: "user", Affiliation: "org1.department1" }
	enrollmentSecret, err := sendRegister(registerUser)
	if err != nil {
		t.Error(err)
	}
	if enrollmentSecret == "" {
		t.Error("bad enrollmentSecret")
	}
}


func TestEnrollUserAPINominal(t *testing.T) {
	username := utils.CreateRandomName()
	registerUser := helpers.UserRegistrer{Name: username, Type: "user", Affiliation: "org1.department1" }
	enrollmentSecret, err := sendRegister(registerUser)
	if err != nil {
		t.Error(err)
	}
	if enrollmentSecret == "" {
		t.Error("bad enrollmentSecret")
	}
	userCredentials := utils.UserCredentials{UserName: username, Password: enrollmentSecret}
	err = sendEnrollUser(userCredentials)
	if err != nil {
		t.Error(err)
	}
}


func TestRevokeUserAPINominal(t *testing.T) {
	username := utils.CreateRandomName()
	registerUser := helpers.UserRegistrer{Name: username, Type: "user", Affiliation: "org1.department1" }
	enrollmentSecret, err := sendRegister(registerUser)
	if err != nil {
		t.Error(err)
	}
	if enrollmentSecret == "" {
		t.Error("bad enrollmentSecret")
	}
	userCredentials := utils.UserCredentials{UserName: username, Password: enrollmentSecret}
	err = sendEnrollUser(userCredentials)
	if err != nil {
		t.Error(err)
	}
	err = sendRevokeUser(userCredentials)
	if err != nil {
		t.Error(err)
	}
}

func sendRevokeUser(userCredentials utils.UserCredentials) error {
	data, _ := json.Marshal(userCredentials)
	request, err := buildRequestWithLoginPassword("POST", httpServerTest.URL+REVOKE, string(data), ADMINNAME, ADMINPWD)
	if err != nil {
		return err
	}
	status, _, err := executeRequest(request)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		return errors.New("bad status")
	}
	return nil
}

func sendEnrollUser(userCredentials utils.UserCredentials) error {
	data, _ := json.Marshal(userCredentials)
	request, err := buildRequestWithLoginPassword("POST", httpServerTest.URL+ENROLL, string(data), ADMINNAME, ADMINPWD)
	if err != nil {
		return err
	}
	status, _, err := executeRequest(request)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		return errors.New("bad status")
	}
	return nil
}

func sendRegister(registerUser helpers.UserRegistrer) (string, error) {
	var response EnrollmentSecret
	data, _ := json.Marshal(registerUser)
	request, err1 := buildRequestWithLoginPassword("POST", httpServerTest.URL+REGISTER, string(data), ADMINNAME, ADMINPWD)
	if err1 != nil {
		return "", err1
	}
	status, body_bytes, err2 := executeRequest(request)
	if err2 != nil {
		return "", err2
	}
	err3 := json.Unmarshal(body_bytes, &response)
	if err3 != nil {
		return "", err3
	}
	if status != http.StatusOK {
		return "", errors.New("bad status")
	}
	return response.Secret, nil
}
