package notification

import (
	"fmt"
	"strings"
)

type RegisterNotification struct {
	context string
	errors  []string
}

func NewRegisterNotification() *RegisterNotification {
	notification := &RegisterNotification{}
	notification.context = "register"
	return notification
}

func (r *RegisterNotification) AddError(errorMessage string) {
	r.errors = append(r.errors, errorMessage)
}

func (r *RegisterNotification) HasErrors() bool {
	return len(r.errors) > 0
}

func (r *RegisterNotification) Messages() string {
	output := fmt.Sprintf("%s: %s", r.context, strings.Join(r.errors, ", "))
	return output
}
