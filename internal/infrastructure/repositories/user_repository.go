package repositories

import (
	"context"
	"log/slog"
	"lovers/internal/domain/models/aggregates/user"
	"lovers/internal/domain/models/entity"
	userid "lovers/internal/domain/models/user/user_id"
	"lovers/internal/shared/infrastructure/db"
)

type UserRepositoryImpl struct {
	logger *slog.Logger
	db     *db.DbClient
}

func NewUserRepository(l *slog.Logger, d *db.DbClient) *UserRepositoryImpl {
	return &UserRepositoryImpl{logger: l, db: d}
}

func (u UserRepositoryImpl) Register(ctx context.Context, user user.UserAggregate) error {
	query := `INSERT INTO users (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	c := u.db.GetClient()
	_, err := c.ExecContext(ctx, query,
		user.GetUserId().GetValue(), user.GetUserName().GetValue(), user.GetEmail().GetValue(), user.GetCreatedAt().GetValue(), user.GetUpdatedAt().GetValue())

	if err != nil {
		u.logger.ErrorContext(ctx, "failed to register user", "error", err)
		return err
	}

	return nil
}

func (u UserRepositoryImpl) GetUser(ctx context.Context, userId userid.UserId) (*entity.UserEntity, error) {
	return nil, nil
}
