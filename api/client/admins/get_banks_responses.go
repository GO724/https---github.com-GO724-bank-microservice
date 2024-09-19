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

// GetBanksReader is a Reader for the GetBanks structure.
type GetBanksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBanksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBanksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBanksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /banks] GetBanks", response, response.Code())
	}
}

// NewGetBanksOK creates a GetBanksOK with default headers values
func NewGetBanksOK() *GetBanksOK {
	return &GetBanksOK{}
}

/*
GetBanksOK describes a response with status code 200, with default header values.

search results matching criteria
*/
type GetBanksOK struct {
	Payload []*models.Bank
}

// IsSuccess returns true when this get banks o k response has a 2xx status code
func (o *GetBanksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get banks o k response has a 3xx status code
func (o *GetBanksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get banks o k response has a 4xx status code
func (o *GetBanksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get banks o k response has a 5xx status code
func (o *GetBanksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get banks o k response a status code equal to that given
func (o *GetBanksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get banks o k response
func (o *GetBanksOK) Code() int {
	return 200
}

func (o *GetBanksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /banks][%d] getBanksOK %s", 200, payload)
}

func (o *GetBanksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /banks][%d] getBanksOK %s", 200, payload)
}

func (o *GetBanksOK) GetPayload() []*models.Bank {
	return o.Payload
}

func (o *GetBanksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBanksBadRequest creates a GetBanksBadRequest with default headers values
func NewGetBanksBadRequest() *GetBanksBadRequest {
	return &GetBanksBadRequest{}
}

/*
GetBanksBadRequest describes a response with status code 400, with default header values.

bad input parameter
*/
type GetBanksBadRequest struct {
}

// IsSuccess returns true when this get banks bad request response has a 2xx status code
func (o *GetBanksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get banks bad request response has a 3xx status code
func (o *GetBanksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get banks bad request response has a 4xx status code
func (o *GetBanksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get banks bad request response has a 5xx status code
func (o *GetBanksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get banks bad request response a status code equal to that given
func (o *GetBanksBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get banks bad request response
func (o *GetBanksBadRequest) Code() int {
	return 400
}

func (o *GetBanksBadRequest) Error() string {
	return fmt.Sprintf("[GET /banks][%d] getBanksBadRequest", 400)
}

func (o *GetBanksBadRequest) String() string {
	return fmt.Sprintf("[GET /banks][%d] getBanksBadRequest", 400)
}

func (o *GetBanksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}