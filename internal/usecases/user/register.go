package user

import (
	"context"
	"errors"
	"lovers/internal/domain/models/aggregates/user"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/user/username"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/email"
	"lovers/internal/domain/models/valueobjects/updatedat"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	userDto "lovers/internal/usecases/dto/user"
)

type UserRegistration struct {
	userRepository repositories.UserRepository
}

func NewUserRegistration(u repositories.UserRepository) *UserRegistration {
	return &UserRegistration{
		userRepository: u,
	}
}

func (ur *UserRegistration) Execute(ctx context.Context, d *userDto.UserRegistrationDto) error {
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

	exist, err := ur.userRepository.Exists(ctx, &userId, &email)
	if exist || err != nil {
		if err == nil {
			return errors.New("ユーザーはすでに登録されています。")
		}
		l.ErrorContext(ctx, "ユーザー登録処理でエラーが発生しました。", "error", err)
		return err
	}

	userName, err := username.NewUserName(d.UserName)
	if err != nil {
		l.ErrorContext(ctx, "ユーザー登録処理でエラーが発生しました。", "error", err)
		return err
	}

	createdAt := createdat.NewCreatedAt()
	updatedAt := updatedat.NewUpdatedAt()
	agg := user.NewUserAggregate(userId, email, userName, createdAt, updatedAt)

	registerErr := ur.userRepository.Register(ctx, *agg)
	if registerErr != nil {
		l.ErrorContext(ctx, "ユーザー登録処理でエラーが発生しました。", "error", err)
		return registerErr
	}

	l.InfoContext(ctx, "ユーザー登録処理を終了します。")
	return nil
}
