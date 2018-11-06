package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Course represents an course record.
type Course struct {
	ID   objectid.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string            `json:"name,omitempty"`
	Link string            `json:"link,omitempty"`
}

// NewCourse creates a new Course
func NewCourse() *Course {
	return &Course{}
}

// Validate validates the Course fields
func (c Course) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(
			&c.Name,
			validation.Required.Error("Nome do curso não foi adicionado."),
			validation.Length(5, 50).Error("Nome do curso deve estar entre 5 à 50 caracteres"),
		),
		validation.Field(
			&c.Link,
			validation.Required.Error("URL vazia."),
			is.URL.Error("URL inválida."),
		),
	)
}
