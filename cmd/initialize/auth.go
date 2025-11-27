package initialize

import (
	"context"
	"log/slog"
	authDi "lovers/cmd/di/auth"
	"lovers/internal/shared/config"
	"lovers/internal/shared/infrastructure/sharedAws"
)

func InitAuth(ctx context.Context, l *slog.Logger, cognitoClient *sharedAws.CognitoClient) (*authDi.AuthSet, error) {
	cognitoConfig := config.LoadCognitoConfig()

	authSet := authDi.Initialize(l, cognitoClient, cognitoConfig)
	return authSet, nil
}
