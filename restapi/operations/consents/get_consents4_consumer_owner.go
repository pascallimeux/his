package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetConsents4ConsumerOwnerHandlerFunc turns a function with the right signature into a get consents4 consumer owner handler
type GetConsents4ConsumerOwnerHandlerFunc func(GetConsents4ConsumerOwnerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConsents4ConsumerOwnerHandlerFunc) Handle(params GetConsents4ConsumerOwnerParams) middleware.Responder {
	return fn(params)
}

// GetConsents4ConsumerOwnerHandler interface for that can handle valid get consents4 consumer owner params
type GetConsents4ConsumerOwnerHandler interface {
	Handle(GetConsents4ConsumerOwnerParams) middleware.Responder
}

// NewGetConsents4ConsumerOwner creates a new http.Handler for the get consents4 consumer owner operation
func NewGetConsents4ConsumerOwner(ctx *middleware.Context, handler GetConsents4ConsumerOwnerHandler) *GetConsents4ConsumerOwner {
	return &GetConsents4ConsumerOwner{Context: ctx, Handler: handler}
}

/*GetConsents4ConsumerOwner swagger:route GET /his/v0/api/app/{appid}/consumer/{consumerid}/owner/{ownerid}/consents consents getConsents4ConsumerOwner

GetConsents4ConsumerOwner get consents4 consumer owner API

*/
type GetConsents4ConsumerOwner struct {
	Context *middleware.Context
	Handler GetConsents4ConsumerOwnerHandler
}

func (o *GetConsents4ConsumerOwner) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetConsents4ConsumerOwnerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}