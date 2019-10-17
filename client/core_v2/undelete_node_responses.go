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

// UndeleteNodeReader is a Reader for the UndeleteNode structure.
type UndeleteNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UndeleteNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUndeleteNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewUndeleteNodeMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUndeleteNodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUndeleteNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUndeleteNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUndeleteNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUndeleteNodeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUndeleteNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUndeleteNodeOK creates a UndeleteNodeOK with default headers values
func NewUndeleteNodeOK() *UndeleteNodeOK {
	return &UndeleteNodeOK{}
}

/*UndeleteNodeOK handles this case with default header values.

Node
*/
type UndeleteNodeOK struct {
	Payload *models.CoreV2Node
}

func (o *UndeleteNodeOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/undelete][%d] undeleteNodeOK  %+v", 200, o.Payload)
}

func (o *UndeleteNodeOK) GetPayload() *models.CoreV2Node {
	return o.Payload
}

func (o *UndeleteNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUndeleteNodeMultiStatus creates a UndeleteNodeMultiStatus with default headers values
func NewUndeleteNodeMultiStatus() *UndeleteNodeMultiStatus {
	return &UndeleteNodeMultiStatus{}
}

/*UndeleteNodeMultiStatus handles this case with default header values.

Multi status if batch request
*/
type UndeleteNodeMultiStatus struct {
}

func (o *UndeleteNodeMultiStatus) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/undelete][%d] undeleteNodeMultiStatus ", 207)
}

func (o *UndeleteNodeMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUndeleteNodeBadRequest creates a UndeleteNodeBadRequest with default headers values
func NewUndeleteNodeBadRequest() *UndeleteNodeBadRequest {
	return &UndeleteNodeBadRequest{}
}

/*UndeleteNodeBadRequest handles this case with default header values.

Bad Reqeust
*/
type UndeleteNodeBadRequest struct {
	Payload *models.CoreV2Error
}

func (o *UndeleteNodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/undelete][%d] undeleteNodeBadRequest  %+v", 400, o.Payload)
}

func (o *UndeleteNodeBadRequest) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *UndeleteNodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUndeleteNodeUnauthorized creates a UndeleteNodeUnauthorized with default headers values
func NewUndeleteNodeUnauthorized() *UndeleteNodeUnauthorized {
	return &UndeleteNodeUnauthorized{}
}

/*UndeleteNodeUnauthorized handles this case with default header values.

Unauthorized
*/
type UndeleteNodeUnauthorized struct {
	Payload *models.CoreV2Error
}

func (o *UndeleteNodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/undelete][%d] undeleteNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *UndeleteNodeUnauthorized) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *UndeleteNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUndeleteNodeForbidden creates a UndeleteNodeForbidden with default headers values
func NewUndeleteNodeForbidden() *UndeleteNodeForbidden {
	return &UndeleteNodeForbidden{}
}

/*UndeleteNodeForbidden handles this case with default header values.

Forbidden
*/
type UndeleteNodeForbidden struct {
	Payload *models.CoreV2Error
}

func (o *UndeleteNodeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/undelete][%d] undeleteNodeForbidden  %+v", 403, o.Payload)
}

func (o *UndeleteNodeForbidden) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *UndeleteNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUndeleteNodeNotFound creates a UndeleteNodeNotFound with default headers values
func NewUndeleteNodeNotFound() *UndeleteNodeNotFound {
	return &UndeleteNodeNotFound{}
}

/*UndeleteNodeNotFound handles this case with default header values.

The specified resource was not found
*/
type UndeleteNodeNotFound struct {
	Payload *models.CoreV2Error
}

func (o *UndeleteNodeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/undelete][%d] undeleteNodeNotFound  %+v", 404, o.Payload)
}

func (o *UndeleteNodeNotFound) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *UndeleteNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUndeleteNodeUnprocessableEntity creates a UndeleteNodeUnprocessableEntity with default headers values
func NewUndeleteNodeUnprocessableEntity() *UndeleteNodeUnprocessableEntity {
	return &UndeleteNodeUnprocessableEntity{}
}

/*UndeleteNodeUnprocessableEntity handles this case with default header values.

Unauthorized
*/
type UndeleteNodeUnprocessableEntity struct {
	Payload *models.CoreV2Error
}

func (o *UndeleteNodeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/undelete][%d] undeleteNodeUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UndeleteNodeUnprocessableEntity) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *UndeleteNodeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUndeleteNodeInternalServerError creates a UndeleteNodeInternalServerError with default headers values
func NewUndeleteNodeInternalServerError() *UndeleteNodeInternalServerError {
	return &UndeleteNodeInternalServerError{}
}

/*UndeleteNodeInternalServerError handles this case with default header values.

Internal Server Error
*/
type UndeleteNodeInternalServerError struct {
	Payload *models.CoreV2Error
}

func (o *UndeleteNodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/undelete][%d] undeleteNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *UndeleteNodeInternalServerError) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *UndeleteNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
