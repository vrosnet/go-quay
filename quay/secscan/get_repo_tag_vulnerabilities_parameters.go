package secscan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
GetRepoTagVulnerabilitiesParams contains all the parameters to send to the API endpoint
for the get repo tag vulnerabilities operation typically these are written to a http.Request
*/
type GetRepoTagVulnerabilitiesParams struct {
	/*
	  Minimum vulnerability priority
	*/
	MinimumPriority string
	/*
	  The full path of the repository. e.g. namespace/name
	*/
	Repository string
	/*
	  The name of the tag
	*/
	Tag string
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoTagVulnerabilitiesParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// query param minimumPriority
	if err := r.SetQueryParam("minimumPriority", o.MinimumPriority); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
