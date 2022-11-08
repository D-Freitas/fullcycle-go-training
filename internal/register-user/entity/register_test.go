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

func TestGivenAnEmptyFullname_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"",
		"email_test",
		"phonenumber_test",
		"password_test",
		"password_confirmation_test",
	}
	assert.Error(t, registration.IsValid(), "invalid fullname")
}

func TestGivenAnEmptyEmail_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"",
		"phonenumber_test",
		"password_test",
		"password_confirmation_test",
	}
	assert.Error(t, registration.IsValid(), "invalid email")
}

func TestGivenAnEmptyPhoneNumber_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"email_test",
		"",
		"password_test",
		"password_confirmation_test",
	}
	assert.Error(t, registration.IsValid(), "invalid phoneNumber")
}

func TestGivenAnEmptyPassword_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"email_test",
		"phonenumber_test",
		"",
		"password_confirmation_test",
	}
	assert.Error(t, registration.IsValid(), "invalid phoneNumber")
}

func TestGivenAnEmptyPasswordConfirmation_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"email_test",
		"phonenumber_test",
		"password_test",
		"",
	}
	assert.Error(t, registration.IsValid(), "invalid passwordConfirmation")
}

func TestGivenAMismatchedPassword_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"email_test",
		"phonenumber_test",
		"password_test",
		"mismatched_password",
	}
	assert.Error(t, registration.IsValid(), "mismatched password")
}

func TestGivenAValidParams_WhenICallNewRegistration_ThenIShouldReceiveCreateRegistrationWithAllParams(t *testing.T) {
	registration := Registration{
		"123",
		"user_test",
		"fullname_test",
		"email_test",
		"phonenumber_test",
		"password_test",
		"password_test",
	}
	assert.Equal(t, "123", registration.ID)
	assert.Equal(t, "user_test", registration.User)
	assert.Equal(t, "fullname_test", registration.FullName)
	assert.Equal(t, "email_test", registration.Email)
	assert.Equal(t, "phonenumber_test", registration.PhoneNumber)
	assert.Equal(t, "password_test", registration.Password)
	assert.Equal(t, "password_test", registration.PasswordConfirmation)
	assert.Nil(t, registration.IsValid())
}

func TestGivenAValidParams_WhenICallNewRegistrationFunc_ThenIShouldReceiveCreateRegistrationWithAllParams(t *testing.T) {
	registration, err := NewRegistration(
		"123",
		"user_test",
		"fullname_test",
		"email_test",
		"phonenumber_test",
		"password_test",
		"password_test",
	)
	assert.Nil(t, err)
	assert.Equal(t, "123", registration.ID)
	assert.Equal(t, "user_test", registration.User)
	assert.Equal(t, "fullname_test", registration.FullName)
	assert.Equal(t, "email_test", registration.Email)
	assert.Equal(t, "phonenumber_test", registration.PhoneNumber)
	assert.Equal(t, "password_test", registration.Password)
	assert.Equal(t, "password_test", registration.PasswordConfirmation)
}
