// Code generated by go-swagger; DO NOT EDIT.

package preview_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetPreviewReader is a Reader for the GetPreview structure.
type GetPreviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPreviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPreviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPreviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPreviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPreviewOK creates a GetPreviewOK with default headers values
func NewGetPreviewOK() *GetPreviewOK {
	return &GetPreviewOK{}
}

/*GetPreviewOK handles this case with default header values.

Binara data
*/
type GetPreviewOK struct {
}

func (o *GetPreviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/files/{file}/preview][%d] getPreviewOK ", 200)
}

func (o *GetPreviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPreviewForbidden creates a GetPreviewForbidden with default headers values
func NewGetPreviewForbidden() *GetPreviewForbidden {
	return &GetPreviewForbidden{}
}

/*GetPreviewForbidden handles this case with default header values.

Access denied
*/
type GetPreviewForbidden struct {
}

func (o *GetPreviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/files/{file}/preview][%d] getPreviewForbidden ", 403)
}

func (o *GetPreviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPreviewNotFound creates a GetPreviewNotFound with default headers values
func NewGetPreviewNotFound() *GetPreviewNotFound {
	return &GetPreviewNotFound{}
}

/*GetPreviewNotFound handles this case with default header values.

Resource does not exists
*/
type GetPreviewNotFound struct {
}

func (o *GetPreviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/files/{file}/preview][%d] getPreviewNotFound ", 404)
}

func (o *GetPreviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
