package user

import (
	"context"
	"errors"
	"lovers/internal/domain/models/aggregates/user"
	userid "lovers/internal/domain/models/user/user_id"
	username "lovers/internal/domain/models/user/user_name"
	createdat "lovers/internal/domain/models/value_objects/createdAt"
	"lovers/internal/domain/models/value_objects/email"
	updatedat "lovers/internal/domain/models/value_objects/updatedAt"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	userDto "lovers/internal/use_cases/dto/user"
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
