package entity

type UserEntity struct {
	userId   string
	email    string
	password string
}

func NewUserEntity(userId string, email string, password string) *UserEntity {
	return &UserEntity{
		userId:   userId,
		email:    email,
		password: password,
	}
}
