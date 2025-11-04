package auth

import (
	"context"
	"log/slog"
	"lovers/internal/domain/models/aggregates/authAggregate"
	"lovers/internal/domain/models/value_objects/email"
	"lovers/internal/domain/models/value_objects/password"
	"lovers/internal/domain/repositories"
	"lovers/internal/use_cases/dto/authDto"
)

type SignUp struct {
	logger         *slog.Logger
	authRepository repositories.AuthRepository
}

func NewSignUp(l *slog.Logger, a repositories.AuthRepository) *SignUp {
	return &SignUp{
		logger:         l,
		authRepository: a,
	}
}

func (s *SignUp) Execute(ctx context.Context, c *authDto.SignUpDto) error {
	s.logger.InfoContext(ctx, "SignUp処理を開始します。")
	email, err := email.NewEmail(c.Email)
	if err != nil {
		s.logger.ErrorContext(ctx, "SignUp処理でエラーが発生しました。", "error", err)
		return err
	}

	password, err := password.NewPassword(c.Password)
	if err != nil {
		s.logger.ErrorContext(ctx, "SignUp処理でエラーが発生しました。", "error", err)
		return err
	}

	a := authAggregate.NewAuthAggregate(*email, *password)
	result, err := s.authRepository.SignUp(ctx, a)
	s.logger.InfoContext(ctx, "result", "value", result)
	if err != nil {
		s.logger.ErrorContext(ctx, "SignUp処理でエラーが発生しました。", "error", err)
		return err
	}

	s.logger.InfoContext(ctx, "SignUp処理を終了します。")
	return nil
}
