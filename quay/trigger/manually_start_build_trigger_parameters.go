package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

// NewManuallyStartBuildTriggerParams creates a new ManuallyStartBuildTriggerParams object
// with the default values initialized.
func NewManuallyStartBuildTriggerParams() *ManuallyStartBuildTriggerParams {
	var ()
	return &ManuallyStartBuildTriggerParams{}
}

/*ManuallyStartBuildTriggerParams contains all the parameters to send to the API endpoint
for the manually start build trigger operation typically these are written to a http.Request
*/
type ManuallyStartBuildTriggerParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.RunParameters
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*TriggerUUID
	  The UUID of the build trigger

	*/
	TriggerUUID string
}

// WithBody adds the body to the manually start build trigger params
func (o *ManuallyStartBuildTriggerParams) WithBody(body *models.RunParameters) *ManuallyStartBuildTriggerParams {
	o.Body = body
	return o
}

// WithRepository adds the repository to the manually start build trigger params
func (o *ManuallyStartBuildTriggerParams) WithRepository(repository string) *ManuallyStartBuildTriggerParams {
	o.Repository = repository
	return o
}

// WithTriggerUUID adds the triggerUuid to the manually start build trigger params
func (o *ManuallyStartBuildTriggerParams) WithTriggerUUID(triggerUuid string) *ManuallyStartBuildTriggerParams {
	o.TriggerUUID = triggerUuid
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ManuallyStartBuildTriggerParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.RunParameters)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param trigger_uuid
	if err := r.SetPathParam("trigger_uuid", o.TriggerUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
