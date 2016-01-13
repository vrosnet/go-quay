package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new permission API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for permission API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Update the existing team permission.
 */
func (a *Client) ChangeTeamPermissions(params *ChangeTeamPermissionsParams, authInfo client.AuthInfoWriter) (*ChangeTeamPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeTeamPermissionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "changeTeamPermissions",
		Method:      "PUT",
		PathPattern: "/api/v1/repository/{repository}/permissions/team/{teamname}",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &ChangeTeamPermissionsReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeTeamPermissionsOK), nil
}

/*Update the perimssions for an existing repository.
 */
func (a *Client) ChangeUserPermissions(params *ChangeUserPermissionsParams, authInfo client.AuthInfoWriter) (*ChangeUserPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeUserPermissionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "changeUserPermissions",
		Method:      "PUT",
		PathPattern: "/api/v1/repository/{repository}/permissions/user/{username}",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &ChangeUserPermissionsReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeUserPermissionsOK), nil
}

/*Delete the permission for the specified team.
 */
func (a *Client) DeleteTeamPermissions(params *DeleteTeamPermissionsParams, authInfo client.AuthInfoWriter) (*DeleteTeamPermissionsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTeamPermissionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "deleteTeamPermissions",
		Method:      "DELETE",
		PathPattern: "/api/v1/repository/{repository}/permissions/team/{teamname}",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &DeleteTeamPermissionsReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteTeamPermissionsNoContent), nil
}

/*Delete the permission for the user.
 */
func (a *Client) DeleteUserPermissions(params *DeleteUserPermissionsParams, authInfo client.AuthInfoWriter) (*DeleteUserPermissionsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserPermissionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "deleteUserPermissions",
		Method:      "DELETE",
		PathPattern: "/api/v1/repository/{repository}/permissions/user/{username}",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &DeleteUserPermissionsReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserPermissionsNoContent), nil
}

/*Fetch the permission for the specified team.
 */
func (a *Client) GetTeamPermissions(params *GetTeamPermissionsParams, authInfo client.AuthInfoWriter) (*GetTeamPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTeamPermissionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "getTeamPermissions",
		Method:      "GET",
		PathPattern: "/api/v1/repository/{repository}/permissions/team/{teamname}",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &GetTeamPermissionsReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTeamPermissionsOK), nil
}

/*Get the Fetch the permission for the specified user.
 */
func (a *Client) GetUserPermissions(params *GetUserPermissionsParams, authInfo client.AuthInfoWriter) (*GetUserPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserPermissionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "getUserPermissions",
		Method:      "GET",
		PathPattern: "/api/v1/repository/{repository}/permissions/user/{username}",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &GetUserPermissionsReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserPermissionsOK), nil
}

/*Get the fetch the permission for the specified user.
 */
func (a *Client) GetUserTransitivePermission(params *GetUserTransitivePermissionParams, authInfo client.AuthInfoWriter) (*GetUserTransitivePermissionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserTransitivePermissionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "getUserTransitivePermission",
		Method:      "GET",
		PathPattern: "/api/v1/repository/{repository}/permissions/user/{username}/transitive",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &GetUserTransitivePermissionReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserTransitivePermissionOK), nil
}

/*List all team permission.
 */
func (a *Client) ListRepoTeamPermissions(params *ListRepoTeamPermissionsParams, authInfo client.AuthInfoWriter) (*ListRepoTeamPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRepoTeamPermissionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "listRepoTeamPermissions",
		Method:      "GET",
		PathPattern: "/api/v1/repository/{repository}/permissions/team/",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &ListRepoTeamPermissionsReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRepoTeamPermissionsOK), nil
}

/*List all user permissions.
 */
func (a *Client) ListRepoUserPermissions(params *ListRepoUserPermissionsParams, authInfo client.AuthInfoWriter) (*ListRepoUserPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRepoUserPermissionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "listRepoUserPermissions",
		Method:      "GET",
		PathPattern: "/api/v1/repository/{repository}/permissions/user/",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &ListRepoUserPermissionsReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRepoUserPermissionsOK), nil
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
