package initialize

import (
	"context"
	authDi "lovers/cmd/di/auth"
	"lovers/internal/shared/config"
	"lovers/internal/shared/infrastructure/sharedAws"
)

func InitAuth(ctx context.Context, cognitoClient *sharedAws.CognitoClient) (*authDi.AuthSet, error) {
	cognitoConfig := config.LoadCognitoConfig()

	authSet := authDi.Initialize(cognitoClient, cognitoConfig)
	return authSet, nil
}
