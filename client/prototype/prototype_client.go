package prototype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new prototype API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for prototype API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Create a new permission prototype.
 */
func (a *Client) CreateOrganizationPrototypePermission(params CreateOrganizationPrototypePermissionParams, authInfo client.AuthInfoWriter) (*CreateOrganizationPrototypePermissionOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "createOrganizationPrototypePermission",
		Params:   &params,
		Reader:   &CreateOrganizationPrototypePermissionReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateOrganizationPrototypePermissionOK), nil
}

/*Delete an existing permission prototype.
 */
func (a *Client) DeleteOrganizationPrototypePermission(params DeleteOrganizationPrototypePermissionParams, authInfo client.AuthInfoWriter) (*DeleteOrganizationPrototypePermissionNoContent, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "deleteOrganizationPrototypePermission",
		Params:   &params,
		Reader:   &DeleteOrganizationPrototypePermissionReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrganizationPrototypePermissionNoContent), nil
}

/*List the existing prototypes for this organization.
 */
func (a *Client) GetOrganizationPrototypePermissions(params GetOrganizationPrototypePermissionsParams, authInfo client.AuthInfoWriter) (*GetOrganizationPrototypePermissionsOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "getOrganizationPrototypePermissions",
		Params:   &params,
		Reader:   &GetOrganizationPrototypePermissionsReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrganizationPrototypePermissionsOK), nil
}

/*Update the role of an existing permission prototype.
 */
func (a *Client) UpdateOrganizationPrototypePermission(params UpdateOrganizationPrototypePermissionParams, authInfo client.AuthInfoWriter) (*UpdateOrganizationPrototypePermissionOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "updateOrganizationPrototypePermission",
		Params:   &params,
		Reader:   &UpdateOrganizationPrototypePermissionReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateOrganizationPrototypePermissionOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
