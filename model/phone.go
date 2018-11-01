package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/lucasfloriani/brazilian-ozzo-validation"
	"gopkg.in/mgo.v2/bson"
)

type Phone struct {
	ID     bson.ObjectId
	Number string
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
