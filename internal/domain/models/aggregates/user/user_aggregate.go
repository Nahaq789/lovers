package user

import (
	userid "lovers/internal/domain/models/user/userid"
	username "lovers/internal/domain/models/user/username"
	createdat "lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/email"
	updatedat "lovers/internal/domain/models/valueobjects/updatedat"
)

type UserAggregate struct {
	userId    userid.UserId
	email     email.Email
	userName  username.UserName
	createdAt createdat.CreatedAt
	updatedAt updatedat.UpdatedAt
}

func NewUserAggregate(
	id userid.UserId,
	email email.Email,
	userName username.UserName,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt) *UserAggregate {
	return &UserAggregate{
		userId:    id,
		email:     email,
		userName:  userName,
		createdAt: createdAt,
		updatedAt: updatedAt,
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
