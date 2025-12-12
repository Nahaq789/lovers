package user

import (
	"context"
	userid "lovers/internal/domain/models/user/user_id"
	"lovers/internal/domain/models/value_objects/email"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	"lovers/internal/use_cases/dto/user"
)

type UserRegistration struct {
	userRepository repositories.UserRepository
}

func NewUserRegistration(u repositories.UserRepository) *UserRegistration {
	return &UserRegistration{
		userRepository: u,
	}
}

func (ur *UserRegistration) Execute(ctx context.Context, d *user.UserRegistrationDto) error {
	l := logger.FromContext(ctx)
	l.InfoContext(ctx, "ユーザー登録処理を開始します。")

	userId, err := userid.NewUserIdFromString(d.UserId)
	if err != nil {
		l.ErrorContext(ctx, "ユーザー登録処理でエラーが発生しました。", "error", err)
		return err
	}

	email, err := email.NewEmail(d.Email)
	if err != nil {
		l.ErrorContext(ctx, "ユーザー登録処理でエラーが発生しました。", "error", err)
		return err
	}

	exist, err := ur.userRepository.Exists(ctx, userId, email)
	if exist || err != nil {
		l.ErrorContext(ctx, "ユーザー登録処理でエラーが発生しました。", "error", err)
		return err
	}

	l.InfoContext(ctx, "ユーザー登録処理を終了します。")
	return nil
}
