package team

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

type DeleteOrganizationTeamReader struct {
	formats strfmt.Registry
}

func (o *DeleteOrganizationTeamReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteOrganizationTeamNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteOrganizationTeamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteOrganizationTeamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteOrganizationTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteOrganizationTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOrganizationTeamNoContent creates a DeleteOrganizationTeamNoContent with default headers values
func NewDeleteOrganizationTeamNoContent() *DeleteOrganizationTeamNoContent {
	return &DeleteOrganizationTeamNoContent{}
}

/*DeleteOrganizationTeamNoContent

Deleted
*/
type DeleteOrganizationTeamNoContent struct {
}

func (o *DeleteOrganizationTeamNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/team/{teamname}][%d] deleteOrganizationTeamNoContent ", 204)
}

func (o *DeleteOrganizationTeamNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationTeamBadRequest creates a DeleteOrganizationTeamBadRequest with default headers values
func NewDeleteOrganizationTeamBadRequest() *DeleteOrganizationTeamBadRequest {
	return &DeleteOrganizationTeamBadRequest{}
}

/*DeleteOrganizationTeamBadRequest

Bad Request
*/
type DeleteOrganizationTeamBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteOrganizationTeamBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/team/{teamname}][%d] deleteOrganizationTeamBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrganizationTeamBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationTeamUnauthorized creates a DeleteOrganizationTeamUnauthorized with default headers values
func NewDeleteOrganizationTeamUnauthorized() *DeleteOrganizationTeamUnauthorized {
	return &DeleteOrganizationTeamUnauthorized{}
}

/*DeleteOrganizationTeamUnauthorized

Session required
*/
type DeleteOrganizationTeamUnauthorized struct {
}

func (o *DeleteOrganizationTeamUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/team/{teamname}][%d] deleteOrganizationTeamUnauthorized ", 401)
}

func (o *DeleteOrganizationTeamUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationTeamForbidden creates a DeleteOrganizationTeamForbidden with default headers values
func NewDeleteOrganizationTeamForbidden() *DeleteOrganizationTeamForbidden {
	return &DeleteOrganizationTeamForbidden{}
}

/*DeleteOrganizationTeamForbidden

Unauthorized access
*/
type DeleteOrganizationTeamForbidden struct {
}

func (o *DeleteOrganizationTeamForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/team/{teamname}][%d] deleteOrganizationTeamForbidden ", 403)
}

func (o *DeleteOrganizationTeamForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationTeamNotFound creates a DeleteOrganizationTeamNotFound with default headers values
func NewDeleteOrganizationTeamNotFound() *DeleteOrganizationTeamNotFound {
	return &DeleteOrganizationTeamNotFound{}
}

/*DeleteOrganizationTeamNotFound

Not found
*/
type DeleteOrganizationTeamNotFound struct {
}

func (o *DeleteOrganizationTeamNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/team/{teamname}][%d] deleteOrganizationTeamNotFound ", 404)
}

func (o *DeleteOrganizationTeamNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
