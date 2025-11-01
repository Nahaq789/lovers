package repositories

import (
	"context"
	"log/slog"
	"lovers/internal/domain/models/aggregates/auth"
	"lovers/internal/shared/infrastructure/aws"
)

type AuthRepositoryImpl struct {
	logger *slog.Logger
	client *aws.CognitoClient
}

func NewAuthRepositoryImpl(l *slog.Logger, c *aws.CognitoClient) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		logger: l,
		client: c,
	}
}

func (a *AuthRepositoryImpl) Signup(ctx context.Context, auth auth.AuthAggregate) error {
	return nil
}
