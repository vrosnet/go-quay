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

type DeleteTeamPermissionsReader struct {
	formats strfmt.Registry
}

func (o *DeleteTeamPermissionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteTeamPermissionsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteTeamPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteTeamPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteTeamPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteTeamPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTeamPermissionsNoContent creates a DeleteTeamPermissionsNoContent with default headers values
func NewDeleteTeamPermissionsNoContent() *DeleteTeamPermissionsNoContent {
	return &DeleteTeamPermissionsNoContent{}
}

/*DeleteTeamPermissionsNoContent

Deleted
*/
type DeleteTeamPermissionsNoContent struct {
}

func (o *DeleteTeamPermissionsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/team/{teamname}][%d] deleteTeamPermissionsNoContent ", 204)
}

func (o *DeleteTeamPermissionsNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTeamPermissionsBadRequest creates a DeleteTeamPermissionsBadRequest with default headers values
func NewDeleteTeamPermissionsBadRequest() *DeleteTeamPermissionsBadRequest {
	return &DeleteTeamPermissionsBadRequest{}
}

/*DeleteTeamPermissionsBadRequest

Bad Request
*/
type DeleteTeamPermissionsBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteTeamPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/team/{teamname}][%d] deleteTeamPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTeamPermissionsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamPermissionsUnauthorized creates a DeleteTeamPermissionsUnauthorized with default headers values
func NewDeleteTeamPermissionsUnauthorized() *DeleteTeamPermissionsUnauthorized {
	return &DeleteTeamPermissionsUnauthorized{}
}

/*DeleteTeamPermissionsUnauthorized

Session required
*/
type DeleteTeamPermissionsUnauthorized struct {
}

func (o *DeleteTeamPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/team/{teamname}][%d] deleteTeamPermissionsUnauthorized ", 401)
}

func (o *DeleteTeamPermissionsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTeamPermissionsForbidden creates a DeleteTeamPermissionsForbidden with default headers values
func NewDeleteTeamPermissionsForbidden() *DeleteTeamPermissionsForbidden {
	return &DeleteTeamPermissionsForbidden{}
}

/*DeleteTeamPermissionsForbidden

Unauthorized access
*/
type DeleteTeamPermissionsForbidden struct {
}

func (o *DeleteTeamPermissionsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/team/{teamname}][%d] deleteTeamPermissionsForbidden ", 403)
}

func (o *DeleteTeamPermissionsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTeamPermissionsNotFound creates a DeleteTeamPermissionsNotFound with default headers values
func NewDeleteTeamPermissionsNotFound() *DeleteTeamPermissionsNotFound {
	return &DeleteTeamPermissionsNotFound{}
}

/*DeleteTeamPermissionsNotFound

Not found
*/
type DeleteTeamPermissionsNotFound struct {
}

func (o *DeleteTeamPermissionsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/team/{teamname}][%d] deleteTeamPermissionsNotFound ", 404)
}

func (o *DeleteTeamPermissionsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
