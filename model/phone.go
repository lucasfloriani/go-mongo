package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/lucasfloriani/brazilian-ozzo-validation"
)

type Phone struct {
	Number string `json:"number"`
}

func (p *Phone) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(
			&p.Number,
			validation.Required.Error("Número de telefone não fornecido."),
			isbr.Phone.Error("Formato do número é inválido."),
		),
	)
}
