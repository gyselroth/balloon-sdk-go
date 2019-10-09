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

// PatchGroupReader is a Reader for the PatchGroup structure.
type PatchGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchGroupOK creates a PatchGroupOK with default headers values
func NewPatchGroupOK() *PatchGroupOK {
	return &PatchGroupOK{}
}

/*PatchGroupOK handles this case with default header values.

Group
*/
type PatchGroupOK struct {
	Payload *models.CoreV2Group
}

func (o *PatchGroupOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/groups/{group}][%d] patchGroupOK  %+v", 200, o.Payload)
}

func (o *PatchGroupOK) GetPayload() *models.CoreV2Group {
	return o.Payload
}

func (o *PatchGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGroupForbidden creates a PatchGroupForbidden with default headers values
func NewPatchGroupForbidden() *PatchGroupForbidden {
	return &PatchGroupForbidden{}
}

/*PatchGroupForbidden handles this case with default header values.

Access denied
*/
type PatchGroupForbidden struct {
}

func (o *PatchGroupForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/groups/{group}][%d] patchGroupForbidden ", 403)
}

func (o *PatchGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGroupNotFound creates a PatchGroupNotFound with default headers values
func NewPatchGroupNotFound() *PatchGroupNotFound {
	return &PatchGroupNotFound{}
}

/*PatchGroupNotFound handles this case with default header values.

Resource does not exists
*/
type PatchGroupNotFound struct {
}

func (o *PatchGroupNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/groups/{group}][%d] patchGroupNotFound ", 404)
}

func (o *PatchGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
