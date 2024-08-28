// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DelСardNoContentCode is the HTTP code returned for type DelСardNoContent
const DelСardNoContentCode int = 204

/*
DelСardNoContent card deleted

swagger:response delСardNoContent
*/
type DelСardNoContent struct {
}

// NewDelСardNoContent creates DelСardNoContent with default headers values
func NewDelСardNoContent() *DelСardNoContent {

	return &DelСardNoContent{}
}

// WriteResponse to the client
func (o *DelСardNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DelСardBadRequestCode is the HTTP code returned for type DelСardBadRequest
const DelСardBadRequestCode int = 400

/*
DelСardBadRequest invalid input, object invalid

swagger:response delСardBadRequest
*/
type DelСardBadRequest struct {
}

// NewDelСardBadRequest creates DelСardBadRequest with default headers values
func NewDelСardBadRequest() *DelСardBadRequest {

	return &DelСardBadRequest{}
}

// WriteResponse to the client
func (o *DelСardBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DelСardConflictCode is the HTTP code returned for type DelСardConflict
const DelСardConflictCode int = 409

/*
DelСardConflict card not exists

swagger:response delСardConflict
*/
type DelСardConflict struct {
}

// NewDelСardConflict creates DelСardConflict with default headers values
func NewDelСardConflict() *DelСardConflict {

	return &DelСardConflict{}
}

// WriteResponse to the client
func (o *DelСardConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}
