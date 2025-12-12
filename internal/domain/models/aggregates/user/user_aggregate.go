package user

import (
	userid "lovers/internal/domain/models/user/user_id"
	username "lovers/internal/domain/models/user/user_name"
	createdat "lovers/internal/domain/models/value_objects/createdAt"
	"lovers/internal/domain/models/value_objects/email"
	"lovers/internal/domain/models/value_objects/password"
	updatedat "lovers/internal/domain/models/value_objects/updatedAt"
)

type UserAggregate struct {
	userId    userid.UserId
	email     email.Email
	password  password.Password
	userName  username.UserName
	createdAt createdat.CreatedAt
	updatedAt updatedat.UpdatedAt
}

func NewUserAggregate(
	id userid.UserId,
	email email.Email,
	password password.Password,
	userName username.UserName,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt) *UserAggregate {
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

func (u *UserAggregate) GetCreatedAt() createdat.CreatedAt {
	return u.createdAt
}

func (u *UserAggregate) GetUpdatedAt() updatedat.UpdatedAt {
	return u.updatedAt
}
