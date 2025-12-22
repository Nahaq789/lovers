package repositories

import (
	"context"
	"lovers/internal/domain/entity"
	"lovers/internal/domain/models/aggregates/user"
	userid "lovers/internal/domain/models/user/user_id"
	"lovers/internal/domain/models/value_objects/email"
)

type UserRepository interface {
	Register(ctx context.Context, user user.UserAggregate) error
	GetUser(ctx context.Context, userId userid.UserId) (*entity.UserEntity, error)
	Exists(ctx context.Context, userId *userid.UserId, email *email.Email) (bool, error)
}
