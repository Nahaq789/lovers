package user

type UserRegistrationDto struct {
	UserId   string
	Email    string
	UserName string
}

func NewUserRegistrationDto(userId, email, userName string) *UserRegistrationDto {
	return &UserRegistrationDto{
		UserId:   userId,
		Email:    email,
		UserName: userName,
	}
}
