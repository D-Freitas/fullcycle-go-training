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
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, registration.IsValid(), "register: invalid id")
}

func TestGivenAnEmptyUser_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, registration.IsValid(), "register: invalid user")
}

func TestGivenAnEmptyFullname_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, registration.IsValid(), "register: invalid fullname")
}

func TestGivenAnEmptyEmail_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, registration.IsValid(), "register: invalid email")
}

func TestGivenAnEmptyPhoneNumber_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, registration.IsValid(), "register: invalid phoneNumber")
}

func TestGivenAnEmptyPassword_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"",
		"",
	}
	assert.EqualError(t, registration.IsValid(), "register: invalid password, invalid passwordConfirmation")
}

func TestGivenAnEmptyUserAndEmptyFullname_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"",
		"",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, registration.IsValid(), "register: invalid user, invalid fullname")
}

func TestGivenAnEmptyUserEmptyFullnameAndInvalidEmail_WhenCreateNewRegistration_ThenShouldReceiveAnError(t *testing.T) {
	registration := Registration{
		"id_test",
		"",
		"",
		"invalid_email",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, registration.IsValid(), "register: invalid user, invalid fullname, invalid email")
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
	assert.EqualError(t, registration.IsValid(), "register: mismatched password")
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
