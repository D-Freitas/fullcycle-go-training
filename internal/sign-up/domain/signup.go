package domain

import (
	"errors"
	"server/internal/sign-up/domain/notification"
	validator "server/internal/sign-up/domain/validator"

	"golang.org/x/crypto/bcrypt"
)

type SignUp struct {
	ID                   string
	User                 string
	FullName             string
	Email                string
	PhoneNumber          string
	Password             string
	PasswordConfirmation string
}

func NewSignUp(
	id string,
	user string,
	fullname string,
	email string,
	phoneNumber string,
	password string,
	passwordConfirmation string,
) (*SignUp, error) {
	signUp := &SignUp{
		ID:                   id,
		User:                 user,
		FullName:             fullname,
		Email:                email,
		PhoneNumber:          phoneNumber,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
	}
	err := signUp.IsValid()
	if err != nil {
		return nil, err
	}
	return signUp, nil
}

func (r *SignUp) IsValid() error {
	notification := notification.NewSignUpNotification()
	if r.ID == "" {
		notification.AddError("invalid id")
	}
	if r.User == "" {
		notification.AddError("invalid user")
	}
	if r.FullName == "" {
		notification.AddError("invalid fullname")
	}
	if !validator.EmailValidator(r.Email) {
		notification.AddError("invalid email")
	}
	if !validator.PhoneNumberValidator(r.PhoneNumber) {
		notification.AddError("invalid phoneNumber")
	}
	if !validator.PasswordValidator(r.Password) {
		notification.AddError("invalid password")
	}
	if !validator.PasswordValidator(r.PasswordConfirmation) {
		notification.AddError("invalid passwordConfirmation")
	}
	if r.Password != r.PasswordConfirmation {
		notification.AddError("mismatched password")
	}
	if notification.HasErrors() {
		return errors.New(notification.Messages())
	}
	return nil
}

func (r *SignUp) EncryptPassword() error {
	password := []byte(r.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	r.Password = string(hashedPassword)
	r.PasswordConfirmation = string(hashedPassword)
	err = r.IsValid()
	if err != nil {
		return err
	}
	return nil
}
