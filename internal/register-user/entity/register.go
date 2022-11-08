package entity

import "errors"

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
	return &Registration{
		ID:                   id,
		User:                 user,
		FullName:             fullname,
		Email:                email,
		PhoneNumber:          phoneNumber,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
	}, nil
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
	if r.Email == "" {
		return errors.New("invalid email")
	}
	if r.PhoneNumber == "" {
		return errors.New("invalid phoneNumber")
	}
	if r.Password == "" {
		return errors.New("invalid password")
	}
	if r.PasswordConfirmation == "" {
		return errors.New("invalid passwordConfirmation")
	}
	if r.Password != r.PasswordConfirmation {
		return errors.New("mismatched password")
	}
	return nil
}
