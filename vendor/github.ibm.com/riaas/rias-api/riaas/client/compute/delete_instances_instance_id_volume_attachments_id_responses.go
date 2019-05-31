// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// DeleteInstancesInstanceIDVolumeAttachmentsIDReader is a Reader for the DeleteInstancesInstanceIDVolumeAttachmentsID structure.
type DeleteInstancesInstanceIDVolumeAttachmentsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteInstancesInstanceIDVolumeAttachmentsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteInstancesInstanceIDVolumeAttachmentsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteInstancesInstanceIDVolumeAttachmentsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewDeleteInstancesInstanceIDVolumeAttachmentsIDNoContent creates a DeleteInstancesInstanceIDVolumeAttachmentsIDNoContent with default headers values
func NewDeleteInstancesInstanceIDVolumeAttachmentsIDNoContent() *DeleteInstancesInstanceIDVolumeAttachmentsIDNoContent {
	return &DeleteInstancesInstanceIDVolumeAttachmentsIDNoContent{}
}

/*DeleteInstancesInstanceIDVolumeAttachmentsIDNoContent handles this case with default header values.

error
*/
type DeleteInstancesInstanceIDVolumeAttachmentsIDNoContent struct {
	Payload *models.Riaaserror
}

func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /instances/{instance_id}/volume_attachments/{id}][%d] deleteInstancesInstanceIdVolumeAttachmentsIdNoContent  %+v", 204, o.Payload)
}

func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstancesInstanceIDVolumeAttachmentsIDNotFound creates a DeleteInstancesInstanceIDVolumeAttachmentsIDNotFound with default headers values
func NewDeleteInstancesInstanceIDVolumeAttachmentsIDNotFound() *DeleteInstancesInstanceIDVolumeAttachmentsIDNotFound {
	return &DeleteInstancesInstanceIDVolumeAttachmentsIDNotFound{}
}

/*DeleteInstancesInstanceIDVolumeAttachmentsIDNotFound handles this case with default header values.

error
*/
type DeleteInstancesInstanceIDVolumeAttachmentsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /instances/{instance_id}/volume_attachments/{id}][%d] deleteInstancesInstanceIdVolumeAttachmentsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstancesInstanceIDVolumeAttachmentsIDDefault creates a DeleteInstancesInstanceIDVolumeAttachmentsIDDefault with default headers values
func NewDeleteInstancesInstanceIDVolumeAttachmentsIDDefault(code int) *DeleteInstancesInstanceIDVolumeAttachmentsIDDefault {
	return &DeleteInstancesInstanceIDVolumeAttachmentsIDDefault{
		_statusCode: code,
	}
}

/*DeleteInstancesInstanceIDVolumeAttachmentsIDDefault handles this case with default header values.

unexpectederror
*/
type DeleteInstancesInstanceIDVolumeAttachmentsIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the delete instances instance ID volume attachments ID default response
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /instances/{instance_id}/volume_attachments/{id}][%d] DeleteInstancesInstanceIDVolumeAttachmentsID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}