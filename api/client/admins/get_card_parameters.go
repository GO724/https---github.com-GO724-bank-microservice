// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetCardParams creates a new GetCardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCardParams() *GetCardParams {
	return &GetCardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCardParamsWithTimeout creates a new GetCardParams object
// with the ability to set a timeout on a request.
func NewGetCardParamsWithTimeout(timeout time.Duration) *GetCardParams {
	return &GetCardParams{
		timeout: timeout,
	}
}

// NewGetCardParamsWithContext creates a new GetCardParams object
// with the ability to set a context for a request.
func NewGetCardParamsWithContext(ctx context.Context) *GetCardParams {
	return &GetCardParams{
		Context: ctx,
	}
}

// NewGetCardParamsWithHTTPClient creates a new GetCardParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCardParamsWithHTTPClient(client *http.Client) *GetCardParams {
	return &GetCardParams{
		HTTPClient: client,
	}
}

/*
GetCardParams contains all the parameters to send to the API endpoint

	for the get card operation.

	Typically these are written to a http.Request.
*/
type GetCardParams struct {

	/* Inn.

	   get card
	*/
	Inn int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get card params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCardParams) WithDefaults() *GetCardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get card params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get card params
func (o *GetCardParams) WithTimeout(timeout time.Duration) *GetCardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get card params
func (o *GetCardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get card params
func (o *GetCardParams) WithContext(ctx context.Context) *GetCardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get card params
func (o *GetCardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get card params
func (o *GetCardParams) WithHTTPClient(client *http.Client) *GetCardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get card params
func (o *GetCardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInn adds the inn to the get card params
func (o *GetCardParams) WithInn(inn int64) *GetCardParams {
	o.SetInn(inn)
	return o
}

// SetInn adds the inn to the get card params
func (o *GetCardParams) SetInn(inn int64) {
	o.Inn = inn
}

// WriteToRequest writes these params to a swagger request
func (o *GetCardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inn
	if err := r.SetPathParam("inn", swag.FormatInt64(o.Inn)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}