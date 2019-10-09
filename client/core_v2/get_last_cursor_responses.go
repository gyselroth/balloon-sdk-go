// Code generated by go-swagger; DO NOT EDIT.

package core_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLastCursorReader is a Reader for the GetLastCursor structure.
type GetLastCursorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLastCursorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLastCursorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLastCursorOK creates a GetLastCursorOK with default headers values
func NewGetLastCursorOK() *GetLastCursorOK {
	return &GetLastCursorOK{}
}

/*GetLastCursorOK handles this case with default header values.

Cursor
*/
type GetLastCursorOK struct {
	Payload string
}

func (o *GetLastCursorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/nodes/last-cursor][%d] getLastCursorOK  %+v", 200, o.Payload)
}

func (o *GetLastCursorOK) GetPayload() string {
	return o.Payload
}

func (o *GetLastCursorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}