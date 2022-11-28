package contract

import "server/internal/sign-up/domain"

type SignUpRepository interface {
	Save(signUp *domain.SignUp) error
}
