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

// GetSupportedFormatsReader is a Reader for the GetSupportedFormats structure.
type GetSupportedFormatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSupportedFormatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSupportedFormatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSupportedFormatsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSupportedFormatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSupportedFormatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSupportedFormatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetSupportedFormatsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSupportedFormatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSupportedFormatsOK creates a GetSupportedFormatsOK with default headers values
func NewGetSupportedFormatsOK() *GetSupportedFormatsOK {
	return &GetSupportedFormatsOK{}
}

/*GetSupportedFormatsOK handles this case with default header values.

Formats
*/
type GetSupportedFormatsOK struct {
	Payload []string
}

func (o *GetSupportedFormatsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/files/{file}/convert/supported-formats][%d] getSupportedFormatsOK  %+v", 200, o.Payload)
}

func (o *GetSupportedFormatsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetSupportedFormatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedFormatsBadRequest creates a GetSupportedFormatsBadRequest with default headers values
func NewGetSupportedFormatsBadRequest() *GetSupportedFormatsBadRequest {
	return &GetSupportedFormatsBadRequest{}
}

/*GetSupportedFormatsBadRequest handles this case with default header values.

Bad Reqeust
*/
type GetSupportedFormatsBadRequest struct {
	Payload *models.CoreV2Error
}

func (o *GetSupportedFormatsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/files/{file}/convert/supported-formats][%d] getSupportedFormatsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSupportedFormatsBadRequest) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetSupportedFormatsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedFormatsUnauthorized creates a GetSupportedFormatsUnauthorized with default headers values
func NewGetSupportedFormatsUnauthorized() *GetSupportedFormatsUnauthorized {
	return &GetSupportedFormatsUnauthorized{}
}

/*GetSupportedFormatsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSupportedFormatsUnauthorized struct {
	Payload *models.CoreV2Error
}

func (o *GetSupportedFormatsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/files/{file}/convert/supported-formats][%d] getSupportedFormatsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSupportedFormatsUnauthorized) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetSupportedFormatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedFormatsForbidden creates a GetSupportedFormatsForbidden with default headers values
func NewGetSupportedFormatsForbidden() *GetSupportedFormatsForbidden {
	return &GetSupportedFormatsForbidden{}
}

/*GetSupportedFormatsForbidden handles this case with default header values.

Forbidden
*/
type GetSupportedFormatsForbidden struct {
	Payload *models.CoreV2Error
}

func (o *GetSupportedFormatsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/files/{file}/convert/supported-formats][%d] getSupportedFormatsForbidden  %+v", 403, o.Payload)
}

func (o *GetSupportedFormatsForbidden) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetSupportedFormatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedFormatsNotFound creates a GetSupportedFormatsNotFound with default headers values
func NewGetSupportedFormatsNotFound() *GetSupportedFormatsNotFound {
	return &GetSupportedFormatsNotFound{}
}

/*GetSupportedFormatsNotFound handles this case with default header values.

The specified resource was not found
*/
type GetSupportedFormatsNotFound struct {
	Payload *models.CoreV2Error
}

func (o *GetSupportedFormatsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/files/{file}/convert/supported-formats][%d] getSupportedFormatsNotFound  %+v", 404, o.Payload)
}

func (o *GetSupportedFormatsNotFound) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetSupportedFormatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedFormatsUnprocessableEntity creates a GetSupportedFormatsUnprocessableEntity with default headers values
func NewGetSupportedFormatsUnprocessableEntity() *GetSupportedFormatsUnprocessableEntity {
	return &GetSupportedFormatsUnprocessableEntity{}
}

/*GetSupportedFormatsUnprocessableEntity handles this case with default header values.

Unauthorized
*/
type GetSupportedFormatsUnprocessableEntity struct {
	Payload *models.CoreV2Error
}

func (o *GetSupportedFormatsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v2/files/{file}/convert/supported-formats][%d] getSupportedFormatsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetSupportedFormatsUnprocessableEntity) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetSupportedFormatsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedFormatsInternalServerError creates a GetSupportedFormatsInternalServerError with default headers values
func NewGetSupportedFormatsInternalServerError() *GetSupportedFormatsInternalServerError {
	return &GetSupportedFormatsInternalServerError{}
}

/*GetSupportedFormatsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSupportedFormatsInternalServerError struct {
	Payload *models.CoreV2Error
}

func (o *GetSupportedFormatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/files/{file}/convert/supported-formats][%d] getSupportedFormatsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSupportedFormatsInternalServerError) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetSupportedFormatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
