package database

import (
	"database/sql"
	"server/internal/sign-up/domain"
)

type SignUpRepository struct {
	Db *sql.DB
}

func NewSignUpRepository(db *sql.DB) *SignUpRepository {
	return &SignUpRepository{Db: db}
}

func (r *SignUpRepository) Save(data *domain.SignUp) error {
	stmt, err := r.Db.Prepare("INSERT INTO users (id, username, fullname, email, phoneNumber, password) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.ID, data.User, data.FullName, data.Email, data.PhoneNumber, data.Password)
	if err != nil {
		return err
	}
	return nil
}
