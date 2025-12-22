package auth

import (
	"context"
	"lovers/internal/domain/models/aggregates/authaggregate"
	"lovers/internal/domain/models/valueobjects/email"
	"lovers/internal/domain/models/valueobjects/password"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	"lovers/internal/usecases/dto/auth"
)

type SignUp struct {
	authRepository repositories.AuthRepository
}

func NewSignUp(a repositories.AuthRepository) *SignUp {
	return &SignUp{
		authRepository: a,
	}
}

func (s *SignUp) Execute(ctx context.Context, c *auth.SignUpDto) error {
	l := logger.FromContext(ctx)
	l.InfoContext(ctx, "SignUp処理を開始します。")
	email, err := email.NewEmail(c.Email)
	if err != nil {
		l.ErrorContext(ctx, "SignUp処理でエラーが発生しました。", "error", err)
		return err
	}

	password, err := password.NewPassword(c.Password)
	if err != nil {
		l.ErrorContext(ctx, "SignUp処理でエラーが発生しました。", "error", err)
		return err
	}

	a := authaggregate.NewAuthAggregate(email, password)
	result, err := s.authRepository.SignUp(ctx, a)
	l.InfoContext(ctx, "result", "value", result)
	if err != nil {
		l.ErrorContext(ctx, "SignUp処理でエラーが発生しました。", "error", err)
		return err
	}

	l.InfoContext(ctx, "SignUp処理を終了します。")
	return nil
}
