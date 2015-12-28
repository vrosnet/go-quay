package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type GetOrganizationMemberReader struct {
	formats strfmt.Registry
}

func (o *GetOrganizationMemberReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result GetOrganizationMemberOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result GetOrganizationMemberBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationMemberBadRequest", &result, response.Code())

	case 401:
		var result GetOrganizationMemberUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationMemberUnauthorized", &result, response.Code())

	case 403:
		var result GetOrganizationMemberForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationMemberForbidden", &result, response.Code())

	case 404:
		var result GetOrganizationMemberNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationMemberNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*
Successful invocation
*/
type GetOrganizationMemberOK struct {
}

func (o *GetOrganizationMemberOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type GetOrganizationMemberBadRequest struct {
	Payload *models.GeneralError
}

func (o *GetOrganizationMemberBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

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
type GetOrganizationMemberUnauthorized struct {
}

func (o *GetOrganizationMemberUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type GetOrganizationMemberForbidden struct {
}

func (o *GetOrganizationMemberForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type GetOrganizationMemberNotFound struct {
}

func (o *GetOrganizationMemberNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
