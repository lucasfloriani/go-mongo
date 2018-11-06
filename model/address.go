package model

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Address struct {
	Name string `json:"name"`
}

func (a *Address) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Name, validation.Required.Error("Endereço vazio.")),
	)
}
