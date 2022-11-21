package entity

import (
	"errors"
	"server/internal/register-user/entity/notification"
	patternvalidator "server/internal/register-user/entity/pattern-validator"

	"golang.org/x/crypto/bcrypt"
)

type Registration struct {
	ID                   string
	User                 string
	FullName             string
	Email                string
	PhoneNumber          string
	Password             string
	PasswordConfirmation string
}

func NewRegistration(
	id string,
	user string,
	fullname string,
	email string,
	phoneNumber string,
	password string,
	passwordConfirmation string,
) (*Registration, error) {
	registration := &Registration{
		ID:                   id,
		User:                 user,
		FullName:             fullname,
		Email:                email,
		PhoneNumber:          phoneNumber,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
	}
	err := registration.IsValid()
	if err != nil {
		return nil, err
	}
	return registration, nil
}

func (r *Registration) IsValid() error {
	notification := notification.NewRegisterNotification()
	if r.ID == "" {
		notification.AddError("invalid id")
	}
	if r.User == "" {
		notification.AddError("invalid user")
	}
	if r.FullName == "" {
		notification.AddError("invalid fullname")
	}
	if !patternvalidator.EmailValidator(r.Email) {
		notification.AddError("invalid email")
	}
	if !patternvalidator.PhoneNumberValidator(r.PhoneNumber) {
		notification.AddError("invalid phoneNumber")
	}
	if !patternvalidator.PasswordValidator(r.Password) {
		notification.AddError("invalid password")
	}
	if !patternvalidator.PasswordValidator(r.PasswordConfirmation) {
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

func (r *Registration) EncryptPassword() error {
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
