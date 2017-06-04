package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/pascallimeux/his/models"
)

// CreateConsentOKCode is the HTTP code returned for type CreateConsentOK
const CreateConsentOKCode int = 200

/*CreateConsentOK consent

swagger:response createConsentOK
*/
type CreateConsentOK struct {

	/*
	  In: Body
	*/
	Payload *models.Consent `json:"body,omitempty"`
}

// NewCreateConsentOK creates CreateConsentOK with default headers values
func NewCreateConsentOK() *CreateConsentOK {
	return &CreateConsentOK{}
}

// WithPayload adds the payload to the create consent o k response
func (o *CreateConsentOK) WithPayload(payload *models.Consent) *CreateConsentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create consent o k response
func (o *CreateConsentOK) SetPayload(payload *models.Consent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConsentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateConsentDefault Genric error for API

swagger:response createConsentDefault
*/
type CreateConsentDefault struct {
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

// NewCreateConsentDefault creates CreateConsentDefault with default headers values
func NewCreateConsentDefault(code int) *CreateConsentDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateConsentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create consent default response
func (o *CreateConsentDefault) WithStatusCode(code int) *CreateConsentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create consent default response
func (o *CreateConsentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithCode adds the code to the create consent default response
func (o *CreateConsentDefault) WithCode(code int64) *CreateConsentDefault {
	o.Code = code
	return o
}

// SetCode sets the code to the create consent default response
func (o *CreateConsentDefault) SetCode(code int64) {
	o.Code = code
}

// WithMessage adds the message to the create consent default response
func (o *CreateConsentDefault) WithMessage(message string) *CreateConsentDefault {
	o.Message = message
	return o
}

// SetMessage sets the message to the create consent default response
func (o *CreateConsentDefault) SetMessage(message string) {
	o.Message = message
}

// WriteResponse to the client
func (o *CreateConsentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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