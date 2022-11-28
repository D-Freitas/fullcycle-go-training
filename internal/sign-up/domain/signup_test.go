package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateNewSignUp_ThenShouldReceiveAnError(t *testing.T) {
	signUp := SignUp{
		"",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, signUp.IsValid(), "signup: invalid id")
}

func TestGivenAnEmptyUser_WhenCreateNewSignUp_ThenShouldReceiveAnError(t *testing.T) {
	signUp := SignUp{
		"id_test",
		"",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, signUp.IsValid(), "signup: invalid user")
}

func TestGivenAnEmptyFullname_WhenCreateNewSignUp_ThenShouldReceiveAnError(t *testing.T) {
	signUp := SignUp{
		"id_test",
		"user_test",
		"",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, signUp.IsValid(), "signup: invalid fullname")
}

func TestGivenAnEmptyEmail_WhenCreateNewSignUp_ThenShouldReceiveAnError(t *testing.T) {
	signUp := SignUp{
		"id_test",
		"user_test",
		"fullname_test",
		"",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, signUp.IsValid(), "signup: invalid email")
}

func TestGivenAnEmptyPhoneNumber_WhenCreateNewSignUp_ThenShouldReceiveAnError(t *testing.T) {
	signUp := SignUp{
		"id_test",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, signUp.IsValid(), "signup: invalid phoneNumber")
}

func TestGivenAnEmptyPassword_WhenCreateNewSignUp_ThenShouldReceiveAnError(t *testing.T) {
	signUp := SignUp{
		"id_test",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"",
		"",
	}
	assert.EqualError(t, signUp.IsValid(), "signup: invalid password, invalid passwordConfirmation")
}

func TestGivenAnEmptyUserAndEmptyFullname_WhenCreateNewSignUp_ThenShouldReceiveAnError(t *testing.T) {
	signUp := SignUp{
		"id_test",
		"",
		"",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, signUp.IsValid(), "signup: invalid user, invalid fullname")
}

func TestGivenAnEmptyUserEmptyFullnameAndInvalidEmail_WhenCreateNewSignUp_ThenShouldReceiveAnError(t *testing.T) {
	signUp := SignUp{
		"id_test",
		"",
		"",
		"invalid_email",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.EqualError(t, signUp.IsValid(), "signup: invalid user, invalid fullname, invalid email")
}

func TestGivenAMismatchedPassword_WhenCreateNewSignUp_ThenShouldReceiveAnError(t *testing.T) {
	signUp := SignUp{
		"id_test",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@1234",
	}
	assert.EqualError(t, signUp.IsValid(), "signup: mismatched password")
}

func TestGivenAValidParams_WhenICallNewSignUp_ThenIShouldReceiveCreateSignUpWithAllParams(t *testing.T) {
	signUp := SignUp{
		"123",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	}
	assert.Equal(t, "123", signUp.ID)
	assert.Equal(t, "user_test", signUp.User)
	assert.Equal(t, "fullname_test", signUp.FullName)
	assert.Equal(t, "test@gmail.com", signUp.Email)
	assert.Equal(t, "11944431351", signUp.PhoneNumber)
	assert.Equal(t, "TestPassword@123", signUp.Password)
	assert.Equal(t, "TestPassword@123", signUp.PasswordConfirmation)
	assert.Nil(t, signUp.IsValid())
}

func TestGivenAValidParams_WhenICallNewSignUpFunc_ThenIShouldReceiveCreateSignUpWithAllParams(t *testing.T) {
	signUp, err := NewSignUp(
		"123",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	)
	assert.Nil(t, err)
	assert.Equal(t, "123", signUp.ID)
	assert.Equal(t, "user_test", signUp.User)
	assert.Equal(t, "fullname_test", signUp.FullName)
	assert.Equal(t, "test@gmail.com", signUp.Email)
	assert.Equal(t, "11944431351", signUp.PhoneNumber)
	assert.Equal(t, "TestPassword@123", signUp.Password)
	assert.Equal(t, "TestPassword@123", signUp.PasswordConfirmation)
}
