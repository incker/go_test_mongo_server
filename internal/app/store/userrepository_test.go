package store_test

import (
	"github.com/stretchr/testify/assert"
	"go_test_learning/internal/app/model"
	"go_test_learning/internal/app/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestDB(t, mongoDBURL)
	defer teardown("user")
	r := s.User()
	user, err := r.Create(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestDB(t, mongoDBURL)
	defer teardown("user")
	r := s.User()

	email := "user@example.org"

	user, err := r.FindByEmail(email)
	assert.Error(t, err)

	user, err = r.Create(&model.User{
		Email: email,
	})
	user, err = r.FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
