package database

import (
	"database/sql"
	"server/internal/register-user/entity"
)

type RegistrationRepository struct {
	Db *sql.DB
}

func NewRegistrationRepository(db *sql.DB) *RegistrationRepository {
	return &RegistrationRepository{Db: db}
}

func (r *RegistrationRepository) Save(registration *entity.Registration) error {
	stmt, err := r.Db.Prepare("INSERT INTO registrations (id, user, fullname, email, phoneNumber, password, passwordConfirmation) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(registration.ID, registration.User, registration.FullName, registration.Email, registration.PhoneNumber, registration.Password, registration.PasswordConfirmation)
	if err != nil {
		return err
	}
	return nil
}
