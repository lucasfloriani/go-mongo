package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gopkg.in/mgo.v2/bson"
)

type Course struct {
	ID   bson.ObjectId
	Name string
	Link string
}

func (c *Course) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.Name,
			validation.Required.Error("Nome do curso não foi adicionado."),
			validation.Length(5, 50).Error("Nome do curso deve estar entre 5 à 50 caracteres"),
		),
		validation.Field(
			c.Link,
			validation.Required.Error("URL vazia."),
			is.URL.Error("URL inválida."),
		),
	)
}
