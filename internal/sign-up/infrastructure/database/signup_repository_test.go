package database

import (
	"database/sql"
	"server/internal/sign-up/domain"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type SignUpRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *SignUpRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE users (id varchar(255) NOT NULL, user varchar(10) NOT NULL, fullname varchar(80) NOT NULL, email varchar(255) NOT NULL, phoneNumber varchar(17) NOT NULL, password varchar(255) NOT NULL)")
	suite.Db = db
}

func (suite *SignUpRepositoryTestSuite) TearDownTest() {
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
	err = suite.Db.QueryRow("SELECT id, user, fullname, email, phoneNumber, password FROM users WHERE id = ?", users.ID).
		Scan(&usersResult.ID, &usersResult.User, &usersResult.FullName, &usersResult.Email, &usersResult.PhoneNumber, &usersResult.Password)
	suite.NoError(err)
	suite.Equal(users.ID, usersResult.ID)
	suite.Equal(users.User, usersResult.User)
	suite.Equal(users.FullName, usersResult.FullName)
	suite.Equal(users.Email, usersResult.Email)
	suite.Equal(users.PhoneNumber, usersResult.PhoneNumber)
	suite.Equal(users.Password, usersResult.Password)
}
