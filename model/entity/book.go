package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type Book struct {
	ID        int       `json:"id" gorm:"primaryKey;type:serial"`
	NameBook  string    `json:"name_book" gorm:"type:varchar(255);unique"`
	Author    string    `json:"author" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp without time zone"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp without time zone"`
}

func (b Book) Validation() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.NameBook, validation.Required.Error("required")),
		validation.Field(&b.Author, validation.Required.Error("required")),
	)
}
