package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*GeneralError GeneralError general error

swagger:model GeneralError
*/
type GeneralError struct {

	/* Message message
	 */
	Message *string `json:"message,omitempty"`
}

// Validate validates this general error
func (m *GeneralError) Validate(formats strfmt.Registry) error {
	return nil
}
