package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// JWorkspace j workspace
// swagger:model JWorkspace
type JWorkspace struct {

	// channel Id
	ChannelID string `json:"channelId,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// layout
	Layout interface{} `json:"layout,omitempty"`

	// machine label
	MachineLabel string `json:"machineLabel,omitempty"`

	// machine UId
	MachineUID string `json:"machineUId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// origin Id
	OriginID string `json:"originId,omitempty"`

	// root path
	RootPath string `json:"rootPath,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this j workspace
func (m *JWorkspace) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
