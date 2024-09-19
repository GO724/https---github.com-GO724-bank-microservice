// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetPersonReader is a Reader for the GetPerson structure.
type GetPersonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPersonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGetPersonCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPersonBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetPersonConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /persons/{inn}] getPerson", response, response.Code())
	}
}

// NewGetPersonCreated creates a GetPersonCreated with default headers values
func NewGetPersonCreated() *GetPersonCreated {
	return &GetPersonCreated{}
}

/*
GetPersonCreated describes a response with status code 201, with default header values.

person found
*/
type GetPersonCreated struct {
}

// IsSuccess returns true when this get person created response has a 2xx status code
func (o *GetPersonCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get person created response has a 3xx status code
func (o *GetPersonCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get person created response has a 4xx status code
func (o *GetPersonCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this get person created response has a 5xx status code
func (o *GetPersonCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this get person created response a status code equal to that given
func (o *GetPersonCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the get person created response
func (o *GetPersonCreated) Code() int {
	return 201
}

func (o *GetPersonCreated) Error() string {
	return fmt.Sprintf("[GET /persons/{inn}][%d] getPersonCreated", 201)
}

func (o *GetPersonCreated) String() string {
	return fmt.Sprintf("[GET /persons/{inn}][%d] getPersonCreated", 201)
}

func (o *GetPersonCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPersonBadRequest creates a GetPersonBadRequest with default headers values
func NewGetPersonBadRequest() *GetPersonBadRequest {
	return &GetPersonBadRequest{}
}

/*
GetPersonBadRequest describes a response with status code 400, with default header values.

invalid input, object invalid
*/
type GetPersonBadRequest struct {
}

// IsSuccess returns true when this get person bad request response has a 2xx status code
func (o *GetPersonBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get person bad request response has a 3xx status code
func (o *GetPersonBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get person bad request response has a 4xx status code
func (o *GetPersonBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get person bad request response has a 5xx status code
func (o *GetPersonBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get person bad request response a status code equal to that given
func (o *GetPersonBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get person bad request response
func (o *GetPersonBadRequest) Code() int {
	return 400
}

func (o *GetPersonBadRequest) Error() string {
	return fmt.Sprintf("[GET /persons/{inn}][%d] getPersonBadRequest", 400)
}

func (o *GetPersonBadRequest) String() string {
	return fmt.Sprintf("[GET /persons/{inn}][%d] getPersonBadRequest", 400)
}

func (o *GetPersonBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPersonConflict creates a GetPersonConflict with default headers values
func NewGetPersonConflict() *GetPersonConflict {
	return &GetPersonConflict{}
}

/*
GetPersonConflict describes a response with status code 409, with default header values.

person not exists
*/
type GetPersonConflict struct {
}

// IsSuccess returns true when this get person conflict response has a 2xx status code
func (o *GetPersonConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get person conflict response has a 3xx status code
func (o *GetPersonConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get person conflict response has a 4xx status code
func (o *GetPersonConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get person conflict response has a 5xx status code
func (o *GetPersonConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get person conflict response a status code equal to that given
func (o *GetPersonConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get person conflict response
func (o *GetPersonConflict) Code() int {
	return 409
}

func (o *GetPersonConflict) Error() string {
	return fmt.Sprintf("[GET /persons/{inn}][%d] getPersonConflict", 409)
}

func (o *GetPersonConflict) String() string {
	return fmt.Sprintf("[GET /persons/{inn}][%d] getPersonConflict", 409)
}

func (o *GetPersonConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}