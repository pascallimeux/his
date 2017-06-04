package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetConsentParams creates a new GetConsentParams object
// with the default values initialized.
func NewGetConsentParams() GetConsentParams {
	var ()
	return GetConsentParams{}
}

// GetConsentParams contains all the bound params for the get consent operation
// typically these are obtained from a http.Request
//
// swagger:parameters getConsent
type GetConsentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The application ID to submit.
	  Required: true
	  In: path
	*/
	Appid string
	/*The consent ID to submit.
	  Required: true
	  In: path
	*/
	Consentid string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetConsentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rAppid, rhkAppid, _ := route.Params.GetOK("appid")
	if err := o.bindAppid(rAppid, rhkAppid, route.Formats); err != nil {
		res = append(res, err)
	}

	rConsentid, rhkConsentid, _ := route.Params.GetOK("consentid")
	if err := o.bindConsentid(rConsentid, rhkConsentid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConsentParams) bindAppid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Appid = raw

	return nil
}

func (o *GetConsentParams) bindConsentid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Consentid = raw

	return nil
}