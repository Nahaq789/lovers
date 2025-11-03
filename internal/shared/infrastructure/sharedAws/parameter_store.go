package sharedAws

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type ParameterStoreClient struct {
	logger *slog.Logger
	client *ssm.Client
}

func InitParameterStoreClient(ctx context.Context, l *slog.Logger) (*ParameterStoreClient, error) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		l.ErrorContext(ctx, "failed init ssm client", "error", err)
		return nil, err
	}

	client := ssm.NewFromConfig(sdkConfig)
	return &ParameterStoreClient{
		logger: l,
		client: client,
	}, nil
}

func (p *ParameterStoreClient) GetParameter(ctx context.Context, name string) (string, error) {
	result, err := p.client.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		p.logger.ErrorContext(ctx, "failed to get parameter", "name", name, "error", err)
		return "", err
	}
	return *result.Parameter.Value, nil
}
