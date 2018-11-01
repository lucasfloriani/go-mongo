package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	ID   bson.ObjectId
	Name string
}

func (a *Address) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Name, validation.Required.Error("Endereço vazio.")),
	)
}
