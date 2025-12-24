package repositories

import (
	"context"
	"database/sql"
	"errors"
	"lovers/internal/domain/entity"
	"lovers/internal/domain/models/aggregates/user"
	userid "lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/email"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/logger"
)

type UserRepositoryImpl struct {
	db *db.DbClient
}

func NewUserRepository(d *db.DbClient) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: d}
}

func (u *UserRepositoryImpl) Register(ctx context.Context, user user.UserAggregate) error {
	l := logger.FromContext(ctx)
	query := `INSERT INTO "user" (user_id, user_name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	c := u.db.GetClient()
	_, err := c.ExecContext(ctx, query,
		user.GetUserId().GetValue(), user.GetUserName().GetValue(), user.GetEmail().GetValue(), user.GetCreatedAt().GetValue(), user.GetUpdatedAt().GetValue())

	if err != nil {
		l.ErrorContext(ctx, "failed to register user", "error", err)
		return err
	}

	return nil
}

func (u *UserRepositoryImpl) GetUser(ctx context.Context, userId userid.UserId) (*entity.UserEntity, error) {
	return nil, nil
}

func (u *UserRepositoryImpl) Exists(ctx context.Context, userId *userid.UserId, email *email.Email) (bool, error) {
	l := logger.FromContext(ctx)
	query := `select user_id from "user" where user_id = $1 or email = $2`
	c := u.db.GetClient()

	var id string
	err := c.QueryRowContext(ctx, query, userId.GetValue(), email.GetValue()).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// ユーザーが存在しない場合
			return false, nil
		}
		// DBエラー
		l.ErrorContext(ctx, "failed to check user existence", "error", err, "user_id", userId.GetValue())
		return false, err
	}

	// ユーザーが存在する場合
	l.WarnContext(ctx, "user already exists", "user_id", id)
	return true, nil
}
