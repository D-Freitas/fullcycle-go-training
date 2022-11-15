package usecase

import (
	"database/sql"
	"server/internal/register-user/entity"
	"server/internal/register-user/infra/database"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type EncryptPasswordUseCaseTestSuite struct {
	suite.Suite
	RegistrationRepository database.RegistrationRepository
	Db                     *sql.DB
}

func (suite *EncryptPasswordUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE registrations (id varchar(255) NOT NULL, user varchar(10) NOT NULL, fullname varchar(80) NOT NULL, email varchar(255) NOT NULL, phoneNumber varchar(17) NOT NULL, password varchar(255) NOT NULL)")
	suite.Db = db
	suite.RegistrationRepository = *database.NewRegistrationRepository(db)
}

func (suite *EncryptPasswordUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EncryptPasswordUseCaseTestSuite))
}

func (suite *EncryptPasswordUseCaseTestSuite) TestEncryptPassword() {
	registration, err := entity.NewRegistration(
		"123",
		"user_test",
		"fullname_test",
		"email_test",
		"phonenumber_test",
		"password_test",
		"password_test",
	)
	suite.NoError(err)
	registration.EncryptPassword()
	encryptPasswordInput := RegistrationInputDTO{
		ID:                   registration.ID,
		User:                 registration.User,
		FullName:             registration.FullName,
		Email:                registration.Email,
		PhoneNumber:          registration.PhoneNumber,
		Password:             registration.Password,
		PasswordConfirmation: registration.PasswordConfirmation,
	}
	encryptPasswordUseCase := NewEncryptPasswordUseCase(suite.RegistrationRepository)
	output, err := encryptPasswordUseCase.Execute(encryptPasswordInput)
	suite.NoError(err)
	suite.Equal(registration.ID, output.ID)
	suite.Equal(registration.User, output.User)
	suite.Equal(registration.FullName, output.FullName)
	suite.Equal(registration.Email, output.Email)
	suite.Equal(registration.PhoneNumber, output.PhoneNumber)
}
