package database

import (
	"database/sql"
	"server/internal/register-user/entity"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type RegistrationRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *RegistrationRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE registrations (id varchar(255) NOT NULL, user varchar(10) NOT NULL, fullname varchar(80) NOT NULL, email varchar(255) NOT NULL, phoneNumber varchar(17) NOT NULL, password varchar(255) NOT NULL)")
	suite.Db = db
}

func (suite *RegistrationRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RegistrationRepositoryTestSuite))
}

func (suite *RegistrationRepositoryTestSuite) TestGivenARegistration_WhenSave_ThenShouldSaveOrder() {
	registration, err := entity.NewRegistration(
		"123",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	)
	suite.NoError(err)
	repo := NewRegistrationRepository(suite.Db)
	err = repo.Save(registration)
	suite.NoError(err)
	var registrationResult entity.Registration
	err = suite.Db.QueryRow("SELECT id, user, fullname, email, phoneNumber, password FROM registrations WHERE id = ?", registration.ID).
		Scan(&registrationResult.ID, &registrationResult.User, &registrationResult.FullName, &registrationResult.Email, &registrationResult.PhoneNumber, &registrationResult.Password)
	suite.NoError(err)
	suite.Equal(registration.ID, registrationResult.ID)
	suite.Equal(registration.User, registrationResult.User)
	suite.Equal(registration.FullName, registrationResult.FullName)
	suite.Equal(registration.Email, registrationResult.Email)
	suite.Equal(registration.PhoneNumber, registrationResult.PhoneNumber)
	suite.Equal(registration.Password, registrationResult.Password)
}
