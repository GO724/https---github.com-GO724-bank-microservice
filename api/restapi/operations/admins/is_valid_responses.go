// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bank-microservice/api/models"
)

// IsValidOKCode is the HTTP code returned for type IsValidOK
const IsValidOKCode int = 200

/*
IsValidOK search results matching criteria

swagger:response isValidOK
*/
type IsValidOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Card `json:"body,omitempty"`
}

// NewIsValidOK creates IsValidOK with default headers values
func NewIsValidOK() *IsValidOK {

	return &IsValidOK{}
}

// WithPayload adds the payload to the is valid o k response
func (o *IsValidOK) WithPayload(payload []*models.Card) *IsValidOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the is valid o k response
func (o *IsValidOK) SetPayload(payload []*models.Card) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IsValidOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Card, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IsValidBadRequestCode is the HTTP code returned for type IsValidBadRequest
const IsValidBadRequestCode int = 400

/*
IsValidBadRequest bad input parameter

swagger:response isValidBadRequest
*/
type IsValidBadRequest struct {
}

// NewIsValidBadRequest creates IsValidBadRequest with default headers values
func NewIsValidBadRequest() *IsValidBadRequest {

	return &IsValidBadRequest{}
}

// WriteResponse to the client
func (o *IsValidBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
