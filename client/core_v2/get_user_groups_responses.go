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

// GetUserGroupsReader is a Reader for the GetUserGroups structure.
type GetUserGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUserGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserGroupsOK creates a GetUserGroupsOK with default headers values
func NewGetUserGroupsOK() *GetUserGroupsOK {
	return &GetUserGroupsOK{}
}

/*GetUserGroupsOK handles this case with default header values.

List of groups
*/
type GetUserGroupsOK struct {
	Payload *models.CoreV2Groups
}

func (o *GetUserGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user}/groups][%d] getUserGroupsOK  %+v", 200, o.Payload)
}

func (o *GetUserGroupsOK) GetPayload() *models.CoreV2Groups {
	return o.Payload
}

func (o *GetUserGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Groups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupsForbidden creates a GetUserGroupsForbidden with default headers values
func NewGetUserGroupsForbidden() *GetUserGroupsForbidden {
	return &GetUserGroupsForbidden{}
}

/*GetUserGroupsForbidden handles this case with default header values.

Access denied
*/
type GetUserGroupsForbidden struct {
}

func (o *GetUserGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user}/groups][%d] getUserGroupsForbidden ", 403)
}

func (o *GetUserGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserGroupsNotFound creates a GetUserGroupsNotFound with default headers values
func NewGetUserGroupsNotFound() *GetUserGroupsNotFound {
	return &GetUserGroupsNotFound{}
}

/*GetUserGroupsNotFound handles this case with default header values.

Resource does not exists
*/
type GetUserGroupsNotFound struct {
}

func (o *GetUserGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user}/groups][%d] getUserGroupsNotFound ", 404)
}

func (o *GetUserGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
