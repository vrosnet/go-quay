package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewListStarredReposParams creates a new ListStarredReposParams object
// with the default values initialized.
func NewListStarredReposParams() *ListStarredReposParams {
	var ()
	return &ListStarredReposParams{}
}

/*ListStarredReposParams contains all the parameters to send to the API endpoint
for the list starred repos operation typically these are written to a http.Request
*/
type ListStarredReposParams struct {

	/*Limit
	  Limit on the number of results (int)

	*/
	Limit *int64
	/*Page
	  Offset page number. (int)

	*/
	Page *int64
}

// WithLimit adds the limit to the list starred repos params
func (o *ListStarredReposParams) WithLimit(limit *int64) *ListStarredReposParams {
	o.Limit = limit
	return o
}

// WithPage adds the page to the list starred repos params
func (o *ListStarredReposParams) WithPage(page *int64) *ListStarredReposParams {
	o.Page = page
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListStarredReposParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
