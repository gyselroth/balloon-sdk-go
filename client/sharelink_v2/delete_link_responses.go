// Code generated by go-swagger; DO NOT EDIT.

package sharelink_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gyselroth/balloon-sdk-go/models"
)

// DeleteLinkReader is a Reader for the DeleteLink structure.
type DeleteLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLinkNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLinkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteLinkUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLinkNoContent creates a DeleteLinkNoContent with default headers values
func NewDeleteLinkNoContent() *DeleteLinkNoContent {
	return &DeleteLinkNoContent{}
}

/*DeleteLinkNoContent handles this case with default header values.

The server responds with 204 if operation was successful
*/
type DeleteLinkNoContent struct {
}

func (o *DeleteLinkNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}/share-link][%d] deleteLinkNoContent ", 204)
}

func (o *DeleteLinkNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLinkBadRequest creates a DeleteLinkBadRequest with default headers values
func NewDeleteLinkBadRequest() *DeleteLinkBadRequest {
	return &DeleteLinkBadRequest{}
}

/*DeleteLinkBadRequest handles this case with default header values.

Bad Reqeust
*/
type DeleteLinkBadRequest struct {
	Payload *models.CoreV2Error
}

func (o *DeleteLinkBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}/share-link][%d] deleteLinkBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLinkBadRequest) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLinkUnauthorized creates a DeleteLinkUnauthorized with default headers values
func NewDeleteLinkUnauthorized() *DeleteLinkUnauthorized {
	return &DeleteLinkUnauthorized{}
}

/*DeleteLinkUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteLinkUnauthorized struct {
	Payload *models.CoreV2Error
}

func (o *DeleteLinkUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}/share-link][%d] deleteLinkUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteLinkUnauthorized) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteLinkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLinkForbidden creates a DeleteLinkForbidden with default headers values
func NewDeleteLinkForbidden() *DeleteLinkForbidden {
	return &DeleteLinkForbidden{}
}

/*DeleteLinkForbidden handles this case with default header values.

Forbidden
*/
type DeleteLinkForbidden struct {
	Payload *models.CoreV2Error
}

func (o *DeleteLinkForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}/share-link][%d] deleteLinkForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLinkForbidden) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLinkNotFound creates a DeleteLinkNotFound with default headers values
func NewDeleteLinkNotFound() *DeleteLinkNotFound {
	return &DeleteLinkNotFound{}
}

/*DeleteLinkNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteLinkNotFound struct {
	Payload *models.CoreV2Error
}

func (o *DeleteLinkNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}/share-link][%d] deleteLinkNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLinkNotFound) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLinkUnprocessableEntity creates a DeleteLinkUnprocessableEntity with default headers values
func NewDeleteLinkUnprocessableEntity() *DeleteLinkUnprocessableEntity {
	return &DeleteLinkUnprocessableEntity{}
}

/*DeleteLinkUnprocessableEntity handles this case with default header values.

Unauthorized
*/
type DeleteLinkUnprocessableEntity struct {
	Payload *models.CoreV2Error
}

func (o *DeleteLinkUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}/share-link][%d] deleteLinkUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteLinkUnprocessableEntity) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteLinkUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLinkInternalServerError creates a DeleteLinkInternalServerError with default headers values
func NewDeleteLinkInternalServerError() *DeleteLinkInternalServerError {
	return &DeleteLinkInternalServerError{}
}

/*DeleteLinkInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteLinkInternalServerError struct {
	Payload *models.CoreV2Error
}

func (o *DeleteLinkInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/nodes/{node}/share-link][%d] deleteLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLinkInternalServerError) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
