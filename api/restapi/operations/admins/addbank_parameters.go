// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"bank-microservice/api/models"
)

// NewAddbankParams creates a new AddbankParams object
//
// There are no default values defined in the spec.
func NewAddbankParams() AddbankParams {

	return AddbankParams{}
}

// AddbankParams contains all the bound params for the addbank operation
// typically these are obtained from a http.Request
//
// swagger:parameters addbank
type AddbankParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*bank to add
	  In: body
	*/
	Bank *models.Bank
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddbankParams() beforehand.
func (o *AddbankParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Bank
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("bank", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Bank = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
