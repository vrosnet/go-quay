package secscan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/coreos/go-quay/models"
)

// GetRepoImageSecurityReader is a Reader for the GetRepoImageSecurity structure.
type GetRepoImageSecurityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetRepoImageSecurityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRepoImageSecurityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRepoImageSecurityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRepoImageSecurityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRepoImageSecurityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRepoImageSecurityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepoImageSecurityOK creates a GetRepoImageSecurityOK with default headers values
func NewGetRepoImageSecurityOK() *GetRepoImageSecurityOK {
	return &GetRepoImageSecurityOK{}
}

/*GetRepoImageSecurityOK handles this case with default header values.

Successful invocation
*/
type GetRepoImageSecurityOK struct {
}

func (o *GetRepoImageSecurityOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/image/{imageid}/security][%d] getRepoImageSecurityOK ", 200)
}

func (o *GetRepoImageSecurityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepoImageSecurityBadRequest creates a GetRepoImageSecurityBadRequest with default headers values
func NewGetRepoImageSecurityBadRequest() *GetRepoImageSecurityBadRequest {
	return &GetRepoImageSecurityBadRequest{}
}

/*GetRepoImageSecurityBadRequest handles this case with default header values.

Bad Request
*/
type GetRepoImageSecurityBadRequest struct {
	Payload *models.APIError
}

func (o *GetRepoImageSecurityBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/image/{imageid}/security][%d] getRepoImageSecurityBadRequest  %+v", 400, o.Payload)
}

func (o *GetRepoImageSecurityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoImageSecurityUnauthorized creates a GetRepoImageSecurityUnauthorized with default headers values
func NewGetRepoImageSecurityUnauthorized() *GetRepoImageSecurityUnauthorized {
	return &GetRepoImageSecurityUnauthorized{}
}

/*GetRepoImageSecurityUnauthorized handles this case with default header values.

Session required
*/
type GetRepoImageSecurityUnauthorized struct {
	Payload *models.APIError
}

func (o *GetRepoImageSecurityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/image/{imageid}/security][%d] getRepoImageSecurityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepoImageSecurityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoImageSecurityForbidden creates a GetRepoImageSecurityForbidden with default headers values
func NewGetRepoImageSecurityForbidden() *GetRepoImageSecurityForbidden {
	return &GetRepoImageSecurityForbidden{}
}

/*GetRepoImageSecurityForbidden handles this case with default header values.

Unauthorized access
*/
type GetRepoImageSecurityForbidden struct {
	Payload *models.APIError
}

func (o *GetRepoImageSecurityForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/image/{imageid}/security][%d] getRepoImageSecurityForbidden  %+v", 403, o.Payload)
}

func (o *GetRepoImageSecurityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoImageSecurityNotFound creates a GetRepoImageSecurityNotFound with default headers values
func NewGetRepoImageSecurityNotFound() *GetRepoImageSecurityNotFound {
	return &GetRepoImageSecurityNotFound{}
}

/*GetRepoImageSecurityNotFound handles this case with default header values.

Not found
*/
type GetRepoImageSecurityNotFound struct {
	Payload *models.APIError
}

func (o *GetRepoImageSecurityNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/image/{imageid}/security][%d] getRepoImageSecurityNotFound  %+v", 404, o.Payload)
}

func (o *GetRepoImageSecurityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}