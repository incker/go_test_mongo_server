package model

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	LastName  string             `bson:"last_name,omitempty" json:"lastName,omitempty"`
	Country   string             `bson:"country,omitempty" json:"country,omitempty"`
	City      string             `bson:"city,omitempty" json:"city,omitempty"`
	Gender    string             `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate time.Time          `bson:"birth_date,omitempty" json:"birthDate,omitempty"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		// Email cannot be empty, and is Email
		validation.Field(&u.Email, validation.Required, is.Email),
		// Gender can be empty, or "Male" or "Female"
		validation.Field(&u.Gender, validation.By(func(value interface{}) error {
			if gender, ok := value.(string); ok {
				if gender != "Male" && gender != "Female" {
					return errors.New("gender has to be 'Male' or 'Female'")
				}
			}
			return nil
		})),
		// Age can not be under 14 y.o.
		validation.Field(&u.BirthDate, validation.By(func(value interface{}) error {
			if birthDate, ok := value.(time.Time); ok {
				if birthDate.Add(14 * 365 * 24 * time.Hour).After(time.Now()) {
					return errors.New("sorry, member can not be under 14 years old")
				}
			}
			return nil
		})),
	)
}
