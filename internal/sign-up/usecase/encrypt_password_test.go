package usecase

import (
	"database/sql"
	"server/internal/sign-up/domain"
	"server/internal/sign-up/infrastructure/database"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type EncryptPasswordUseCaseTestSuite struct {
	suite.Suite
	SignUpRepository database.SignUpRepository
	Db               *sql.DB
}

func (suite *EncryptPasswordUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE users (id varchar(255) NOT NULL, user varchar(10) NOT NULL, fullname varchar(80) NOT NULL, email varchar(255) NOT NULL, phoneNumber varchar(17) NOT NULL, password varchar(255) NOT NULL)")
	suite.Db = db
	suite.SignUpRepository = *database.NewSignUpRepository(db)
}

func (suite *EncryptPasswordUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EncryptPasswordUseCaseTestSuite))
}

func (suite *EncryptPasswordUseCaseTestSuite) TestEncryptPassword() {
	signUp, err := domain.NewSignUp(
		"123",
		"user_test",
		"fullname_test",
		"test@gmail.com",
		"11944431351",
		"TestPassword@123",
		"TestPassword@123",
	)
	suite.NoError(err)
	err = signUp.EncryptPassword()
	suite.NoError(err)
	encryptPasswordInput := UserInputDTO{
		ID:                   signUp.ID,
		User:                 signUp.User,
		FullName:             signUp.FullName,
		Email:                signUp.Email,
		PhoneNumber:          signUp.PhoneNumber,
		Password:             signUp.Password,
		PasswordConfirmation: signUp.PasswordConfirmation,
	}
	encryptPasswordUseCase := NewEncryptPasswordUseCase(suite.SignUpRepository)
	output, err := encryptPasswordUseCase.Execute(encryptPasswordInput)
	suite.NoError(err)
	suite.Equal(signUp.ID, output.ID)
	suite.Equal(signUp.User, output.User)
	suite.Equal(signUp.FullName, output.FullName)
	suite.Equal(signUp.Email, output.Email)
	suite.Equal(signUp.PhoneNumber, output.PhoneNumber)
}
