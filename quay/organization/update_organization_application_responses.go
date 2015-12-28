package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type UpdateOrganizationApplicationReader struct {
	formats strfmt.Registry
}

func (o *UpdateOrganizationApplicationReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result UpdateOrganizationApplicationOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result UpdateOrganizationApplicationBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("updateOrganizationApplicationBadRequest", &result, response.Code())

	case 401:
		var result UpdateOrganizationApplicationUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("updateOrganizationApplicationUnauthorized", &result, response.Code())

	case 403:
		var result UpdateOrganizationApplicationForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("updateOrganizationApplicationForbidden", &result, response.Code())

	case 404:
		var result UpdateOrganizationApplicationNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("updateOrganizationApplicationNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*
Successful invocation
*/
type UpdateOrganizationApplicationOK struct {
}

func (o *UpdateOrganizationApplicationOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type UpdateOrganizationApplicationBadRequest struct {
	Payload *models.GeneralError
}

func (o *UpdateOrganizationApplicationBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*
Session required
*/
type UpdateOrganizationApplicationUnauthorized struct {
}

func (o *UpdateOrganizationApplicationUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type UpdateOrganizationApplicationForbidden struct {
}

func (o *UpdateOrganizationApplicationForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type UpdateOrganizationApplicationNotFound struct {
}

func (o *UpdateOrganizationApplicationNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
