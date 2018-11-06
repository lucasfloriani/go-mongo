package model

import (
	"github.com/go-ozzo/ozzo-validation"
)

// Address represents an address record.
type Address struct {
	Name string `json:"name"`
}

// Validate validates the Address fields
func (a *Address) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Name, validation.Required.Error("Endereço vazio.")),
	)
}
