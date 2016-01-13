package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*RepositoryBuildRequest Description of a new repository build.

swagger:model RepositoryBuildRequest
*/
type RepositoryBuildRequest struct {

	/* The URL of the .tar.gz to build. Must start with "http" or "https".
	 */
	ArchiveURL *string `json:"archive_url,omitempty"`

	/* The tags to which the built images will be pushed. If none specified, "latest" is used.

	Min Items: 1
	Unique: true
	*/
	DockerTags []string `json:"docker_tags,omitempty"`

	/* The file id that was generated when the build spec was uploaded
	 */
	FileID *string `json:"file_id,omitempty"`

	/* Username of a Quay robot account to use as pull credentials
	 */
	PullRobot *string `json:"pull_robot,omitempty"`

	/* Subdirectory in which the Dockerfile can be found
	 */
	Subdirectory *string `json:"subdirectory,omitempty"`
}

// Validate validates this repository build request
func (m *RepositoryBuildRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDockerTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepositoryBuildRequest) validateDockerTags(formats strfmt.Registry) error {

	if swag.IsZero(m.DockerTags) { // not required
		return nil
	}

	iDockerTagsSize := int64(len(m.DockerTags))

	if err := validate.MinItems("docker_tags", "body", iDockerTagsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("docker_tags", "body", m.DockerTags); err != nil {
		return err
	}

	for i := 0; i < len(m.DockerTags); i++ {

		if err := validate.Required("docker_tags"+"."+strconv.Itoa(i), "body", string(m.DockerTags[i])); err != nil {
			return err
		}

	}

	return nil
}
