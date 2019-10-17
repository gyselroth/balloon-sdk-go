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

// AddUserReader is a Reader for the AddUser structure.
type AddUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAddUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddUserCreated creates a AddUserCreated with default headers values
func NewAddUserCreated() *AddUserCreated {
	return &AddUserCreated{}
}

/*AddUserCreated handles this case with default header values.

If successful the server will respond with 201 Created
*/
type AddUserCreated struct {
	Payload *models.CoreV2User
}

func (o *AddUserCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/users][%d] addUserCreated  %+v", 201, o.Payload)
}

func (o *AddUserCreated) GetPayload() *models.CoreV2User {
	return o.Payload
}

func (o *AddUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserBadRequest creates a AddUserBadRequest with default headers values
func NewAddUserBadRequest() *AddUserBadRequest {
	return &AddUserBadRequest{}
}

/*AddUserBadRequest handles this case with default header values.

Bad Reqeust
*/
type AddUserBadRequest struct {
	Payload *models.CoreV2Error
}

func (o *AddUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users][%d] addUserBadRequest  %+v", 400, o.Payload)
}

func (o *AddUserBadRequest) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserUnauthorized creates a AddUserUnauthorized with default headers values
func NewAddUserUnauthorized() *AddUserUnauthorized {
	return &AddUserUnauthorized{}
}

/*AddUserUnauthorized handles this case with default header values.

Unauthorized
*/
type AddUserUnauthorized struct {
	Payload *models.CoreV2Error
}

func (o *AddUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/users][%d] addUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AddUserUnauthorized) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserForbidden creates a AddUserForbidden with default headers values
func NewAddUserForbidden() *AddUserForbidden {
	return &AddUserForbidden{}
}

/*AddUserForbidden handles this case with default header values.

Forbidden
*/
type AddUserForbidden struct {
	Payload *models.CoreV2Error
}

func (o *AddUserForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users][%d] addUserForbidden  %+v", 403, o.Payload)
}

func (o *AddUserForbidden) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserNotFound creates a AddUserNotFound with default headers values
func NewAddUserNotFound() *AddUserNotFound {
	return &AddUserNotFound{}
}

/*AddUserNotFound handles this case with default header values.

The specified resource was not found
*/
type AddUserNotFound struct {
	Payload *models.CoreV2Error
}

func (o *AddUserNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/users][%d] addUserNotFound  %+v", 404, o.Payload)
}

func (o *AddUserNotFound) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserUnprocessableEntity creates a AddUserUnprocessableEntity with default headers values
func NewAddUserUnprocessableEntity() *AddUserUnprocessableEntity {
	return &AddUserUnprocessableEntity{}
}

/*AddUserUnprocessableEntity handles this case with default header values.

Unauthorized
*/
type AddUserUnprocessableEntity struct {
	Payload *models.CoreV2Error
}

func (o *AddUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/users][%d] addUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AddUserUnprocessableEntity) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserInternalServerError creates a AddUserInternalServerError with default headers values
func NewAddUserInternalServerError() *AddUserInternalServerError {
	return &AddUserInternalServerError{}
}

/*AddUserInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddUserInternalServerError struct {
	Payload *models.CoreV2Error
}

func (o *AddUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/users][%d] addUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUserInternalServerError) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *AddUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
