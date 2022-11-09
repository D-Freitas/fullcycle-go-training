package entity

type RegistrationRepositoryInterface interface {
	Save(registration *Registration) error
}
