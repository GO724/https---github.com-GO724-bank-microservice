// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetCardsParams creates a new GetCardsParams object
//
// There are no default values defined in the spec.
func NewGetCardsParams() GetCardsParams {

	return GetCardsParams{}
}

// GetCardsParams contains all the bound params for the get cards operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetCards
type GetCardsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*maximum number of records to return
	  Maximum: 50
	  Minimum: 0
	  In: query
	*/
	Limit *int32
	/*pass an optional search string for looking up card
	  In: query
	*/
	SearchString *string
	/*number of records to skip for pagination
	  Minimum: 0
	  In: query
	*/
	Skip *int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetCardsParams() beforehand.
func (o *GetCardsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qSearchString, qhkSearchString, _ := qs.GetOK("searchString")
	if err := o.bindSearchString(qSearchString, qhkSearchString, route.Formats); err != nil {
		res = append(res, err)
	}

	qSkip, qhkSkip, _ := qs.GetOK("skip")
	if err := o.bindSkip(qSkip, qhkSkip, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetCardsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int32", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetCardsParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", int64(*o.Limit), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "query", int64(*o.Limit), 50, false); err != nil {
		return err
	}

	return nil
}

// bindSearchString binds and validates parameter SearchString from query.
func (o *GetCardsParams) bindSearchString(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.SearchString = &raw

	return nil
}

// bindSkip binds and validates parameter Skip from query.
func (o *GetCardsParams) bindSkip(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("skip", "query", "int32", raw)
	}
	o.Skip = &value

	if err := o.validateSkip(formats); err != nil {
		return err
	}

	return nil
}

// validateSkip carries on validations for parameter Skip
func (o *GetCardsParams) validateSkip(formats strfmt.Registry) error {

	if err := validate.MinimumInt("skip", "query", int64(*o.Skip), 0, false); err != nil {
		return err
	}

	return nil
}
