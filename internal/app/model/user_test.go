package model_test

import (
	"github.com/stretchr/testify/assert"
	"go_test_learning/internal/app/model"
	"testing"
	"time"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		getUser func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			getUser: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			getUser: func() *model.User {
				user := model.TestUser(t)
				user.Email = ""
				return user
			},
			isValid: false,
		},
		{
			name: "invalid email",
			getUser: func() *model.User {
				user := model.TestUser(t)
				user.Email = "some.invalid.email"
				return user
			},
			isValid: false,
		},
		{
			name: "younger 14 y.o.",
			getUser: func() *model.User {
				user := model.TestUser(t)
				user.BirthDate = time.Now()
				return user
			},
			isValid: false,
		},

		{
			name: "valid gender Male",
			getUser: func() *model.User {
				user := model.TestUser(t)
				user.Gender = "Male"
				return user
			},
			isValid: true,
		},
		{
			name: "valid gender Female",
			getUser: func() *model.User {
				user := model.TestUser(t)
				user.Gender = "Female"
				return user
			},
			isValid: true,
		},
		{
			name: "invalid gender",
			getUser: func() *model.User {
				user := model.TestUser(t)
				user.Gender = "Not sure"
				return user
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		user := tc.getUser()
		err := user.Validate()
		if tc.isValid {
			assert.NoError(t, err, tc.name)
		} else {
			assert.Error(t, err, tc.name)
		}
	}
}
