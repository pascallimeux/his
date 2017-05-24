package version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
)

// GetVersionOKCode is the HTTP code returned for type GetVersionOK
const GetVersionOKCode int = 200

/*GetVersionOK Version of smart contract consent

swagger:response getVersionOK
*/
type GetVersionOK struct {
	/*
	  Required: true
	*/
	Version string `json:"version"`
}

// NewGetVersionOK creates GetVersionOK with default headers values
func NewGetVersionOK() *GetVersionOK {
	return &GetVersionOK{}
}

// WithVersion adds the version to the get version o k response
func (o *GetVersionOK) WithVersion(version string) *GetVersionOK {
	o.Version = version
	return o
}

// SetVersion sets the version to the get version o k response
func (o *GetVersionOK) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *GetVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.WriteHeader(200)
}

/*GetVersionDefault Genric error for API

swagger:response getVersionDefault
*/
type GetVersionDefault struct {
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

// NewGetVersionDefault creates GetVersionDefault with default headers values
func NewGetVersionDefault(code int) *GetVersionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetVersionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get version default response
func (o *GetVersionDefault) WithStatusCode(code int) *GetVersionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get version default response
func (o *GetVersionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithCode adds the code to the get version default response
func (o *GetVersionDefault) WithCode(code int64) *GetVersionDefault {
	o.Code = code
	return o
}

// SetCode sets the code to the get version default response
func (o *GetVersionDefault) SetCode(code int64) {
	o.Code = code
}

// WithMessage adds the message to the get version default response
func (o *GetVersionDefault) WithMessage(message string) *GetVersionDefault {
	o.Message = message
	return o
}

// SetMessage sets the message to the get version default response
func (o *GetVersionDefault) SetMessage(message string) {
	o.Message = message
}

// WriteResponse to the client
func (o *GetVersionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
