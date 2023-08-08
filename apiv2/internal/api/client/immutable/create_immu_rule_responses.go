// Code generated by go-swagger; DO NOT EDIT.

package immutable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/peresureda/goharbor-client/v5/apiv2/model"
)

// CreateImmuRuleReader is a Reader for the CreateImmuRule structure.
type CreateImmuRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateImmuRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateImmuRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateImmuRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateImmuRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateImmuRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateImmuRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateImmuRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateImmuRuleCreated creates a CreateImmuRuleCreated with default headers values
func NewCreateImmuRuleCreated() *CreateImmuRuleCreated {
	return &CreateImmuRuleCreated{}
}

/*CreateImmuRuleCreated handles this case with default header values.

Created
*/
type CreateImmuRuleCreated struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CreateImmuRuleCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/immutabletagrules][%d] createImmuRuleCreated ", 201)
}

func (o *CreateImmuRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewCreateImmuRuleBadRequest creates a CreateImmuRuleBadRequest with default headers values
func NewCreateImmuRuleBadRequest() *CreateImmuRuleBadRequest {
	return &CreateImmuRuleBadRequest{}
}

/*CreateImmuRuleBadRequest handles this case with default header values.

Bad request
*/
type CreateImmuRuleBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateImmuRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/immutabletagrules][%d] createImmuRuleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateImmuRuleBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateImmuRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImmuRuleUnauthorized creates a CreateImmuRuleUnauthorized with default headers values
func NewCreateImmuRuleUnauthorized() *CreateImmuRuleUnauthorized {
	return &CreateImmuRuleUnauthorized{}
}

/*CreateImmuRuleUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateImmuRuleUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateImmuRuleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/immutabletagrules][%d] createImmuRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateImmuRuleUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateImmuRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImmuRuleForbidden creates a CreateImmuRuleForbidden with default headers values
func NewCreateImmuRuleForbidden() *CreateImmuRuleForbidden {
	return &CreateImmuRuleForbidden{}
}

/*CreateImmuRuleForbidden handles this case with default header values.

Forbidden
*/
type CreateImmuRuleForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateImmuRuleForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/immutabletagrules][%d] createImmuRuleForbidden  %+v", 403, o.Payload)
}

func (o *CreateImmuRuleForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateImmuRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImmuRuleNotFound creates a CreateImmuRuleNotFound with default headers values
func NewCreateImmuRuleNotFound() *CreateImmuRuleNotFound {
	return &CreateImmuRuleNotFound{}
}

/*CreateImmuRuleNotFound handles this case with default header values.

Not found
*/
type CreateImmuRuleNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateImmuRuleNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/immutabletagrules][%d] createImmuRuleNotFound  %+v", 404, o.Payload)
}

func (o *CreateImmuRuleNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateImmuRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImmuRuleInternalServerError creates a CreateImmuRuleInternalServerError with default headers values
func NewCreateImmuRuleInternalServerError() *CreateImmuRuleInternalServerError {
	return &CreateImmuRuleInternalServerError{}
}

/*CreateImmuRuleInternalServerError handles this case with default header values.

Internal server error
*/
type CreateImmuRuleInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateImmuRuleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/immutabletagrules][%d] createImmuRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateImmuRuleInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateImmuRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
