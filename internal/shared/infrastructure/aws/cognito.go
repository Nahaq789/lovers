package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoClient struct {
	client cognitoidentityprovider.Client
}

func InitCognitoClient(ctx context.Context) (*CognitoClient, error) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, nil
	}

	return nil, nil
}
