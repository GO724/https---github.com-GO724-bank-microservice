// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"bank-microservice/api/models"
)

// GetCardsReader is a Reader for the GetCards structure.
type GetCardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCardsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cards] GetCards", response, response.Code())
	}
}

// NewGetCardsOK creates a GetCardsOK with default headers values
func NewGetCardsOK() *GetCardsOK {
	return &GetCardsOK{}
}

/*
GetCardsOK describes a response with status code 200, with default header values.

search results matching criteria
*/
type GetCardsOK struct {
	Payload []*models.Card
}

// IsSuccess returns true when this get cards o k response has a 2xx status code
func (o *GetCardsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cards o k response has a 3xx status code
func (o *GetCardsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cards o k response has a 4xx status code
func (o *GetCardsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cards o k response has a 5xx status code
func (o *GetCardsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cards o k response a status code equal to that given
func (o *GetCardsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cards o k response
func (o *GetCardsOK) Code() int {
	return 200
}

func (o *GetCardsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cards][%d] getCardsOK %s", 200, payload)
}

func (o *GetCardsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cards][%d] getCardsOK %s", 200, payload)
}

func (o *GetCardsOK) GetPayload() []*models.Card {
	return o.Payload
}

func (o *GetCardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCardsBadRequest creates a GetCardsBadRequest with default headers values
func NewGetCardsBadRequest() *GetCardsBadRequest {
	return &GetCardsBadRequest{}
}

/*
GetCardsBadRequest describes a response with status code 400, with default header values.

bad input parameter
*/
type GetCardsBadRequest struct {
}

// IsSuccess returns true when this get cards bad request response has a 2xx status code
func (o *GetCardsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cards bad request response has a 3xx status code
func (o *GetCardsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cards bad request response has a 4xx status code
func (o *GetCardsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cards bad request response has a 5xx status code
func (o *GetCardsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cards bad request response a status code equal to that given
func (o *GetCardsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get cards bad request response
func (o *GetCardsBadRequest) Code() int {
	return 400
}

func (o *GetCardsBadRequest) Error() string {
	return fmt.Sprintf("[GET /cards][%d] getCardsBadRequest", 400)
}

func (o *GetCardsBadRequest) String() string {
	return fmt.Sprintf("[GET /cards][%d] getCardsBadRequest", 400)
}

func (o *GetCardsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}