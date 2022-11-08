package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"",
		"user_test",
		"fullname_test",
		"email_test",
		"phonenumber_test",
		"password_test",
		"password_confirmation_test",
	}
	assert.Error(t, registration.IsValid(), "invalid id")
}

func TestGivenAnEmptyUser_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"",
		"fullname_test",
		"email_test",
		"phonenumber_test",
		"password_test",
		"password_confirmation_test",
	}
	assert.Error(t, registration.IsValid(), "invalid user")
}
