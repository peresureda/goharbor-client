// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/peresureda/goharbor-client/v5/apiv1/model"
)

// GetReplicationExecutionsIDReader is a Reader for the GetReplicationExecutionsID structure.
type GetReplicationExecutionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationExecutionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationExecutionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReplicationExecutionsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReplicationExecutionsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplicationExecutionsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReplicationExecutionsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetReplicationExecutionsIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReplicationExecutionsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReplicationExecutionsIDOK creates a GetReplicationExecutionsIDOK with default headers values
func NewGetReplicationExecutionsIDOK() *GetReplicationExecutionsIDOK {
	return &GetReplicationExecutionsIDOK{}
}

/*GetReplicationExecutionsIDOK handles this case with default header values.

Success.
*/
type GetReplicationExecutionsIDOK struct {
	Payload *model.ReplicationExecution
}

func (o *GetReplicationExecutionsIDOK) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionsIdOK  %+v", 200, o.Payload)
}

func (o *GetReplicationExecutionsIDOK) GetPayload() *model.ReplicationExecution {
	return o.Payload
}

func (o *GetReplicationExecutionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ReplicationExecution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionsIDBadRequest creates a GetReplicationExecutionsIDBadRequest with default headers values
func NewGetReplicationExecutionsIDBadRequest() *GetReplicationExecutionsIDBadRequest {
	return &GetReplicationExecutionsIDBadRequest{}
}

/*GetReplicationExecutionsIDBadRequest handles this case with default header values.

Bad request.
*/
type GetReplicationExecutionsIDBadRequest struct {
}

func (o *GetReplicationExecutionsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionsIdBadRequest ", 400)
}

func (o *GetReplicationExecutionsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsIDUnauthorized creates a GetReplicationExecutionsIDUnauthorized with default headers values
func NewGetReplicationExecutionsIDUnauthorized() *GetReplicationExecutionsIDUnauthorized {
	return &GetReplicationExecutionsIDUnauthorized{}
}

/*GetReplicationExecutionsIDUnauthorized handles this case with default header values.

User need to login first.
*/
type GetReplicationExecutionsIDUnauthorized struct {
}

func (o *GetReplicationExecutionsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionsIdUnauthorized ", 401)
}

func (o *GetReplicationExecutionsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsIDForbidden creates a GetReplicationExecutionsIDForbidden with default headers values
func NewGetReplicationExecutionsIDForbidden() *GetReplicationExecutionsIDForbidden {
	return &GetReplicationExecutionsIDForbidden{}
}

/*GetReplicationExecutionsIDForbidden handles this case with default header values.

User has no privilege for the operation.
*/
type GetReplicationExecutionsIDForbidden struct {
}

func (o *GetReplicationExecutionsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionsIdForbidden ", 403)
}

func (o *GetReplicationExecutionsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsIDNotFound creates a GetReplicationExecutionsIDNotFound with default headers values
func NewGetReplicationExecutionsIDNotFound() *GetReplicationExecutionsIDNotFound {
	return &GetReplicationExecutionsIDNotFound{}
}

/*GetReplicationExecutionsIDNotFound handles this case with default header values.

Resource requested does not exist.
*/
type GetReplicationExecutionsIDNotFound struct {
}

func (o *GetReplicationExecutionsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionsIdNotFound ", 404)
}

func (o *GetReplicationExecutionsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsIDUnsupportedMediaType creates a GetReplicationExecutionsIDUnsupportedMediaType with default headers values
func NewGetReplicationExecutionsIDUnsupportedMediaType() *GetReplicationExecutionsIDUnsupportedMediaType {
	return &GetReplicationExecutionsIDUnsupportedMediaType{}
}

/*GetReplicationExecutionsIDUnsupportedMediaType handles this case with default header values.

The Media Type of the request is not supported, it has to be "application/json"
*/
type GetReplicationExecutionsIDUnsupportedMediaType struct {
}

func (o *GetReplicationExecutionsIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionsIdUnsupportedMediaType ", 415)
}

func (o *GetReplicationExecutionsIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsIDInternalServerError creates a GetReplicationExecutionsIDInternalServerError with default headers values
func NewGetReplicationExecutionsIDInternalServerError() *GetReplicationExecutionsIDInternalServerError {
	return &GetReplicationExecutionsIDInternalServerError{}
}

/*GetReplicationExecutionsIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetReplicationExecutionsIDInternalServerError struct {
}

func (o *GetReplicationExecutionsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionsIdInternalServerError ", 500)
}

func (o *GetReplicationExecutionsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
