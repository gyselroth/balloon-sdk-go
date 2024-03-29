// Code generated by go-swagger; DO NOT EDIT.

package notification_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gyselroth/balloon-sdk-go/models"
)

// SendBroadcastReader is a Reader for the SendBroadcast structure.
type SendBroadcastReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendBroadcastReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSendBroadcastAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSendBroadcastForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendBroadcastAccepted creates a SendBroadcastAccepted with default headers values
func NewSendBroadcastAccepted() *SendBroadcastAccepted {
	return &SendBroadcastAccepted{}
}

/*SendBroadcastAccepted handles this case with default header values.

Notification
*/
type SendBroadcastAccepted struct {
	Payload *models.NotificationV2Notification
}

func (o *SendBroadcastAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/broadcast][%d] sendBroadcastAccepted  %+v", 202, o.Payload)
}

func (o *SendBroadcastAccepted) GetPayload() *models.NotificationV2Notification {
	return o.Payload
}

func (o *SendBroadcastAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationV2Notification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendBroadcastForbidden creates a SendBroadcastForbidden with default headers values
func NewSendBroadcastForbidden() *SendBroadcastForbidden {
	return &SendBroadcastForbidden{}
}

/*SendBroadcastForbidden handles this case with default header values.

Access denied
*/
type SendBroadcastForbidden struct {
}

func (o *SendBroadcastForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/broadcast][%d] sendBroadcastForbidden ", 403)
}

func (o *SendBroadcastForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
