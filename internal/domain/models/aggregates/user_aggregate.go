package aggregates

import (
	"lovers/internal/domain/models/user/email"
	"lovers/internal/domain/models/user/password"
)

type UserAggregate struct {
	userId   string
	email    email.Email
	password password.Password
}
