package organization

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

type RemoveOrganizationMemberReader struct {
	formats strfmt.Registry
}

func (o *RemoveOrganizationMemberReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRemoveOrganizationMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveOrganizationMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveOrganizationMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveOrganizationMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveOrganizationMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveOrganizationMemberNoContent creates a RemoveOrganizationMemberNoContent with default headers values
func NewRemoveOrganizationMemberNoContent() *RemoveOrganizationMemberNoContent {
	return &RemoveOrganizationMemberNoContent{}
}

/*RemoveOrganizationMemberNoContent

Deleted
*/
type RemoveOrganizationMemberNoContent struct {
}

func (o *RemoveOrganizationMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/members/{membername}][%d] removeOrganizationMemberNoContent ", 204)
}

func (o *RemoveOrganizationMemberNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveOrganizationMemberBadRequest creates a RemoveOrganizationMemberBadRequest with default headers values
func NewRemoveOrganizationMemberBadRequest() *RemoveOrganizationMemberBadRequest {
	return &RemoveOrganizationMemberBadRequest{}
}

/*RemoveOrganizationMemberBadRequest

Bad Request
*/
type RemoveOrganizationMemberBadRequest struct {
	Payload *models.GeneralError
}

func (o *RemoveOrganizationMemberBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/members/{membername}][%d] removeOrganizationMemberBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveOrganizationMemberBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrganizationMemberUnauthorized creates a RemoveOrganizationMemberUnauthorized with default headers values
func NewRemoveOrganizationMemberUnauthorized() *RemoveOrganizationMemberUnauthorized {
	return &RemoveOrganizationMemberUnauthorized{}
}

/*RemoveOrganizationMemberUnauthorized

Session required
*/
type RemoveOrganizationMemberUnauthorized struct {
}

func (o *RemoveOrganizationMemberUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/members/{membername}][%d] removeOrganizationMemberUnauthorized ", 401)
}

func (o *RemoveOrganizationMemberUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveOrganizationMemberForbidden creates a RemoveOrganizationMemberForbidden with default headers values
func NewRemoveOrganizationMemberForbidden() *RemoveOrganizationMemberForbidden {
	return &RemoveOrganizationMemberForbidden{}
}

/*RemoveOrganizationMemberForbidden

Unauthorized access
*/
type RemoveOrganizationMemberForbidden struct {
}

func (o *RemoveOrganizationMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/members/{membername}][%d] removeOrganizationMemberForbidden ", 403)
}

func (o *RemoveOrganizationMemberForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveOrganizationMemberNotFound creates a RemoveOrganizationMemberNotFound with default headers values
func NewRemoveOrganizationMemberNotFound() *RemoveOrganizationMemberNotFound {
	return &RemoveOrganizationMemberNotFound{}
}

/*RemoveOrganizationMemberNotFound

Not found
*/
type RemoveOrganizationMemberNotFound struct {
}

func (o *RemoveOrganizationMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/members/{membername}][%d] removeOrganizationMemberNotFound ", 404)
}

func (o *RemoveOrganizationMemberNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
