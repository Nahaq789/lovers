package aggregates

import (
	"lovers/internal/domain/models/user/email"
	"lovers/internal/domain/models/user/password"
	userid "lovers/internal/domain/models/user/user_id"
)

type UserAggregate struct {
	userId   userid.UserId
	email    email.Email
	password password.Password
}

func NewUserAggregate(id userid.UserId, email email.Email, password password.Password) *UserAggregate {
	return &UserAggregate{
		userId:   id,
		email:    email,
		password: password,
	}
}
