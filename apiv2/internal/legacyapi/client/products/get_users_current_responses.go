// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v2/apiv2/model/legacy"
)

// GetUsersCurrentReader is a Reader for the GetUsersCurrent structure.
type GetUsersCurrentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersCurrentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersCurrentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsersCurrentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersCurrentOK creates a GetUsersCurrentOK with default headers values
func NewGetUsersCurrentOK() *GetUsersCurrentOK {
	return &GetUsersCurrentOK{}
}

/*GetUsersCurrentOK handles this case with default header values.

Get current user information successfully.
*/
type GetUsersCurrentOK struct {
	Payload *legacy.User
}

func (o *GetUsersCurrentOK) Error() string {
	return fmt.Sprintf("[GET /users/current][%d] getUsersCurrentOK  %+v", 200, o.Payload)
}

func (o *GetUsersCurrentOK) GetPayload() *legacy.User {
	return o.Payload
}

func (o *GetUsersCurrentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legacy.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersCurrentUnauthorized creates a GetUsersCurrentUnauthorized with default headers values
func NewGetUsersCurrentUnauthorized() *GetUsersCurrentUnauthorized {
	return &GetUsersCurrentUnauthorized{}
}

/*GetUsersCurrentUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetUsersCurrentUnauthorized struct {
}

func (o *GetUsersCurrentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current][%d] getUsersCurrentUnauthorized ", 401)
}

func (o *GetUsersCurrentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
