package api

import (
	"errors"
	"encoding/json"
	"github.com/pascallimeux/his/modules/ocms"
	"net/http"
	"testing"
	"time"
)

func TestCreateConsentFromAPINominal(t *testing.T) {
	consent := ocms.Consent{OwnerID: "1111", ConsumerID: "2222"}
	consentID, err := createConsent(consent)
	if err != nil {
		t.Error(err)
	}
	if consentID == "" {
		t.Error("bad consent ID")
	}
}

func TestGetConsentDetailFromAPINominal(t *testing.T) {
	consent := ocms.Consent{OwnerID: "OOOO", ConsumerID: "AAAA"}
	consentID, err := createConsent(consent)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(TransactionTimeout)
	consent2, err2 := getConsent(consentID)
	if err2 != nil {
		t.Error(err2)
	}
	if consent2.ConsentID != consentID || consent2.ConsumerID != consent.ConsumerID {
		t.Error(err)
	}

}

func TestGetConsentsFromAPINominal(t *testing.T) {
	createConsent(ocms.Consent{OwnerID: "1111", ConsumerID: "2222"})
	createConsent(ocms.Consent{OwnerID: "1111", ConsumerID: "3333"})
	createConsent(ocms.Consent{OwnerID: "1111", ConsumerID: "4444"})
	consents, err := getListOfConsents("", "")
	if err != nil {
		t.Error(err)
	}
	for _, consent := range consents {
		t.Log(consent.Print())
	}
}

func TestGetConsents4OwnerFromAPINominal(t *testing.T) {
	ownerid := "1111"
	createConsent(ocms.Consent{OwnerID: "1111", ConsumerID: "2222"})
	createConsent(ocms.Consent{OwnerID: "1111", ConsumerID: "3333"})
	createConsent(ocms.Consent{OwnerID: "1111", ConsumerID: "4444"})
	consents, err := getListOfConsents(ownerid, "")
	if err != nil {
		t.Error(err)
	}
	for _, consent := range consents {
		t.Log(consent.Print())
	}
}

func TestGetConsents4ConsumerFromAPINominal(t *testing.T) {
	consumerid := "3333"
	createConsent(ocms.Consent{OwnerID: "1111", ConsumerID: consumerid})
	createConsent(ocms.Consent{OwnerID: "2222", ConsumerID: consumerid})
	createConsent(ocms.Consent{OwnerID: "3333", ConsumerID: consumerid})
	consents, err := getListOfConsents("", consumerid)
	if err != nil {
		t.Error(err)
	}
	for _, consent := range consents {
		t.Log(consent.Print())
	}
}



func createConsent(consent ocms.Consent) (string, error) {
	var responseConsent ocms.Consent
	consent.Action = "create"
	consent.AppID = APPID
	data, _ := json.Marshal(consent)
	request, err1 := buildRequestWithLoginPassword("POST", httpServerTest.URL+ocms.CONSENTAPI, string(data), ADMINNAME, ADMINPWD)
	if err1 != nil {
		return "", err1
	}
	status, body_bytes, err2 := executeRequest(request)
	if err2 != nil {
		return "", err2
	}
	err3 := json.Unmarshal(body_bytes, &responseConsent)
	if err3 != nil {
		return "", err3
	}

	if status != http.StatusOK {
		return "", errors.New("bad status")
	}
	return responseConsent.ConsentID, nil
}

func getConsent(consentID string) (ocms.Consent, error) {
	consent := ocms.Consent{Action: "get", AppID: APPID, ConsentID: consentID}
	responseConsent := ocms.Consent{}
	data, _ := json.Marshal(consent)
	request, err1 := buildRequestWithLoginPassword("POST", httpServerTest.URL+ocms.CONSENTAPI, string(data), ADMINNAME, ADMINPWD)
	if err1 != nil {
		return responseConsent, err1
	}
	status, body_bytes, err2 := executeRequest(request)
	if err2 != nil {
		return responseConsent, err2
	}
	err3 := json.Unmarshal(body_bytes, &responseConsent)
	if err3 != nil {
		return responseConsent, err3
	}
	if status != http.StatusOK {
		return responseConsent, errors.New("bad status")
	}
	return responseConsent, nil
}

func getListOfConsents(ownerID, consumerID string) ([]ocms.Consent, error) {
	consent := ocms.Consent{Action: "list", AppID: APPID}
	consents := []ocms.Consent{}
	if ownerID != "" {
		consent.OwnerID = ownerID
		consent.Action = "list4owner"
	} else if consumerID != "" {
		consent.ConsumerID = consumerID
		consent.Action = "list4consumer"
	}
	data, _ := json.Marshal(consent)
	request, err1 := buildRequestWithLoginPassword("POST", httpServerTest.URL+ocms.CONSENTAPI, string(data), ADMINNAME, ADMINPWD)
	if err1 != nil {
		return consents, err1
	}
	status, body_bytes, err2 := executeRequest(request)
	if err2 != nil {
		return consents, err2
	}
	err3 := json.Unmarshal(body_bytes, &consents)
	if err3 != nil {
		return consents, err3
	}
	if status != http.StatusOK {
		return consents, errors.New("bad status")
	}
	return consents, nil
}

