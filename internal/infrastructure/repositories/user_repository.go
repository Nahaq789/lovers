package repositories

import (
	"context"
	"log/slog"
	"lovers/internal/domain/models/aggregates"
	"lovers/internal/domain/models/entity"
	userid "lovers/internal/domain/models/user/user_id"
)

type UserRepositoryImpl struct {
	logger *slog.Logger
}

func NewUserRepository(l *slog.Logger) *UserRepositoryImpl {
	return &UserRepositoryImpl{logger: l}
}

func (u UserRepositoryImpl) Register(ctx context.Context, user aggregates.UserAggregate) error {
	return nil
}

func (u UserRepositoryImpl) GetUser(ctx context.Context, userId userid.UserId) (*entity.UserEntity, error) {
	return nil, nil
}
