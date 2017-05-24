package helpers

import (
	"testing"
	"github.com/pascallimeux/his/modules/utils"
)

func TestRegisterUser(t *testing.T) {
	username := utils.CreateRandomName()
	registerUser := UserRegistrer{Name: username, Type: "user", Affiliation: "org1.department1"}
	enrollSecret, err :=userhelper.RegisterUser(registerUser)
	if err != nil {
		t.Error(err)
	}
	if enrollSecret == ""{
		t.Error("no enrollSecret received")
	}
}

func TestEnrollUser(t *testing.T) {
	username := utils.CreateRandomName()
	registerUser := UserRegistrer{Name: username, Type: "user", Affiliation: "org1.department1"}
	enrollSecret, err :=userhelper.RegisterUser(registerUser)
	if err != nil {
		t.Error(err)
	}
	userCredentials := utils.UserCredentials{UserName: username, Password: enrollSecret}
	err = userhelper.EnrollUser(userCredentials)
	if err != nil {
		t.Error(err)
	}
}

func TestRevokeUser(t *testing.T) {
	username := utils.CreateRandomName()
	registerUser := UserRegistrer{Name: username, Type: "user", Affiliation: "org1.department1"}
	enrollSecret, err :=userhelper.RegisterUser(registerUser)
	if err != nil {
		t.Error(err)
	}
	userCredentials := utils.UserCredentials{UserName: username, Password: enrollSecret}
	err = userhelper.EnrollUser(userCredentials)
	if err != nil {
		t.Error(err)
	}
	err = userhelper.RevokeUser(userCredentials)
	if err != nil {
		t.Error(err)
	}
}

func TestGetUser(t *testing.T) {
	username := utils.CreateRandomName()
	registerUser := UserRegistrer{Name: username, Type: "user", Affiliation: "org1.department1"}
	enrollSecret, err :=userhelper.RegisterUser(registerUser)
	if err != nil {
		t.Error(err)
	}
	userCredentials := utils.UserCredentials{UserName: username, Password: enrollSecret}
	user, err := userhelper.GetUser(userCredentials)
	if err != nil {
		t.Error(err)
	}
	t.Log(user.GetName())
}


func TestGetClient(t *testing.T) {
	username := utils.CreateRandomName()
	registerUser := UserRegistrer{Name: username, Type: "user", Affiliation: "org1.department1"}
	enrollSecret, err :=userhelper.RegisterUser(registerUser)
	if err != nil {
		t.Error(err)
	}
	userCredentials := utils.UserCredentials{UserName: username, Password: enrollSecret}
	_, err = utils.GetClient(userCredentials, statStorePath)
	if err != nil {
		t.Error(err)
	}
}

func __TestReenrollUser(t *testing.T) {
	username := utils.CreateRandomName()
	registerUser := UserRegistrer{Name: username, Type: "user", Affiliation: "org1.department1"}
	enrollSecret, err :=userhelper.RegisterUser(registerUser)
	if err != nil {
		t.Error(err)
	}
	t.Log("enrollSecret string: ",enrollSecret)
	userCredentials := utils.UserCredentials{UserName: username, Password: enrollSecret}
	err = userhelper.EnrollUser(userCredentials)
	if err != nil {
		t.Error(err)
	}
	err = userhelper.RevokeUser(userCredentials)
	if err != nil {
		t.Error(err)
	}
	err = userhelper.EnrollUser(userCredentials)
	if err != nil {
		t.Error(err)
	}
}
