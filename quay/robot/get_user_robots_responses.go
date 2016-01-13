package robot

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

type GetUserRobotsReader struct {
	formats strfmt.Registry
}

func (o *GetUserRobotsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserRobotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUserRobotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetUserRobotsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUserRobotsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUserRobotsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserRobotsOK creates a GetUserRobotsOK with default headers values
func NewGetUserRobotsOK() *GetUserRobotsOK {
	return &GetUserRobotsOK{}
}

/*GetUserRobotsOK

Successful invocation
*/
type GetUserRobotsOK struct {
}

func (o *GetUserRobotsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/robots][%d] getUserRobotsOK ", 200)
}

func (o *GetUserRobotsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserRobotsBadRequest creates a GetUserRobotsBadRequest with default headers values
func NewGetUserRobotsBadRequest() *GetUserRobotsBadRequest {
	return &GetUserRobotsBadRequest{}
}

/*GetUserRobotsBadRequest

Bad Request
*/
type GetUserRobotsBadRequest struct {
	Payload *models.GeneralError
}

func (o *GetUserRobotsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/robots][%d] getUserRobotsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserRobotsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRobotsUnauthorized creates a GetUserRobotsUnauthorized with default headers values
func NewGetUserRobotsUnauthorized() *GetUserRobotsUnauthorized {
	return &GetUserRobotsUnauthorized{}
}

/*GetUserRobotsUnauthorized

Session required
*/
type GetUserRobotsUnauthorized struct {
}

func (o *GetUserRobotsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/robots][%d] getUserRobotsUnauthorized ", 401)
}

func (o *GetUserRobotsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserRobotsForbidden creates a GetUserRobotsForbidden with default headers values
func NewGetUserRobotsForbidden() *GetUserRobotsForbidden {
	return &GetUserRobotsForbidden{}
}

/*GetUserRobotsForbidden

Unauthorized access
*/
type GetUserRobotsForbidden struct {
}

func (o *GetUserRobotsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/robots][%d] getUserRobotsForbidden ", 403)
}

func (o *GetUserRobotsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserRobotsNotFound creates a GetUserRobotsNotFound with default headers values
func NewGetUserRobotsNotFound() *GetUserRobotsNotFound {
	return &GetUserRobotsNotFound{}
}

/*GetUserRobotsNotFound

Not found
*/
type GetUserRobotsNotFound struct {
}

func (o *GetUserRobotsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/robots][%d] getUserRobotsNotFound ", 404)
}

func (o *GetUserRobotsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
