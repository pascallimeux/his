package api

import (
	"errors"
	"encoding/json"
	"github.com/pascallimeux/his/his/modules/ocms"
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
	log.Debug("consentID=", consentID)
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
		t.Log(consent.ToString())
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
		t.Log(consent.ToString())
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
		t.Log(consent.ToString())
	}
}

func TestGetConsents4ConsumerOwnerFromAPINominal(t *testing.T) {
	consumerid := "3333"
	ownerid := "1111"
	createConsent(ocms.Consent{OwnerID: ownerid, ConsumerID: consumerid})
	createConsent(ocms.Consent{OwnerID: "2222", ConsumerID: consumerid})
	createConsent(ocms.Consent{OwnerID: ownerid, ConsumerID: "2222"})
	consents, err := getListOfConsents(ownerid, consumerid)
	if err != nil {
		t.Error(err)
	}
	for _, consent := range consents {
		t.Log(consent.ToString())
	}
}


func createConsent(consent ocms.Consent) (string, error) {
	var responseConsent ocms.Consent
	consent.AppID = APPID
	data, _ := json.Marshal(consent)
	request, err := buildRequestWithLoginPassword("POST", httpServerTest.URL+ocms.APICONSENTSURI, string(data), ADMINNAME, ADMINPWD)
	if err != nil {
		return "", err
	}
	status, body_bytes, err := executeRequest(request)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body_bytes, &responseConsent)
	if err != nil {
		return "", err
	}

	if status != http.StatusOK {
		return "", errors.New("bad status")
	}
	return responseConsent.ConsentID, nil
}

func getConsent(consentID string) (ocms.Consent, error) {
	responseConsent := ocms.Consent{}
	request, err1 := buildRequestWithLoginPassword("GET", httpServerTest.URL+ocms.APICONSENTSURI+"/"+APPID+","+consentID, "", ADMINNAME, ADMINPWD)
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
	consents := []ocms.Consent{}
	var request *http.Request
	var err error
	if ownerID != ""&& consumerID !="" {
		request, err = buildRequestWithLoginPassword("GET", httpServerTest.URL+ocms.API+ocms.APP+"/"+APPID+ocms.CONSUMER+"/"+consumerID+ocms.OWNER+"/"+ownerID+ocms.CONSENTS, "", ADMINNAME, ADMINPWD)
	} else if consumerID != "" {
		request, err = buildRequestWithLoginPassword("GET", httpServerTest.URL+ocms.API+ocms.APP+"/"+APPID+ocms.CONSUMER+"/"+consumerID+ocms.CONSENTS, "", ADMINNAME, ADMINPWD)
	} else if ownerID !=""{
		request, err = buildRequestWithLoginPassword("GET", httpServerTest.URL+ocms.API+ocms.APP+"/"+APPID+ocms.OWNER+"/"+ownerID+ocms.CONSENTS, "", ADMINNAME, ADMINPWD)
	}
	if err != nil {
		return consents, err
	}
	status, body_bytes, err := executeRequest(request)
	if err != nil {
		return consents, err
	}
	err = json.Unmarshal(body_bytes, &consents)
	if err != nil {
		return consents, err
	}
	if status != http.StatusOK {
		return consents, errors.New("bad status")
	}
	return consents, nil
}

