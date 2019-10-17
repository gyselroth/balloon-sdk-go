// Code generated by go-swagger; DO NOT EDIT.

package convert_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gyselroth/balloon-sdk-go/models"
)

// DeleteSlaveReader is a Reader for the DeleteSlave structure.
type DeleteSlaveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSlaveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSlaveNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSlaveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSlaveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSlaveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSlaveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteSlaveUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSlaveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSlaveNoContent creates a DeleteSlaveNoContent with default headers values
func NewDeleteSlaveNoContent() *DeleteSlaveNoContent {
	return &DeleteSlaveNoContent{}
}

/*DeleteSlaveNoContent handles this case with default header values.

If successful the server will respond with 204 No Content
*/
type DeleteSlaveNoContent struct {
}

func (o *DeleteSlaveNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/files/{file}/convert/slaves/{slave}][%d] deleteSlaveNoContent ", 204)
}

func (o *DeleteSlaveNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSlaveBadRequest creates a DeleteSlaveBadRequest with default headers values
func NewDeleteSlaveBadRequest() *DeleteSlaveBadRequest {
	return &DeleteSlaveBadRequest{}
}

/*DeleteSlaveBadRequest handles this case with default header values.

Bad Reqeust
*/
type DeleteSlaveBadRequest struct {
	Payload *models.CoreV2Error
}

func (o *DeleteSlaveBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/files/{file}/convert/slaves/{slave}][%d] deleteSlaveBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSlaveBadRequest) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteSlaveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSlaveUnauthorized creates a DeleteSlaveUnauthorized with default headers values
func NewDeleteSlaveUnauthorized() *DeleteSlaveUnauthorized {
	return &DeleteSlaveUnauthorized{}
}

/*DeleteSlaveUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteSlaveUnauthorized struct {
	Payload *models.CoreV2Error
}

func (o *DeleteSlaveUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/files/{file}/convert/slaves/{slave}][%d] deleteSlaveUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteSlaveUnauthorized) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteSlaveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSlaveForbidden creates a DeleteSlaveForbidden with default headers values
func NewDeleteSlaveForbidden() *DeleteSlaveForbidden {
	return &DeleteSlaveForbidden{}
}

/*DeleteSlaveForbidden handles this case with default header values.

Forbidden
*/
type DeleteSlaveForbidden struct {
	Payload *models.CoreV2Error
}

func (o *DeleteSlaveForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/files/{file}/convert/slaves/{slave}][%d] deleteSlaveForbidden  %+v", 403, o.Payload)
}

func (o *DeleteSlaveForbidden) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteSlaveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSlaveNotFound creates a DeleteSlaveNotFound with default headers values
func NewDeleteSlaveNotFound() *DeleteSlaveNotFound {
	return &DeleteSlaveNotFound{}
}

/*DeleteSlaveNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteSlaveNotFound struct {
	Payload *models.CoreV2Error
}

func (o *DeleteSlaveNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/files/{file}/convert/slaves/{slave}][%d] deleteSlaveNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSlaveNotFound) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteSlaveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSlaveUnprocessableEntity creates a DeleteSlaveUnprocessableEntity with default headers values
func NewDeleteSlaveUnprocessableEntity() *DeleteSlaveUnprocessableEntity {
	return &DeleteSlaveUnprocessableEntity{}
}

/*DeleteSlaveUnprocessableEntity handles this case with default header values.

Unauthorized
*/
type DeleteSlaveUnprocessableEntity struct {
	Payload *models.CoreV2Error
}

func (o *DeleteSlaveUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/files/{file}/convert/slaves/{slave}][%d] deleteSlaveUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteSlaveUnprocessableEntity) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteSlaveUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSlaveInternalServerError creates a DeleteSlaveInternalServerError with default headers values
func NewDeleteSlaveInternalServerError() *DeleteSlaveInternalServerError {
	return &DeleteSlaveInternalServerError{}
}

/*DeleteSlaveInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteSlaveInternalServerError struct {
	Payload *models.CoreV2Error
}

func (o *DeleteSlaveInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/files/{file}/convert/slaves/{slave}][%d] deleteSlaveInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSlaveInternalServerError) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *DeleteSlaveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
