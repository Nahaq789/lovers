package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/auth"
)

type AuthRepository interface {
	SignUp(ctx context.Context, auth *authaggregate.AuthAggregate) (*string, error)
}
