package sharedAws

import (
	"context"
	"lovers/internal/shared/infrastructure/logger"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoClient struct {
	client *cognitoidentityprovider.Client
}

func InitCognitoClient(ctx context.Context) (*CognitoClient, error) {
	l := logger.FromContext(ctx)
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		l.ErrorContext(ctx, "failed init cognito client", "error", err)
		return nil, err
	}

	client := cognitoidentityprovider.NewFromConfig(sdkConfig)

	return &CognitoClient{
		client: client,
	}, nil
}

func (c *CognitoClient) GetClient() *cognitoidentityprovider.Client {
	return c.client
}
