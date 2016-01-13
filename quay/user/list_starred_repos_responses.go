package user

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

type ListStarredReposReader struct {
	formats strfmt.Registry
}

func (o *ListStarredReposReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListStarredReposOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListStarredReposBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListStarredReposUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListStarredReposForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListStarredReposNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewListStarredReposOK creates a ListStarredReposOK with default headers values
func NewListStarredReposOK() *ListStarredReposOK {
	return &ListStarredReposOK{}
}

/*ListStarredReposOK

Successful invocation
*/
type ListStarredReposOK struct {
}

func (o *ListStarredReposOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/starred][%d] listStarredReposOK ", 200)
}

func (o *ListStarredReposOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListStarredReposBadRequest creates a ListStarredReposBadRequest with default headers values
func NewListStarredReposBadRequest() *ListStarredReposBadRequest {
	return &ListStarredReposBadRequest{}
}

/*ListStarredReposBadRequest

Bad Request
*/
type ListStarredReposBadRequest struct {
	Payload *models.GeneralError
}

func (o *ListStarredReposBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/starred][%d] listStarredReposBadRequest  %+v", 400, o.Payload)
}

func (o *ListStarredReposBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListStarredReposUnauthorized creates a ListStarredReposUnauthorized with default headers values
func NewListStarredReposUnauthorized() *ListStarredReposUnauthorized {
	return &ListStarredReposUnauthorized{}
}

/*ListStarredReposUnauthorized

Session required
*/
type ListStarredReposUnauthorized struct {
}

func (o *ListStarredReposUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/starred][%d] listStarredReposUnauthorized ", 401)
}

func (o *ListStarredReposUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListStarredReposForbidden creates a ListStarredReposForbidden with default headers values
func NewListStarredReposForbidden() *ListStarredReposForbidden {
	return &ListStarredReposForbidden{}
}

/*ListStarredReposForbidden

Unauthorized access
*/
type ListStarredReposForbidden struct {
}

func (o *ListStarredReposForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/starred][%d] listStarredReposForbidden ", 403)
}

func (o *ListStarredReposForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListStarredReposNotFound creates a ListStarredReposNotFound with default headers values
func NewListStarredReposNotFound() *ListStarredReposNotFound {
	return &ListStarredReposNotFound{}
}

/*ListStarredReposNotFound

Not found
*/
type ListStarredReposNotFound struct {
}

func (o *ListStarredReposNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/starred][%d] listStarredReposNotFound ", 404)
}

func (o *ListStarredReposNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
