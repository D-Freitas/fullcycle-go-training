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
		"TestPassword@123",
		"password_confirmation_test",
	}
	assert.EqualError(t, registration.IsValid(), "invalid id")
}

func TestGivenAnEmptyUser_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"",
		"fullname_test",
		"email_test",
		"phonenumber_test",
		"TestPassword@123",
		"password_confirmation_test",
	}
	assert.EqualError(t, registration.IsValid(), "invalid user")
}

func TestGivenAnEmptyFullname_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"",
		"email_test",
		"phonenumber_test",
		"TestPassword@123",
		"password_confirmation_test",
	}
	assert.EqualError(t, registration.IsValid(), "invalid fullname")
}

func TestGivenAnEmptyEmail_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"",
		"phonenumber_test",
		"TestPassword@123",
		"password_confirmation_test",
	}
	assert.EqualError(t, registration.IsValid(), "invalid email")
}

func TestGivenAnEmptyPhoneNumber_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"",
		"TestPassword@123",
		"password_confirmation_test",
	}
	assert.EqualError(t, registration.IsValid(), "invalid phoneNumber")
}

func TestGivenAnEmptyPassword_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"",
		"password_confirmation_test",
	}
	assert.EqualError(t, registration.IsValid(), "invalid password")
}

func TestGivenAnEmptyPasswordConfirmation_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"",
	}
	assert.EqualError(t, registration.IsValid(), "invalid passwordConfirmation")
}

func TestGivenAMismatchedPassword_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@1234",
	}
	assert.EqualError(t, registration.IsValid(), "mismatched password")
}

func TestGivenAValidParams_WhenICallNewRegistration_ThenIShouldReceiveCreateRegistrationWithAllParams(t *testing.T) {
	registration := Registration{
		"123",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.Equal(t, "123", registration.ID)
	assert.Equal(t, "user_test", registration.User)
	assert.Equal(t, "fullname_test", registration.FullName)
	assert.Equal(t, "test@gmail.com", registration.Email)
	assert.Equal(t, "11944431351", registration.PhoneNumber)
	assert.Equal(t, "TestPassword@123", registration.Password)
	assert.Equal(t, "TestPassword@123", registration.PasswordConfirmation)
	assert.Nil(t, registration.IsValid())
}

func TestGivenAValidParams_WhenICallNewRegistrationFunc_ThenIShouldReceiveCreateRegistrationWithAllParams(t *testing.T) {
	registration, err := NewRegistration(
		"123",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	)
	assert.Nil(t, err)
	assert.Equal(t, "123", registration.ID)
	assert.Equal(t, "user_test", registration.User)
	assert.Equal(t, "fullname_test", registration.FullName)
	assert.Equal(t, "test@gmail.com", registration.Email)
	assert.Equal(t, "11944431351", registration.PhoneNumber)
	assert.Equal(t, "TestPassword@123", registration.Password)
	assert.Equal(t, "TestPassword@123", registration.PasswordConfirmation)
}
