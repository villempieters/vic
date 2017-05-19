package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VolumeStoresListResponse volume stores list response
// swagger:model VolumeStoresListResponse
type VolumeStoresListResponse struct {

	// stores
	// Required: true
	Stores []string `json:"Stores"`
}

// Validate validates this volume stores list response
func (m *VolumeStoresListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStores(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeStoresListResponse) validateStores(formats strfmt.Registry) error {

	if err := validate.Required("Stores", "body", m.Stores); err != nil {
		return err
	}

	return nil
}
