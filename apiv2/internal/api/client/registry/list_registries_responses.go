// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/peresureda/goharbor-client/v5/apiv2/model"
)

// ListRegistriesReader is a Reader for the ListRegistries structure.
type ListRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRegistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRegistriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRegistriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRegistriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRegistriesOK creates a ListRegistriesOK with default headers values
func NewListRegistriesOK() *ListRegistriesOK {
	return &ListRegistriesOK{}
}

/*ListRegistriesOK handles this case with default header values.

Success
*/
type ListRegistriesOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of the resources
	 */
	XTotalCount int64

	Payload []*model.Registry
}

func (o *ListRegistriesOK) Error() string {
	return fmt.Sprintf("[GET /registries][%d] listRegistriesOK  %+v", 200, o.Payload)
}

func (o *ListRegistriesOK) GetPayload() []*model.Registry {
	return o.Payload
}

func (o *ListRegistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistriesUnauthorized creates a ListRegistriesUnauthorized with default headers values
func NewListRegistriesUnauthorized() *ListRegistriesUnauthorized {
	return &ListRegistriesUnauthorized{}
}

/*ListRegistriesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListRegistriesUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRegistriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /registries][%d] listRegistriesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRegistriesUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRegistriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistriesForbidden creates a ListRegistriesForbidden with default headers values
func NewListRegistriesForbidden() *ListRegistriesForbidden {
	return &ListRegistriesForbidden{}
}

/*ListRegistriesForbidden handles this case with default header values.

Forbidden
*/
type ListRegistriesForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRegistriesForbidden) Error() string {
	return fmt.Sprintf("[GET /registries][%d] listRegistriesForbidden  %+v", 403, o.Payload)
}

func (o *ListRegistriesForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRegistriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistriesInternalServerError creates a ListRegistriesInternalServerError with default headers values
func NewListRegistriesInternalServerError() *ListRegistriesInternalServerError {
	return &ListRegistriesInternalServerError{}
}

/*ListRegistriesInternalServerError handles this case with default header values.

Internal server error
*/
type ListRegistriesInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRegistriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registries][%d] listRegistriesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRegistriesInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRegistriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
