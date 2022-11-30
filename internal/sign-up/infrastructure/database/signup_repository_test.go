package database

import (
	"database/sql"
	"fmt"
	"server/internal/sign-up/domain"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type SignUpRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

var dbsql *sql.DB

func (suite *SignUpRepositoryTestSuite) SetupSuite() {
	var (
		host     = "127.0.0.1"
		port     = 5432
		user     = "admin"
		password = "admin"
		dbname   = "postgres"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	suite.NoError(err)
	db.Exec(`
	CREATE TABLE users (
		id varchar(255) NOT NULL,
		username varchar(20) UNIQUE NOT NULL,
		fullname varchar(80) UNIQUE NOT NULL,
		email varchar(255) UNIQUE NOT NULL,
		phoneNumber varchar(17) NOT NULL,
		password varchar(255) NOT NULL
	)
	`)
	suite.Db = db
	dbsql = db
}

func (suite *SignUpRepositoryTestSuite) TearDownTest() {
	dbsql.Exec("drop table users")
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SignUpRepositoryTestSuite))
}

func (suite *SignUpRepositoryTestSuite) TestGivenARegistration_WhenSave_ThenShouldSaveOrder() {
	users, err := domain.NewSignUp(
		"123",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	)
	suite.NoError(err)
	repo := NewSignUpRepository(suite.Db)
	err = repo.Save(users)
	suite.NoError(err)
	var usersResult domain.SignUp
	err = suite.Db.QueryRow("SELECT id, username, fullname, email, phoneNumber, password FROM users WHERE id = $1", users.ID).
		Scan(&usersResult.ID, &usersResult.User, &usersResult.FullName, &usersResult.Email, &usersResult.PhoneNumber, &usersResult.Password)
	suite.NoError(err)
	suite.Equal(users.ID, usersResult.ID)
	suite.Equal(users.User, usersResult.User)
	suite.Equal(users.FullName, usersResult.FullName)
	suite.Equal(users.Email, usersResult.Email)
	suite.Equal(users.PhoneNumber, usersResult.PhoneNumber)
	suite.Equal(users.Password, usersResult.Password)
}
