package aws

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoClient struct {
	logger *slog.Logger
	client *cognitoidentityprovider.Client
}

func InitCognitoClient(ctx context.Context, l *slog.Logger) (*CognitoClient, error) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		l.ErrorContext(ctx, "failed init cognito client", "error", err)
		return nil, err
	}

	client := cognitoidentityprovider.NewFromConfig(sdkConfig)

	return &CognitoClient{
		logger: l,
		client: client,
	}, nil
}
