//go:build wireinject
// +build wireinject

package aws

import (
	"context"
	"log/slog"
	"lovers/internal/shared/infrastructure/aws"

	"github.com/google/wire"
)

var cognitoSet = wire.NewSet(aws.InitCognitoClient)
var parameterStoreSet = wire.NewSet(aws.InitParameterStoreClient)

type AwsSet struct {
	Cognito        *aws.CognitoClient
	ParameterStore *aws.ParameterStoreClient
}

func Initialize(ctx context.Context, logger *slog.Logger) (*AwsSet, error) {
	wire.Build(cognitoSet, parameterStoreSet, wire.Struct(new(AwsSet), "*"))
	return nil, nil
}
