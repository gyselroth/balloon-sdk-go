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

// GetNodeEventLogReader is a Reader for the GetNodeEventLog structure.
type GetNodeEventLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeEventLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeEventLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNodeEventLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNodeEventLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNodeEventLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNodeEventLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetNodeEventLogUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNodeEventLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNodeEventLogOK creates a GetNodeEventLogOK with default headers values
func NewGetNodeEventLogOK() *GetNodeEventLogOK {
	return &GetNodeEventLogOK{}
}

/*GetNodeEventLogOK handles this case with default header values.

Event log
*/
type GetNodeEventLogOK struct {
	Payload *models.CoreV2EventLogs
}

func (o *GetNodeEventLogOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/nodes/{node}/event-log][%d] getNodeEventLogOK  %+v", 200, o.Payload)
}

func (o *GetNodeEventLogOK) GetPayload() *models.CoreV2EventLogs {
	return o.Payload
}

func (o *GetNodeEventLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2EventLogs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeEventLogBadRequest creates a GetNodeEventLogBadRequest with default headers values
func NewGetNodeEventLogBadRequest() *GetNodeEventLogBadRequest {
	return &GetNodeEventLogBadRequest{}
}

/*GetNodeEventLogBadRequest handles this case with default header values.

Bad Reqeust
*/
type GetNodeEventLogBadRequest struct {
	Payload *models.CoreV2Error
}

func (o *GetNodeEventLogBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/nodes/{node}/event-log][%d] getNodeEventLogBadRequest  %+v", 400, o.Payload)
}

func (o *GetNodeEventLogBadRequest) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetNodeEventLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeEventLogUnauthorized creates a GetNodeEventLogUnauthorized with default headers values
func NewGetNodeEventLogUnauthorized() *GetNodeEventLogUnauthorized {
	return &GetNodeEventLogUnauthorized{}
}

/*GetNodeEventLogUnauthorized handles this case with default header values.

Unauthorized
*/
type GetNodeEventLogUnauthorized struct {
	Payload *models.CoreV2Error
}

func (o *GetNodeEventLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/nodes/{node}/event-log][%d] getNodeEventLogUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNodeEventLogUnauthorized) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetNodeEventLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeEventLogForbidden creates a GetNodeEventLogForbidden with default headers values
func NewGetNodeEventLogForbidden() *GetNodeEventLogForbidden {
	return &GetNodeEventLogForbidden{}
}

/*GetNodeEventLogForbidden handles this case with default header values.

Forbidden
*/
type GetNodeEventLogForbidden struct {
	Payload *models.CoreV2Error
}

func (o *GetNodeEventLogForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/nodes/{node}/event-log][%d] getNodeEventLogForbidden  %+v", 403, o.Payload)
}

func (o *GetNodeEventLogForbidden) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetNodeEventLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeEventLogNotFound creates a GetNodeEventLogNotFound with default headers values
func NewGetNodeEventLogNotFound() *GetNodeEventLogNotFound {
	return &GetNodeEventLogNotFound{}
}

/*GetNodeEventLogNotFound handles this case with default header values.

The specified resource was not found
*/
type GetNodeEventLogNotFound struct {
	Payload *models.CoreV2Error
}

func (o *GetNodeEventLogNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/nodes/{node}/event-log][%d] getNodeEventLogNotFound  %+v", 404, o.Payload)
}

func (o *GetNodeEventLogNotFound) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetNodeEventLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeEventLogUnprocessableEntity creates a GetNodeEventLogUnprocessableEntity with default headers values
func NewGetNodeEventLogUnprocessableEntity() *GetNodeEventLogUnprocessableEntity {
	return &GetNodeEventLogUnprocessableEntity{}
}

/*GetNodeEventLogUnprocessableEntity handles this case with default header values.

Unauthorized
*/
type GetNodeEventLogUnprocessableEntity struct {
	Payload *models.CoreV2Error
}

func (o *GetNodeEventLogUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v2/nodes/{node}/event-log][%d] getNodeEventLogUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetNodeEventLogUnprocessableEntity) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetNodeEventLogUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeEventLogInternalServerError creates a GetNodeEventLogInternalServerError with default headers values
func NewGetNodeEventLogInternalServerError() *GetNodeEventLogInternalServerError {
	return &GetNodeEventLogInternalServerError{}
}

/*GetNodeEventLogInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetNodeEventLogInternalServerError struct {
	Payload *models.CoreV2Error
}

func (o *GetNodeEventLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/nodes/{node}/event-log][%d] getNodeEventLogInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNodeEventLogInternalServerError) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetNodeEventLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
