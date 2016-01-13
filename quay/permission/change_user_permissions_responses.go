package permission

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

type ChangeUserPermissionsReader struct {
	formats strfmt.Registry
}

func (o *ChangeUserPermissionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeUserPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewChangeUserPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewChangeUserPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewChangeUserPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewChangeUserPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeUserPermissionsOK creates a ChangeUserPermissionsOK with default headers values
func NewChangeUserPermissionsOK() *ChangeUserPermissionsOK {
	return &ChangeUserPermissionsOK{}
}

/*ChangeUserPermissionsOK

Successful invocation
*/
type ChangeUserPermissionsOK struct {
}

func (o *ChangeUserPermissionsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/permissions/user/{username}][%d] changeUserPermissionsOK ", 200)
}

func (o *ChangeUserPermissionsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeUserPermissionsBadRequest creates a ChangeUserPermissionsBadRequest with default headers values
func NewChangeUserPermissionsBadRequest() *ChangeUserPermissionsBadRequest {
	return &ChangeUserPermissionsBadRequest{}
}

/*ChangeUserPermissionsBadRequest

Bad Request
*/
type ChangeUserPermissionsBadRequest struct {
	Payload *models.GeneralError
}

func (o *ChangeUserPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/permissions/user/{username}][%d] changeUserPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *ChangeUserPermissionsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeUserPermissionsUnauthorized creates a ChangeUserPermissionsUnauthorized with default headers values
func NewChangeUserPermissionsUnauthorized() *ChangeUserPermissionsUnauthorized {
	return &ChangeUserPermissionsUnauthorized{}
}

/*ChangeUserPermissionsUnauthorized

Session required
*/
type ChangeUserPermissionsUnauthorized struct {
}

func (o *ChangeUserPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/permissions/user/{username}][%d] changeUserPermissionsUnauthorized ", 401)
}

func (o *ChangeUserPermissionsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeUserPermissionsForbidden creates a ChangeUserPermissionsForbidden with default headers values
func NewChangeUserPermissionsForbidden() *ChangeUserPermissionsForbidden {
	return &ChangeUserPermissionsForbidden{}
}

/*ChangeUserPermissionsForbidden

Unauthorized access
*/
type ChangeUserPermissionsForbidden struct {
}

func (o *ChangeUserPermissionsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/permissions/user/{username}][%d] changeUserPermissionsForbidden ", 403)
}

func (o *ChangeUserPermissionsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeUserPermissionsNotFound creates a ChangeUserPermissionsNotFound with default headers values
func NewChangeUserPermissionsNotFound() *ChangeUserPermissionsNotFound {
	return &ChangeUserPermissionsNotFound{}
}

/*ChangeUserPermissionsNotFound

Not found
*/
type ChangeUserPermissionsNotFound struct {
}

func (o *ChangeUserPermissionsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/permissions/user/{username}][%d] changeUserPermissionsNotFound ", 404)
}

func (o *ChangeUserPermissionsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
