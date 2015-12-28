package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*
GetMatchingEntitiesParams contains all the parameters to send to the API endpoint
for the get matching entities operation typically these are written to a http.Request
*/
type GetMatchingEntitiesParams struct {
	/*
	  Whether to include orgs names.
	*/
	IncludeOrgs bool
	/*
	  Whether to include team names.
	*/
	IncludeTeams bool
	/*
	  Namespace to use when querying for org entities.
	*/
	Namespace string

	Prefix string
}

// WriteToRequest writes these params to a swagger request
func (o *GetMatchingEntitiesParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// query param includeOrgs
	if err := r.SetQueryParam("includeOrgs", swag.FormatBool(o.IncludeOrgs)); err != nil {
		return err
	}

	// query param includeTeams
	if err := r.SetQueryParam("includeTeams", swag.FormatBool(o.IncludeTeams)); err != nil {
		return err
	}

	// query param namespace
	if err := r.SetQueryParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param prefix
	if err := r.SetPathParam("prefix", o.Prefix); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
