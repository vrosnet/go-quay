package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
CreateUserRobotParams contains all the parameters to send to the API endpoint
for the create user robot operation typically these are written to a http.Request
*/
type CreateUserRobotParams struct {
	/*
	  The short name for the robot, without any user or organization prefix
	*/
	RobotShortname string
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserRobotParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param robot_shortname
	if err := r.SetPathParam("robot_shortname", o.RobotShortname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
