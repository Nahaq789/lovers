package request

type SignUpDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSignUpDto(email, password string) *SignUpDto {
	return &SignUpDto{
		Email:    email,
		Password: password,
	}
}
