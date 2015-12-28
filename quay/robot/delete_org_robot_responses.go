package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type DeleteOrgRobotReader struct {
	formats strfmt.Registry
}

func (o *DeleteOrgRobotReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		var result DeleteOrgRobotNoContent
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result DeleteOrgRobotBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteOrgRobotBadRequest", &result, response.Code())

	case 401:
		var result DeleteOrgRobotUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteOrgRobotUnauthorized", &result, response.Code())

	case 403:
		var result DeleteOrgRobotForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteOrgRobotForbidden", &result, response.Code())

	case 404:
		var result DeleteOrgRobotNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteOrgRobotNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*
Deleted
*/
type DeleteOrgRobotNoContent struct {
}

func (o *DeleteOrgRobotNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type DeleteOrgRobotBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteOrgRobotBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

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
type DeleteOrgRobotUnauthorized struct {
}

func (o *DeleteOrgRobotUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type DeleteOrgRobotForbidden struct {
}

func (o *DeleteOrgRobotForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type DeleteOrgRobotNotFound struct {
}

func (o *DeleteOrgRobotNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
