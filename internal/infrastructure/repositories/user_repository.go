package repositories

import (
	"context"
	"log/slog"
	"lovers/internal/domain/models/aggregates"
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

func (u UserRepositoryImpl) Register(ctx context.Context, user aggregates.UserAggregate) error {
	query := `INSERT INTO users (id, name, email) VALUES (?, ?, ?)`
	c := u.db.GetClient()
	_, err := c.ExecContext(ctx, query, "", "", "")

	if err != nil {
		u.logger.ErrorContext(ctx, "failed to register user", "error", err)
		return err
	}

	return nil
}

func (u UserRepositoryImpl) GetUser(ctx context.Context, userId userid.UserId) (*entity.UserEntity, error) {
	return nil, nil
}
