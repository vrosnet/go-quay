package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetImageParams creates a new GetImageParams object
// with the default values initialized.
func NewGetImageParams() *GetImageParams {
	var ()
	return &GetImageParams{}
}

/*GetImageParams contains all the parameters to send to the API endpoint
for the get image operation typically these are written to a http.Request
*/
type GetImageParams struct {

	/*ImageID
	  The Docker image ID

	*/
	ImageID string
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
}

// WithImageID adds the imageId to the get image params
func (o *GetImageParams) WithImageID(imageId string) *GetImageParams {
	o.ImageID = imageId
	return o
}

// WithRepository adds the repository to the get image params
func (o *GetImageParams) WithRepository(repository string) *GetImageParams {
	o.Repository = repository
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param image_id
	if err := r.SetPathParam("image_id", o.ImageID); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
