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

type ListRepoTeamPermissionsReader struct {
	formats strfmt.Registry
}

func (o *ListRepoTeamPermissionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRepoTeamPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListRepoTeamPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListRepoTeamPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListRepoTeamPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListRepoTeamPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRepoTeamPermissionsOK creates a ListRepoTeamPermissionsOK with default headers values
func NewListRepoTeamPermissionsOK() *ListRepoTeamPermissionsOK {
	return &ListRepoTeamPermissionsOK{}
}

/*ListRepoTeamPermissionsOK

Successful invocation
*/
type ListRepoTeamPermissionsOK struct {
}

func (o *ListRepoTeamPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/team/][%d] listRepoTeamPermissionsOK ", 200)
}

func (o *ListRepoTeamPermissionsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRepoTeamPermissionsBadRequest creates a ListRepoTeamPermissionsBadRequest with default headers values
func NewListRepoTeamPermissionsBadRequest() *ListRepoTeamPermissionsBadRequest {
	return &ListRepoTeamPermissionsBadRequest{}
}

/*ListRepoTeamPermissionsBadRequest

Bad Request
*/
type ListRepoTeamPermissionsBadRequest struct {
	Payload *models.GeneralError
}

func (o *ListRepoTeamPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/team/][%d] listRepoTeamPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *ListRepoTeamPermissionsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoTeamPermissionsUnauthorized creates a ListRepoTeamPermissionsUnauthorized with default headers values
func NewListRepoTeamPermissionsUnauthorized() *ListRepoTeamPermissionsUnauthorized {
	return &ListRepoTeamPermissionsUnauthorized{}
}

/*ListRepoTeamPermissionsUnauthorized

Session required
*/
type ListRepoTeamPermissionsUnauthorized struct {
}

func (o *ListRepoTeamPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/team/][%d] listRepoTeamPermissionsUnauthorized ", 401)
}

func (o *ListRepoTeamPermissionsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRepoTeamPermissionsForbidden creates a ListRepoTeamPermissionsForbidden with default headers values
func NewListRepoTeamPermissionsForbidden() *ListRepoTeamPermissionsForbidden {
	return &ListRepoTeamPermissionsForbidden{}
}

/*ListRepoTeamPermissionsForbidden

Unauthorized access
*/
type ListRepoTeamPermissionsForbidden struct {
}

func (o *ListRepoTeamPermissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/team/][%d] listRepoTeamPermissionsForbidden ", 403)
}

func (o *ListRepoTeamPermissionsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRepoTeamPermissionsNotFound creates a ListRepoTeamPermissionsNotFound with default headers values
func NewListRepoTeamPermissionsNotFound() *ListRepoTeamPermissionsNotFound {
	return &ListRepoTeamPermissionsNotFound{}
}

/*ListRepoTeamPermissionsNotFound

Not found
*/
type ListRepoTeamPermissionsNotFound struct {
}

func (o *ListRepoTeamPermissionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/team/][%d] listRepoTeamPermissionsNotFound ", 404)
}

func (o *ListRepoTeamPermissionsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
