package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		ID:        primitive.ObjectID{},
		Email:     "user@example.org",
		LastName:  "Smith",
		Country:   "German",
		City:      "Berlin",
		Gender:    "Male",
		BirthDate: time.Time{},
	}
}
