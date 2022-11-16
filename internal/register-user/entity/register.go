package entity

import (
	"errors"
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
	if r.ID == "" {
		return errors.New("invalid id")
	}
	if r.User == "" {
		return errors.New("invalid user")
	}
	if r.FullName == "" {
		return errors.New("invalid fullname")
	}
	if patternvalidator.EmailValidator(r.Email) {
		return errors.New("invalid email")
	}
	if patternvalidator.PhoneNumberValidator(r.PhoneNumber) {
		return errors.New("invalid phoneNumber")
	}
	if patternvalidator.PasswordValidator(r.Password) {
		return errors.New("invalid password")
	}
	if patternvalidator.PasswordValidator(r.Password) {
		return errors.New("invalid passwordConfirmation")
	}
	if r.Password != r.PasswordConfirmation {
		return errors.New("mismatched password")
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
