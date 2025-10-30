package user

import (
	"log/slog"
	"lovers/internal/use_cases/dto/user"
)

type UserRegistration struct {
	logger *slog.Logger
}

func NewUserRegistration(l *slog.Logger) *UserRegistration {
	return &UserRegistration{
		logger: l,
	}
}

func (ur *UserRegistration) Execute(d *user.UserRegistrationDto) error {
	return nil
}
