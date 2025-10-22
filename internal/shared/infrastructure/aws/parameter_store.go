package aws

import (
	"context"
	"log/slog"

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
