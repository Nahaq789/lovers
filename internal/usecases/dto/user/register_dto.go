package user

type UserRegistrationDto struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

func NewUserRegistrationDto(userId, email, userName string) *UserRegistrationDto {
	return &UserRegistrationDto{
		UserId:   userId,
		Email:    email,
		UserName: userName,
	}
}
