package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/auth"
)

type AuthRepository interface {
	Signup(ctx context.Context, auth auth.AuthAggregate) error
}
