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

// PatchUserReader is a Reader for the PatchUser structure.
type PatchUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchUserOK creates a PatchUserOK with default headers values
func NewPatchUserOK() *PatchUserOK {
	return &PatchUserOK{}
}

/*PatchUserOK handles this case with default header values.

User
*/
type PatchUserOK struct {
	Payload *models.CoreV2User
}

func (o *PatchUserOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{user}][%d] patchUserOK  %+v", 200, o.Payload)
}

func (o *PatchUserOK) GetPayload() *models.CoreV2User {
	return o.Payload
}

func (o *PatchUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserForbidden creates a PatchUserForbidden with default headers values
func NewPatchUserForbidden() *PatchUserForbidden {
	return &PatchUserForbidden{}
}

/*PatchUserForbidden handles this case with default header values.

Access denied
*/
type PatchUserForbidden struct {
}

func (o *PatchUserForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{user}][%d] patchUserForbidden ", 403)
}

func (o *PatchUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchUserNotFound creates a PatchUserNotFound with default headers values
func NewPatchUserNotFound() *PatchUserNotFound {
	return &PatchUserNotFound{}
}

/*PatchUserNotFound handles this case with default header values.

Resource does not exists
*/
type PatchUserNotFound struct {
}

func (o *PatchUserNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{user}][%d] patchUserNotFound ", 404)
}

func (o *PatchUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
