package usecase

import (
	"server/internal/sign-up/domain"
	"server/internal/sign-up/domain/contract"
	"server/internal/sign-up/infrastructure/database"
)

type UserInputDTO struct {
	ID                   string
	User                 string
	FullName             string
	Email                string
	PhoneNumber          string
	Password             string
	PasswordConfirmation string
}

type UserOutputDTO struct {
	ID          string
	User        string
	FullName    string
	Email       string
	PhoneNumber string
}

type EncryptPasswordUseCase struct {
	SignUpRepository contract.SignUpRepository
}

func NewEncryptPasswordUseCase(signUpRepository database.SignUpRepository) *EncryptPasswordUseCase {
	return &EncryptPasswordUseCase{
		SignUpRepository: &signUpRepository,
	}
}

func (e *EncryptPasswordUseCase) Execute(input UserInputDTO) (*UserOutputDTO, error) {
	user, err := domain.NewSignUp(input.ID, input.User, input.FullName, input.Email, input.PhoneNumber, input.Password, input.PasswordConfirmation)
	if err != nil {
		return nil, err
	}
	err = user.EncryptPassword()
	if err != nil {
		return nil, err
	}
	err = e.SignUpRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return &UserOutputDTO{
		ID:          user.ID,
		User:        user.User,
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}, nil
}
