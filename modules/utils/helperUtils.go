package utils

import (
	"net/http"
)

type Helper interface {
	Init(UserCredentials) error
}

func InitHelper (r *http.Request, helper Helper, credentials UserCredentials, authentication bool)  error {
	var err error
	if authentication{
		credentials, err = GetUserCredentials(r)
		if err != nil {
			return err
		}
	}
	err = helper.Init(credentials)
	return err
}
