package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetTeamPermissionsParams creates a new GetTeamPermissionsParams object
// with the default values initialized.
func NewGetTeamPermissionsParams() *GetTeamPermissionsParams {
	var ()
	return &GetTeamPermissionsParams{}
}

/*GetTeamPermissionsParams contains all the parameters to send to the API endpoint
for the get team permissions operation typically these are written to a http.Request
*/
type GetTeamPermissionsParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Teamname
	  The name of the team to which the permission applies

	*/
	Teamname string
}

// WithRepository adds the repository to the get team permissions params
func (o *GetTeamPermissionsParams) WithRepository(repository string) *GetTeamPermissionsParams {
	o.Repository = repository
	return o
}

// WithTeamname adds the teamname to the get team permissions params
func (o *GetTeamPermissionsParams) WithTeamname(teamname string) *GetTeamPermissionsParams {
	o.Teamname = teamname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamPermissionsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param teamname
	if err := r.SetPathParam("teamname", o.Teamname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
