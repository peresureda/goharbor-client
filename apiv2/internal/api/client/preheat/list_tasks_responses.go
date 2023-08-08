// Code generated by go-swagger; DO NOT EDIT.

package preheat

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

// ListTasksReader is a Reader for the ListTasks structure.
type ListTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTasksOK creates a ListTasksOK with default headers values
func NewListTasksOK() *ListTasksOK {
	return &ListTasksOK{}
}

/*ListTasksOK handles this case with default header values.

List tasks success
*/
type ListTasksOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of tasks
	 */
	XTotalCount int64

	Payload []*model.Task
}

func (o *ListTasksOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks][%d] listTasksOK  %+v", 200, o.Payload)
}

func (o *ListTasksOK) GetPayload() []*model.Task {
	return o.Payload
}

func (o *ListTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListTasksBadRequest creates a ListTasksBadRequest with default headers values
func NewListTasksBadRequest() *ListTasksBadRequest {
	return &ListTasksBadRequest{}
}

/*ListTasksBadRequest handles this case with default header values.

Bad request
*/
type ListTasksBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks][%d] listTasksBadRequest  %+v", 400, o.Payload)
}

func (o *ListTasksBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksUnauthorized creates a ListTasksUnauthorized with default headers values
func NewListTasksUnauthorized() *ListTasksUnauthorized {
	return &ListTasksUnauthorized{}
}

/*ListTasksUnauthorized handles this case with default header values.

Unauthorized
*/
type ListTasksUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListTasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks][%d] listTasksUnauthorized  %+v", 401, o.Payload)
}

func (o *ListTasksUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksForbidden creates a ListTasksForbidden with default headers values
func NewListTasksForbidden() *ListTasksForbidden {
	return &ListTasksForbidden{}
}

/*ListTasksForbidden handles this case with default header values.

Forbidden
*/
type ListTasksForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks][%d] listTasksForbidden  %+v", 403, o.Payload)
}

func (o *ListTasksForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksNotFound creates a ListTasksNotFound with default headers values
func NewListTasksNotFound() *ListTasksNotFound {
	return &ListTasksNotFound{}
}

/*ListTasksNotFound handles this case with default header values.

Not found
*/
type ListTasksNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListTasksNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks][%d] listTasksNotFound  %+v", 404, o.Payload)
}

func (o *ListTasksNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListTasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksInternalServerError creates a ListTasksInternalServerError with default headers values
func NewListTasksInternalServerError() *ListTasksInternalServerError {
	return &ListTasksInternalServerError{}
}

/*ListTasksInternalServerError handles this case with default header values.

Internal server error
*/
type ListTasksInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks][%d] listTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *ListTasksInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
