package usecase

import (
	"server/internal/register-user/entity"
	"server/internal/register-user/infra/database"
)

type RegistrationInputDTO struct {
	ID                   string
	User                 string
	FullName             string
	Email                string
	PhoneNumber          string
	Password             string
	PasswordConfirmation string
}

type RegistrationOutputDTO struct {
	ID          string
	User        string
	FullName    string
	Email       string
	PhoneNumber string
}

type EncryptPasswordUseCase struct {
	RegistrationRepository entity.RegistrationRepositoryInterface
}

func NewEncryptPasswordUseCase(registrationRepository database.RegistrationRepository) *EncryptPasswordUseCase {
	return &EncryptPasswordUseCase{
		RegistrationRepository: &registrationRepository,
	}
}

func (e *EncryptPasswordUseCase) Execute(input RegistrationInputDTO) (*RegistrationOutputDTO, error) {
	registration, err := entity.NewRegistration(input.ID, input.User, input.FullName, input.Email, input.PhoneNumber, input.Password, input.PasswordConfirmation)
	if err != nil {
		return nil, err
	}
	err = registration.EncryptPassword()
	if err != nil {
		return nil, err
	}
	err = e.RegistrationRepository.Save(registration)
	if err != nil {
		return nil, err
	}
	return &RegistrationOutputDTO{
		ID:          registration.ID,
		User:        registration.User,
		FullName:    registration.FullName,
		Email:       registration.Email,
		PhoneNumber: registration.PhoneNumber,
	}, nil
}
