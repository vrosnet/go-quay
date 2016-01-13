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

type DeleteUserPermissionsReader struct {
	formats strfmt.Registry
}

func (o *DeleteUserPermissionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserPermissionsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteUserPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteUserPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserPermissionsNoContent creates a DeleteUserPermissionsNoContent with default headers values
func NewDeleteUserPermissionsNoContent() *DeleteUserPermissionsNoContent {
	return &DeleteUserPermissionsNoContent{}
}

/*DeleteUserPermissionsNoContent

Deleted
*/
type DeleteUserPermissionsNoContent struct {
}

func (o *DeleteUserPermissionsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/user/{username}][%d] deleteUserPermissionsNoContent ", 204)
}

func (o *DeleteUserPermissionsNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserPermissionsBadRequest creates a DeleteUserPermissionsBadRequest with default headers values
func NewDeleteUserPermissionsBadRequest() *DeleteUserPermissionsBadRequest {
	return &DeleteUserPermissionsBadRequest{}
}

/*DeleteUserPermissionsBadRequest

Bad Request
*/
type DeleteUserPermissionsBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteUserPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/user/{username}][%d] deleteUserPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserPermissionsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserPermissionsUnauthorized creates a DeleteUserPermissionsUnauthorized with default headers values
func NewDeleteUserPermissionsUnauthorized() *DeleteUserPermissionsUnauthorized {
	return &DeleteUserPermissionsUnauthorized{}
}

/*DeleteUserPermissionsUnauthorized

Session required
*/
type DeleteUserPermissionsUnauthorized struct {
}

func (o *DeleteUserPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/user/{username}][%d] deleteUserPermissionsUnauthorized ", 401)
}

func (o *DeleteUserPermissionsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserPermissionsForbidden creates a DeleteUserPermissionsForbidden with default headers values
func NewDeleteUserPermissionsForbidden() *DeleteUserPermissionsForbidden {
	return &DeleteUserPermissionsForbidden{}
}

/*DeleteUserPermissionsForbidden

Unauthorized access
*/
type DeleteUserPermissionsForbidden struct {
}

func (o *DeleteUserPermissionsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/user/{username}][%d] deleteUserPermissionsForbidden ", 403)
}

func (o *DeleteUserPermissionsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserPermissionsNotFound creates a DeleteUserPermissionsNotFound with default headers values
func NewDeleteUserPermissionsNotFound() *DeleteUserPermissionsNotFound {
	return &DeleteUserPermissionsNotFound{}
}

/*DeleteUserPermissionsNotFound

Not found
*/
type DeleteUserPermissionsNotFound struct {
}

func (o *DeleteUserPermissionsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/user/{username}][%d] deleteUserPermissionsNotFound ", 404)
}

func (o *DeleteUserPermissionsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
