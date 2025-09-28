//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"log/slog"
	"lovers/internal/shared/infrastructure/aws"

	"github.com/google/wire"
)

var cognitoSet = wire.NewSet(aws.InitCognitoClient)

type AwsSet struct {
	Cognito *aws.CognitoClient
}

func Initialize(ctx context.Context, logger *slog.Logger) (*AwsSet, error) {
	wire.Build(cognitoSet, wire.Struct(new(AwsSet), "*"))
	return nil, nil
}
