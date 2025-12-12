package user

import (
	"context"
	userid "lovers/internal/domain/models/user/user_id"
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

	// IDチェック
	userId, err := userid.NewUserIdFromString(d.UserId)
	if err != nil {
		l.ErrorContext(ctx, "ユーザー登録処理でエラーが発生しました。", "error", err)
		return err
	}
	exist, err := ur.userRepository.ExistsUserId(ctx, userId)
	if exist || err != nil {
		l.ErrorContext(ctx, "ユーザー登録処理でエラーが発生しました。", "error", err)
		return err
	}

	l.InfoContext(ctx, "ユーザー登録処理を終了します。")
	return nil
}
