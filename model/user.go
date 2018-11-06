package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// User represents an user record.
type User struct {
	ID      objectid.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string            `json:"name,omitempty"`
	Age     uint              `json:"age,omitempty"`
	Address Address           `json:"address"`
	Phones  []Phone           `json:"phones"`
	Courses []Course          `json:"courses"`
}

// NewUser creates a new User
func NewUser() *User {
	return &User{}
}

// Validate validates the User fields
func (u User) Validate() error {
	if err := u.Address.Validate(); err != nil {
		return err
	}

	for _, phone := range u.Phones {
		if err := phone.Validate(); err != nil {
			return err
		}
	}

	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Name,
			validation.Required.Error("Nome vazio."),
			validation.Length(5, 50).Error("Nome deve estar entre 5 à 50 caracteres"),
		),
		validation.Field(
			&u.Age,
			validation.Required.Error("Idade vazia."),
			validation.Min(uint(18)).Error("Idade mínima de 18 anos."),
		),
		validation.Field(
			&u.Phones,
			validation.Required.Error("É necessário pelo menos um telefone de contato."),
			validation.Length(1, 0).Error("É necessário pelo menos um telefone de contato."),
		),
	)
}
