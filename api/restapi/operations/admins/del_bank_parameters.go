// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDelBankParams creates a new DelBankParams object
//
// There are no default values defined in the spec.
func NewDelBankParams() DelBankParams {

	return DelBankParams{}
}

// DelBankParams contains all the bound params for the del bank operation
// typically these are obtained from a http.Request
//
// swagger:parameters delBank
type DelBankParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*bank to remove
	  Required: true
	  In: path
	*/
	Inn int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDelBankParams() beforehand.
func (o *DelBankParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rInn, rhkInn, _ := route.Params.GetOK("inn")
	if err := o.bindInn(rInn, rhkInn, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindInn binds and validates parameter Inn from path.
func (o *DelBankParams) bindInn(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("inn", "path", "int64", raw)
	}
	o.Inn = value

	return nil
}
