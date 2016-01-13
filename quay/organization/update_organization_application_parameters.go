package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

// NewUpdateOrganizationApplicationParams creates a new UpdateOrganizationApplicationParams object
// with the default values initialized.
func NewUpdateOrganizationApplicationParams() *UpdateOrganizationApplicationParams {
	var ()
	return &UpdateOrganizationApplicationParams{}
}

/*UpdateOrganizationApplicationParams contains all the parameters to send to the API endpoint
for the update organization application operation typically these are written to a http.Request
*/
type UpdateOrganizationApplicationParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.UpdateApp
	/*ClientID
	  The OAuth client ID

	*/
	ClientID string
	/*Orgname
	  The name of the organization

	*/
	Orgname string
}

// WithBody adds the body to the update organization application params
func (o *UpdateOrganizationApplicationParams) WithBody(body *models.UpdateApp) *UpdateOrganizationApplicationParams {
	o.Body = body
	return o
}

// WithClientID adds the clientId to the update organization application params
func (o *UpdateOrganizationApplicationParams) WithClientID(clientId string) *UpdateOrganizationApplicationParams {
	o.ClientID = clientId
	return o
}

// WithOrgname adds the orgname to the update organization application params
func (o *UpdateOrganizationApplicationParams) WithOrgname(orgname string) *UpdateOrganizationApplicationParams {
	o.Orgname = orgname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationApplicationParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.UpdateApp)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param client_id
	if err := r.SetPathParam("client_id", o.ClientID); err != nil {
		return err
	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
