package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetConsents4ConsumerParams creates a new GetConsents4ConsumerParams object
// with the default values initialized.
func NewGetConsents4ConsumerParams() GetConsents4ConsumerParams {
	var ()
	return GetConsents4ConsumerParams{}
}

// GetConsents4ConsumerParams contains all the bound params for the get consents4 consumer operation
// typically these are obtained from a http.Request
//
// swagger:parameters getConsents4Consumer
type GetConsents4ConsumerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The application ID to submit.
	  Required: true
	  In: path
	*/
	Appid string
	/*The consumer ID to submit.
	  Required: true
	  In: path
	*/
	Consumerid string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetConsents4ConsumerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rAppid, rhkAppid, _ := route.Params.GetOK("appid")
	if err := o.bindAppid(rAppid, rhkAppid, route.Formats); err != nil {
		res = append(res, err)
	}

	rConsumerid, rhkConsumerid, _ := route.Params.GetOK("consumerid")
	if err := o.bindConsumerid(rConsumerid, rhkConsumerid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConsents4ConsumerParams) bindAppid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Appid = raw

	return nil
}

func (o *GetConsents4ConsumerParams) bindConsumerid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Consumerid = raw

	return nil
}