package repository

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

type GetRepoReader struct {
	formats strfmt.Registry
}

func (o *GetRepoReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRepoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRepoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRepoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRepoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRepoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepoOK creates a GetRepoOK with default headers values
func NewGetRepoOK() *GetRepoOK {
	return &GetRepoOK{}
}

/*GetRepoOK

Successful invocation
*/
type GetRepoOK struct {
}

func (o *GetRepoOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}][%d] getRepoOK ", 200)
}

func (o *GetRepoOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepoBadRequest creates a GetRepoBadRequest with default headers values
func NewGetRepoBadRequest() *GetRepoBadRequest {
	return &GetRepoBadRequest{}
}

/*GetRepoBadRequest

Bad Request
*/
type GetRepoBadRequest struct {
	Payload *models.GeneralError
}

func (o *GetRepoBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}][%d] getRepoBadRequest  %+v", 400, o.Payload)
}

func (o *GetRepoBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoUnauthorized creates a GetRepoUnauthorized with default headers values
func NewGetRepoUnauthorized() *GetRepoUnauthorized {
	return &GetRepoUnauthorized{}
}

/*GetRepoUnauthorized

Session required
*/
type GetRepoUnauthorized struct {
}

func (o *GetRepoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}][%d] getRepoUnauthorized ", 401)
}

func (o *GetRepoUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepoForbidden creates a GetRepoForbidden with default headers values
func NewGetRepoForbidden() *GetRepoForbidden {
	return &GetRepoForbidden{}
}

/*GetRepoForbidden

Unauthorized access
*/
type GetRepoForbidden struct {
}

func (o *GetRepoForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}][%d] getRepoForbidden ", 403)
}

func (o *GetRepoForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepoNotFound creates a GetRepoNotFound with default headers values
func NewGetRepoNotFound() *GetRepoNotFound {
	return &GetRepoNotFound{}
}

/*GetRepoNotFound

Not found
*/
type GetRepoNotFound struct {
}

func (o *GetRepoNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}][%d] getRepoNotFound ", 404)
}

func (o *GetRepoNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
