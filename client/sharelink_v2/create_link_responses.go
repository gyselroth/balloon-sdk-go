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

// CreateLinkReader is a Reader for the CreateLink structure.
type CreateLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateLinkCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateLinkCreated creates a CreateLinkCreated with default headers values
func NewCreateLinkCreated() *CreateLinkCreated {
	return &CreateLinkCreated{}
}

/*CreateLinkCreated handles this case with default header values.

Node
*/
type CreateLinkCreated struct {
	Payload *models.CoreV2Node
}

func (o *CreateLinkCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/share-link][%d] createLinkCreated  %+v", 201, o.Payload)
}

func (o *CreateLinkCreated) GetPayload() *models.CoreV2Node {
	return o.Payload
}

func (o *CreateLinkCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLinkForbidden creates a CreateLinkForbidden with default headers values
func NewCreateLinkForbidden() *CreateLinkForbidden {
	return &CreateLinkForbidden{}
}

/*CreateLinkForbidden handles this case with default header values.

Access denied
*/
type CreateLinkForbidden struct {
}

func (o *CreateLinkForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/share-link][%d] createLinkForbidden ", 403)
}

func (o *CreateLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLinkNotFound creates a CreateLinkNotFound with default headers values
func NewCreateLinkNotFound() *CreateLinkNotFound {
	return &CreateLinkNotFound{}
}

/*CreateLinkNotFound handles this case with default header values.

Resource does not exists
*/
type CreateLinkNotFound struct {
}

func (o *CreateLinkNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/share-link][%d] createLinkNotFound ", 404)
}

func (o *CreateLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
