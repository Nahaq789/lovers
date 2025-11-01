package user

import (
	"lovers/internal/domain/models/user/email"
	"lovers/internal/domain/models/user/password"
	userid "lovers/internal/domain/models/user/user_id"
	username "lovers/internal/domain/models/user/user_name"
)

type UserAggregate struct {
	userId   userid.UserId
	email    email.Email
	password password.Password
	userName username.UserName
}

func NewUserAggregate(id userid.UserId, email email.Email, password password.Password, userName username.UserName) *UserAggregate {
	return &UserAggregate{
		userId:   id,
		email:    email,
		password: password,
		userName: userName,
	}
}

func (u *UserAggregate) GetUserId() userid.UserId {
	return u.userId
}

func (u *UserAggregate) GetEmail() email.Email {
	return u.email
}

func (u *UserAggregate) GetUserName() username.UserName {
	return u.userName
}
