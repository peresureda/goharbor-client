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

// GetRegistriesReader is a Reader for the GetRegistries structure.
type GetRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRegistriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegistriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegistriesOK creates a GetRegistriesOK with default headers values
func NewGetRegistriesOK() *GetRegistriesOK {
	return &GetRegistriesOK{}
}

/*GetRegistriesOK handles this case with default header values.

List registries successfully.
*/
type GetRegistriesOK struct {
	Payload []*legacy.Registry
}

func (o *GetRegistriesOK) Error() string {
	return fmt.Sprintf("[GET /registries][%d] getRegistriesOK  %+v", 200, o.Payload)
}

func (o *GetRegistriesOK) GetPayload() []*legacy.Registry {
	return o.Payload
}

func (o *GetRegistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistriesUnauthorized creates a GetRegistriesUnauthorized with default headers values
func NewGetRegistriesUnauthorized() *GetRegistriesUnauthorized {
	return &GetRegistriesUnauthorized{}
}

/*GetRegistriesUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetRegistriesUnauthorized struct {
}

func (o *GetRegistriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /registries][%d] getRegistriesUnauthorized ", 401)
}

func (o *GetRegistriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistriesInternalServerError creates a GetRegistriesInternalServerError with default headers values
func NewGetRegistriesInternalServerError() *GetRegistriesInternalServerError {
	return &GetRegistriesInternalServerError{}
}

/*GetRegistriesInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetRegistriesInternalServerError struct {
}

func (o *GetRegistriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registries][%d] getRegistriesInternalServerError ", 500)
}

func (o *GetRegistriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
