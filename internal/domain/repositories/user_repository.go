package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates"
	"lovers/internal/domain/models/entity"
	userid "lovers/internal/domain/models/user/user_id"
)

type UserRepository interface {
	Register(ctx context.Context, user aggregates.UserAggregate) error
	GetUser(ctx context.Context, userId userid.UserId) (*entity.UserEntity, error)
}
