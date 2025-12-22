package auth

type SignUpDto struct {
	Email    string
	Password string
}

func NewSignUpDto(email, password string) *SignUpDto {
	return &SignUpDto{
		Email:    email,
		Password: password,
	}
}
