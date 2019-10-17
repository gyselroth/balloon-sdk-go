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

// AddSlaveReader is a Reader for the AddSlave structure.
type AddSlaveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSlaveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewAddSlaveAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddSlaveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddSlaveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddSlaveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddSlaveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAddSlaveUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddSlaveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddSlaveAccepted creates a AddSlaveAccepted with default headers values
func NewAddSlaveAccepted() *AddSlaveAccepted {
	return &AddSlaveAccepted{}
}

/*AddSlaveAccepted handles this case with default header values.

Slave
*/
type AddSlaveAccepted struct {
	Payload *models.ConvertV2Slave
}

func (o *AddSlaveAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/files/{file}/convert/slaves][%d] addSlaveAccepted  %+v", 202, o.Payload)
}

func (o *AddSlaveAccepted) GetPayload() *models.ConvertV2Slave {
	return o.Payload
}

func (o *AddSlaveAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConvertV2Slave)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSlaveBadRequest creates a AddSlaveBadRequest with default headers values
func NewAddSlaveBadRequest() *AddSlaveBadRequest {
	return &AddSlaveBadRequest{}
}

/*AddSlaveBadRequest handles this case with default header values.

Bad Reqeust
*/
type AddSlaveBadRequest struct {
	Payload *models.CoreV2Error
}

func (o *AddSlaveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/files/{file}/convert/slaves][%d] addSlaveBadRequest  %+v", 400, o.Payload)
}

func (o *AddSlaveBadRequest) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddSlaveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSlaveUnauthorized creates a AddSlaveUnauthorized with default headers values
func NewAddSlaveUnauthorized() *AddSlaveUnauthorized {
	return &AddSlaveUnauthorized{}
}

/*AddSlaveUnauthorized handles this case with default header values.

Unauthorized
*/
type AddSlaveUnauthorized struct {
	Payload *models.CoreV2Error
}

func (o *AddSlaveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/files/{file}/convert/slaves][%d] addSlaveUnauthorized  %+v", 401, o.Payload)
}

func (o *AddSlaveUnauthorized) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddSlaveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSlaveForbidden creates a AddSlaveForbidden with default headers values
func NewAddSlaveForbidden() *AddSlaveForbidden {
	return &AddSlaveForbidden{}
}

/*AddSlaveForbidden handles this case with default header values.

Forbidden
*/
type AddSlaveForbidden struct {
	Payload *models.CoreV2Error
}

func (o *AddSlaveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/files/{file}/convert/slaves][%d] addSlaveForbidden  %+v", 403, o.Payload)
}

func (o *AddSlaveForbidden) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddSlaveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSlaveNotFound creates a AddSlaveNotFound with default headers values
func NewAddSlaveNotFound() *AddSlaveNotFound {
	return &AddSlaveNotFound{}
}

/*AddSlaveNotFound handles this case with default header values.

The specified resource was not found
*/
type AddSlaveNotFound struct {
	Payload *models.CoreV2Error
}

func (o *AddSlaveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/files/{file}/convert/slaves][%d] addSlaveNotFound  %+v", 404, o.Payload)
}

func (o *AddSlaveNotFound) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddSlaveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSlaveUnprocessableEntity creates a AddSlaveUnprocessableEntity with default headers values
func NewAddSlaveUnprocessableEntity() *AddSlaveUnprocessableEntity {
	return &AddSlaveUnprocessableEntity{}
}

/*AddSlaveUnprocessableEntity handles this case with default header values.

Unauthorized
*/
type AddSlaveUnprocessableEntity struct {
	Payload *models.CoreV2Error
}

func (o *AddSlaveUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/files/{file}/convert/slaves][%d] addSlaveUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AddSlaveUnprocessableEntity) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddSlaveUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSlaveInternalServerError creates a AddSlaveInternalServerError with default headers values
func NewAddSlaveInternalServerError() *AddSlaveInternalServerError {
	return &AddSlaveInternalServerError{}
}

/*AddSlaveInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddSlaveInternalServerError struct {
	Payload *models.CoreV2Error
}

func (o *AddSlaveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/files/{file}/convert/slaves][%d] addSlaveInternalServerError  %+v", 500, o.Payload)
}

func (o *AddSlaveInternalServerError) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddSlaveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
