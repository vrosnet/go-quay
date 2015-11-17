package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new robot API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for robot API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Create a new robot in the organization.
 */
func (a *Client) CreateOrgRobot(params CreateOrgRobotParams, authInfo client.AuthInfoWriter) (*CreateOrgRobotOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "createOrgRobot",
		Params:   &params,
		Reader:   &CreateOrgRobotReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateOrgRobotOK), nil
}

/*Create a new user robot with the specified name.
 */
func (a *Client) CreateUserRobot(params CreateUserRobotParams) (*CreateUserRobotOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "createUserRobot",
		Params: &params,
		Reader: &CreateUserRobotReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserRobotOK), nil
}

/*Delete an existing organization robot.
 */
func (a *Client) DeleteOrgRobot(params DeleteOrgRobotParams, authInfo client.AuthInfoWriter) (*DeleteOrgRobotNoContent, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "deleteOrgRobot",
		Params:   &params,
		Reader:   &DeleteOrgRobotReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrgRobotNoContent), nil
}

/*Delete an existing robot.
 */
func (a *Client) DeleteUserRobot(params DeleteUserRobotParams) (*DeleteUserRobotNoContent, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "deleteUserRobot",
		Params: &params,
		Reader: &DeleteUserRobotReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserRobotNoContent), nil
}

/*Returns the organization's robot with the specified name.
 */
func (a *Client) GetOrgRobot(params GetOrgRobotParams, authInfo client.AuthInfoWriter) (*GetOrgRobotOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "getOrgRobot",
		Params:   &params,
		Reader:   &GetOrgRobotReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrgRobotOK), nil
}

/*Returns the list of repository permissions for the org's robot.
 */
func (a *Client) GetOrgRobotPermissions(params GetOrgRobotPermissionsParams) (*GetOrgRobotPermissionsOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "getOrgRobotPermissions",
		Params: &params,
		Reader: &GetOrgRobotPermissionsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrgRobotPermissionsOK), nil
}

/*List the organization's robots.
 */
func (a *Client) GetOrgRobots(params GetOrgRobotsParams, authInfo client.AuthInfoWriter) (*GetOrgRobotsOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "getOrgRobots",
		Params:   &params,
		Reader:   &GetOrgRobotsReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrgRobotsOK), nil
}

/*Returns the user's robot with the specified name.
 */
func (a *Client) GetUserRobot(params GetUserRobotParams) (*GetUserRobotOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "getUserRobot",
		Params: &params,
		Reader: &GetUserRobotReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserRobotOK), nil
}

/*Returns the list of repository permissions for the user's robot.
 */
func (a *Client) GetUserRobotPermissions(params GetUserRobotPermissionsParams) (*GetUserRobotPermissionsOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "getUserRobotPermissions",
		Params: &params,
		Reader: &GetUserRobotPermissionsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserRobotPermissionsOK), nil
}

/*List the available robots for the user.
 */
func (a *Client) GetUserRobots(params GetUserRobotsParams) (*GetUserRobotsOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "getUserRobots",
		Params: &params,
		Reader: &GetUserRobotsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserRobotsOK), nil
}

/*Regenerates the token for an organization robot.
 */
func (a *Client) RegenerateOrgRobotToken(params RegenerateOrgRobotTokenParams, authInfo client.AuthInfoWriter) (*RegenerateOrgRobotTokenOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "regenerateOrgRobotToken",
		Params:   &params,
		Reader:   &RegenerateOrgRobotTokenReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegenerateOrgRobotTokenOK), nil
}

/*Regenerates the token for a user's robot.
 */
func (a *Client) RegenerateUserRobotToken(params RegenerateUserRobotTokenParams) (*RegenerateUserRobotTokenOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "regenerateUserRobotToken",
		Params: &params,
		Reader: &RegenerateUserRobotTokenReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegenerateUserRobotTokenOK), nil
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
