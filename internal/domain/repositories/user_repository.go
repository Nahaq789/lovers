package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/user"
	"lovers/internal/domain/models/entity"
	userid "lovers/internal/domain/models/user/user_id"
	"lovers/internal/domain/models/value_objects/email"
)

type UserRepository interface {
	Register(ctx context.Context, user user.UserAggregate) error
	GetUser(ctx context.Context, userId userid.UserId) (*entity.UserEntity, error)
	ExistsUserId(ctx context.Context, userId userid.UserId) (bool, error)
	ExistsEmail(ctx context.Context, email email.Email) (bool, error)
}
