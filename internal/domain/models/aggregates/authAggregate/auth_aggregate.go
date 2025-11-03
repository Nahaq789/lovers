package authAggregate

import (
	"lovers/internal/domain/models/value_objects/email"
	"lovers/internal/domain/models/value_objects/password"
)

type AuthAggregate struct {
	email    email.Email
	password password.Password
}

func NewAuthAggregate(email email.Email, password password.Password) *AuthAggregate {
	return &AuthAggregate{
		email:    email,
		password: password,
	}
}

func (a *AuthAggregate) GetEmail() email.Email {
	return a.email
}

func (a *AuthAggregate) GetPassword() password.Password {
	return a.password
}
