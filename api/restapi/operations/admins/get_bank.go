// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBankHandlerFunc turns a function with the right signature into a get bank handler
type GetBankHandlerFunc func(GetBankParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBankHandlerFunc) Handle(params GetBankParams) middleware.Responder {
	return fn(params)
}

// GetBankHandler interface for that can handle valid get bank params
type GetBankHandler interface {
	Handle(GetBankParams) middleware.Responder
}

// NewGetBank creates a new http.Handler for the get bank operation
func NewGetBank(ctx *middleware.Context, handler GetBankHandler) *GetBank {
	return &GetBank{Context: ctx, Handler: handler}
}

/*
	GetBank swagger:route GET /banks/{inn} admins getBank

get an bank item

Gets an bank from the system
*/
type GetBank struct {
	Context *middleware.Context
	Handler GetBankHandler
}

func (o *GetBank) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBankParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}