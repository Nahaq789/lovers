package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/authAggregate"
)

type AuthRepository interface {
	SignUp(ctx context.Context, auth *authAggregate.AuthAggregate) error
}
