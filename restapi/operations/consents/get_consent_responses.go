package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/pascallimeux/his/models"
)

// GetConsentOKCode is the HTTP code returned for type GetConsentOK
const GetConsentOKCode int = 200

/*GetConsentOK consent

swagger:response getConsentOK
*/
type GetConsentOK struct {

	/*
	  In: Body
	*/
	Payload *models.Consent `json:"body,omitempty"`
}

// NewGetConsentOK creates GetConsentOK with default headers values
func NewGetConsentOK() *GetConsentOK {
	return &GetConsentOK{}
}

// WithPayload adds the payload to the get consent o k response
func (o *GetConsentOK) WithPayload(payload *models.Consent) *GetConsentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get consent o k response
func (o *GetConsentOK) SetPayload(payload *models.Consent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConsentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetConsentDefault Genric error for API

swagger:response getConsentDefault
*/
type GetConsentDefault struct {
	_statusCode int
	/*in: path
	  Required: true
	*/
	Code int64 `json:"code"`
	/*
	  Required: true
	*/
	Message string `json:"message"`
}

// NewGetConsentDefault creates GetConsentDefault with default headers values
func NewGetConsentDefault(code int) *GetConsentDefault {
	if code <= 0 {
		code = 500
	}

	return &GetConsentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get consent default response
func (o *GetConsentDefault) WithStatusCode(code int) *GetConsentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get consent default response
func (o *GetConsentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithCode adds the code to the get consent default response
func (o *GetConsentDefault) WithCode(code int64) *GetConsentDefault {
	o.Code = code
	return o
}

// SetCode sets the code to the get consent default response
func (o *GetConsentDefault) SetCode(code int64) {
	o.Code = code
}

// WithMessage adds the message to the get consent default response
func (o *GetConsentDefault) WithMessage(message string) *GetConsentDefault {
	o.Message = message
	return o
}

// SetMessage sets the message to the get consent default response
func (o *GetConsentDefault) SetMessage(message string) {
	o.Message = message
}

// WriteResponse to the client
func (o *GetConsentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header code

	code := swag.FormatInt64(o.Code)
	if code != "" {
		rw.Header().Set("code", code)
	}

	// response header message

	message := o.Message
	if message != "" {
		rw.Header().Set("message", message)
	}

	rw.WriteHeader(o._statusCode)
}