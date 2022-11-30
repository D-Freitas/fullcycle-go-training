package application

import (
	"database/sql"
	"fmt"
	"server/internal/sign-up/domain"
	"server/internal/sign-up/infrastructure/database"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type RegisterUserUseCaseTestSuite struct {
	suite.Suite
	SignUpRepository database.SignUpRepository
	Db               *sql.DB
}

var dbsql *sql.DB

func (suite *RegisterUserUseCaseTestSuite) SetupSuite() {
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
	suite.SignUpRepository = *database.NewSignUpRepository(db)
	dbsql = db
}

func (suite *RegisterUserUseCaseTestSuite) TearDownTest() {
	dbsql.Exec("drop table users")
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RegisterUserUseCaseTestSuite))
}

func (suite *RegisterUserUseCaseTestSuite) TestRegisterUserUseCase() {
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
	registerUserUseCaseInput := UserInputDTO{
		ID:                   signUp.ID,
		User:                 signUp.User,
		FullName:             signUp.FullName,
		Email:                signUp.Email,
		PhoneNumber:          signUp.PhoneNumber,
		Password:             signUp.Password,
		PasswordConfirmation: signUp.PasswordConfirmation,
	}
	registerUserUseCase := NewRegisterUserUseCase(suite.SignUpRepository)
	output, err := registerUserUseCase.Execute(registerUserUseCaseInput)
	suite.NoError(err)
	suite.Equal(signUp.ID, output.ID)
	suite.Equal(signUp.User, output.User)
	suite.Equal(signUp.FullName, output.FullName)
	suite.Equal(signUp.Email, output.Email)
	suite.Equal(signUp.PhoneNumber, output.PhoneNumber)
}
