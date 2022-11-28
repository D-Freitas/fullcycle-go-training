package notification

import (
	"fmt"
	"strings"
)

type SignUpNotification struct {
	context string
	errors  []string
}

func NewSignUpNotification() *SignUpNotification {
	notification := &SignUpNotification{}
	notification.context = "signup"
	return notification
}

func (r *SignUpNotification) AddError(errorMessage string) {
	r.errors = append(r.errors, errorMessage)
}

func (r *SignUpNotification) HasErrors() bool {
	return len(r.errors) > 0
}

func (r *SignUpNotification) Messages() string {
	output := fmt.Sprintf("%s: %s", r.context, strings.Join(r.errors, ", "))
	return output
}
