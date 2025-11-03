package initialize

import (
	"context"
	"log/slog"
	authDi "lovers/cmd/di/auth"
	"lovers/internal/shared/config"
	"lovers/internal/shared/infrastructure/sharedAws"
)

func InitAuth(ctx context.Context, l *slog.Logger) (*authDi.AuthSet, error) {
	cognitoClient, err := sharedAws.InitCognitoClient(ctx, l)
	if err != nil {
		return nil, err
	}

	cognitoConfig := config.LoadCognitoConfig()

	authSet := authDi.Initialize(l, cognitoClient, cognitoConfig)
	return authSet, nil
}
