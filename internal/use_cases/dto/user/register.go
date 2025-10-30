package user

type UserRegistrationDto struct {
	UserId string
}

func NewUserRegistrationDto(userId string) *UserRegistrationDto {
	return &UserRegistrationDto{
		UserId: userId,
	}
}
