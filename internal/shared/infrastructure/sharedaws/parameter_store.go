package sharedaws

import (
	"context"
	"lovers/internal/shared/infrastructure/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type ParameterStoreClient struct {
	client *ssm.Client
}

func InitParameterStoreClient(ctx context.Context) (*ParameterStoreClient, error) {
	l := logger.FromContext(ctx)
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		l.ErrorContext(ctx, "failed init ssm client", "error", err)
		return nil, err
	}

	client := ssm.NewFromConfig(sdkConfig)
	return &ParameterStoreClient{
		client: client,
	}, nil
}

func (p *ParameterStoreClient) GetParameter(ctx context.Context, name string) (string, error) {
	result, err := p.client.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		l := logger.FromContext(ctx)
		l.ErrorContext(ctx, "failed to get parameter", "name", name, "error", err)
		return "", err
	}
	return *result.Parameter.Value, nil
}
