package utils

import (
	"net/http"
)

type Helper interface {
	Init(UserCredentials) error
}

func InitHelper (r *http.Request, helper Helper)  error {
	userCredentials, err := GetUserCredentials(r)
	if err != nil {
		return err
	}
	err = helper.Init(userCredentials)
	return err
}
