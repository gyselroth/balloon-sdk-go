// Code generated by go-swagger; DO NOT EDIT.

package core_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gyselroth/balloon-sdk-go/models"
)

// DeleteNodeReader is a Reader for the DeleteNode structure.
type DeleteNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNodeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteNodeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNodeNoContent creates a DeleteNodeNoContent with default headers values
func NewDeleteNodeNoContent() *DeleteNodeNoContent {
	return &DeleteNodeNoContent{}
}

/*DeleteNodeNoContent handles this case with default header values.

If successful the server will respond with 204 No Content
*/
type DeleteNodeNoContent struct {
}

func (o *DeleteNodeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}][%d] deleteNodeNoContent ", 204)
}

func (o *DeleteNodeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodeBadRequest creates a DeleteNodeBadRequest with default headers values
func NewDeleteNodeBadRequest() *DeleteNodeBadRequest {
	return &DeleteNodeBadRequest{}
}

/*DeleteNodeBadRequest handles this case with default header values.

Bad Reqeust
*/
type DeleteNodeBadRequest struct {
	Payload *models.CoreV2Error
}

func (o *DeleteNodeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}][%d] deleteNodeBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteNodeBadRequest) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteNodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodeUnauthorized creates a DeleteNodeUnauthorized with default headers values
func NewDeleteNodeUnauthorized() *DeleteNodeUnauthorized {
	return &DeleteNodeUnauthorized{}
}

/*DeleteNodeUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteNodeUnauthorized struct {
	Payload *models.CoreV2Error
}

func (o *DeleteNodeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}][%d] deleteNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteNodeUnauthorized) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodeForbidden creates a DeleteNodeForbidden with default headers values
func NewDeleteNodeForbidden() *DeleteNodeForbidden {
	return &DeleteNodeForbidden{}
}

/*DeleteNodeForbidden handles this case with default header values.

Forbidden
*/
type DeleteNodeForbidden struct {
	Payload *models.CoreV2Error
}

func (o *DeleteNodeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}][%d] deleteNodeForbidden  %+v", 403, o.Payload)
}

func (o *DeleteNodeForbidden) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodeNotFound creates a DeleteNodeNotFound with default headers values
func NewDeleteNodeNotFound() *DeleteNodeNotFound {
	return &DeleteNodeNotFound{}
}

/*DeleteNodeNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteNodeNotFound struct {
	Payload *models.CoreV2Error
}

func (o *DeleteNodeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}][%d] deleteNodeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNodeNotFound) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodeUnprocessableEntity creates a DeleteNodeUnprocessableEntity with default headers values
func NewDeleteNodeUnprocessableEntity() *DeleteNodeUnprocessableEntity {
	return &DeleteNodeUnprocessableEntity{}
}

/*DeleteNodeUnprocessableEntity handles this case with default header values.

Unauthorized
*/
type DeleteNodeUnprocessableEntity struct {
	Payload *models.CoreV2Error
}

func (o *DeleteNodeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}][%d] deleteNodeUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteNodeUnprocessableEntity) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteNodeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodeInternalServerError creates a DeleteNodeInternalServerError with default headers values
func NewDeleteNodeInternalServerError() *DeleteNodeInternalServerError {
	return &DeleteNodeInternalServerError{}
}

/*DeleteNodeInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteNodeInternalServerError struct {
	Payload *models.CoreV2Error
}

func (o *DeleteNodeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}][%d] deleteNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNodeInternalServerError) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
