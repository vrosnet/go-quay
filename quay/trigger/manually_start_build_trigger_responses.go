package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type ManuallyStartBuildTriggerReader struct {
	formats strfmt.Registry
}

func (o *ManuallyStartBuildTriggerReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewManuallyStartBuildTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewManuallyStartBuildTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewManuallyStartBuildTriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewManuallyStartBuildTriggerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewManuallyStartBuildTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewManuallyStartBuildTriggerOK creates a ManuallyStartBuildTriggerOK with default headers values
func NewManuallyStartBuildTriggerOK() *ManuallyStartBuildTriggerOK {
	return &ManuallyStartBuildTriggerOK{}
}

/*ManuallyStartBuildTriggerOK

Successful invocation
*/
type ManuallyStartBuildTriggerOK struct {
}

func (o *ManuallyStartBuildTriggerOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/start][%d] manuallyStartBuildTriggerOK ", 200)
}

func (o *ManuallyStartBuildTriggerOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewManuallyStartBuildTriggerBadRequest creates a ManuallyStartBuildTriggerBadRequest with default headers values
func NewManuallyStartBuildTriggerBadRequest() *ManuallyStartBuildTriggerBadRequest {
	return &ManuallyStartBuildTriggerBadRequest{}
}

/*ManuallyStartBuildTriggerBadRequest

Bad Request
*/
type ManuallyStartBuildTriggerBadRequest struct {
	Payload *models.GeneralError
}

func (o *ManuallyStartBuildTriggerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/start][%d] manuallyStartBuildTriggerBadRequest  %+v", 400, o.Payload)
}

func (o *ManuallyStartBuildTriggerBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManuallyStartBuildTriggerUnauthorized creates a ManuallyStartBuildTriggerUnauthorized with default headers values
func NewManuallyStartBuildTriggerUnauthorized() *ManuallyStartBuildTriggerUnauthorized {
	return &ManuallyStartBuildTriggerUnauthorized{}
}

/*ManuallyStartBuildTriggerUnauthorized

Session required
*/
type ManuallyStartBuildTriggerUnauthorized struct {
}

func (o *ManuallyStartBuildTriggerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/start][%d] manuallyStartBuildTriggerUnauthorized ", 401)
}

func (o *ManuallyStartBuildTriggerUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewManuallyStartBuildTriggerForbidden creates a ManuallyStartBuildTriggerForbidden with default headers values
func NewManuallyStartBuildTriggerForbidden() *ManuallyStartBuildTriggerForbidden {
	return &ManuallyStartBuildTriggerForbidden{}
}

/*ManuallyStartBuildTriggerForbidden

Unauthorized access
*/
type ManuallyStartBuildTriggerForbidden struct {
}

func (o *ManuallyStartBuildTriggerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/start][%d] manuallyStartBuildTriggerForbidden ", 403)
}

func (o *ManuallyStartBuildTriggerForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewManuallyStartBuildTriggerNotFound creates a ManuallyStartBuildTriggerNotFound with default headers values
func NewManuallyStartBuildTriggerNotFound() *ManuallyStartBuildTriggerNotFound {
	return &ManuallyStartBuildTriggerNotFound{}
}

/*ManuallyStartBuildTriggerNotFound

Not found
*/
type ManuallyStartBuildTriggerNotFound struct {
}

func (o *ManuallyStartBuildTriggerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/start][%d] manuallyStartBuildTriggerNotFound ", 404)
}

func (o *ManuallyStartBuildTriggerNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
