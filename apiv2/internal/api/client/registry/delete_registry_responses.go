// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/peresureda/goharbor-client/v5/apiv2/model"
)

// DeleteRegistryReader is a Reader for the DeleteRegistry structure.
type DeleteRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeleteRegistryPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRegistryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRegistryOK creates a DeleteRegistryOK with default headers values
func NewDeleteRegistryOK() *DeleteRegistryOK {
	return &DeleteRegistryOK{}
}

/*DeleteRegistryOK handles this case with default header values.

Success
*/
type DeleteRegistryOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteRegistryOK) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryOK ", 200)
}

func (o *DeleteRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewDeleteRegistryUnauthorized creates a DeleteRegistryUnauthorized with default headers values
func NewDeleteRegistryUnauthorized() *DeleteRegistryUnauthorized {
	return &DeleteRegistryUnauthorized{}
}

/*DeleteRegistryUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteRegistryUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRegistryUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistryForbidden creates a DeleteRegistryForbidden with default headers values
func NewDeleteRegistryForbidden() *DeleteRegistryForbidden {
	return &DeleteRegistryForbidden{}
}

/*DeleteRegistryForbidden handles this case with default header values.

Forbidden
*/
type DeleteRegistryForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRegistryForbidden) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRegistryForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistryNotFound creates a DeleteRegistryNotFound with default headers values
func NewDeleteRegistryNotFound() *DeleteRegistryNotFound {
	return &DeleteRegistryNotFound{}
}

/*DeleteRegistryNotFound handles this case with default header values.

Not found
*/
type DeleteRegistryNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRegistryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRegistryNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistryPreconditionFailed creates a DeleteRegistryPreconditionFailed with default headers values
func NewDeleteRegistryPreconditionFailed() *DeleteRegistryPreconditionFailed {
	return &DeleteRegistryPreconditionFailed{}
}

/*DeleteRegistryPreconditionFailed handles this case with default header values.

Precondition failed
*/
type DeleteRegistryPreconditionFailed struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRegistryPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteRegistryPreconditionFailed) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRegistryPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistryInternalServerError creates a DeleteRegistryInternalServerError with default headers values
func NewDeleteRegistryInternalServerError() *DeleteRegistryInternalServerError {
	return &DeleteRegistryInternalServerError{}
}

/*DeleteRegistryInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteRegistryInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRegistryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRegistryInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRegistryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
