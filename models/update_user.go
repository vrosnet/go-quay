package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
Fields which can be updated in a user.

swagger:model UpdateUser
*/
type UpdateUser struct {

	/* The user's email address
	 */
	Email string `json:"email,omitempty"`

	/* Whether the user desires to receive an invoice email.
	 */
	InvoiceEmail bool `json:"invoice_email,omitempty"`

	/* The user's password
	 */
	Password string `json:"password,omitempty"`

	/* TagExpiration tag expiration

	Maximum: 2.592e+06
	Minimum: 0
	*/
	TagExpiration int64 `json:"tag_expiration,omitempty"`

	/* The user's username
	 */
	Username string `json:"username,omitempty"`
}

// Validate validates this update user
func (m *UpdateUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTagExpiration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateUser) validateTagExpiration(formats strfmt.Registry) error {

	if err := validate.Minimum("tag_expiration", "body", float64(m.TagExpiration), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("tag_expiration", "body", float64(m.TagExpiration), 2.592e+06, false); err != nil {
		return err
	}

	return nil
}
