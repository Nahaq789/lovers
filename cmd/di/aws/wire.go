//go:build wireinject
// +build wireinject

package aws

import (
	"context"
	"lovers/internal/shared/infrastructure/sharedaws"

	"github.com/google/wire"
)

var cognitoSet = wire.NewSet(sharedaws.InitCognitoClient)
var parameterStoreSet = wire.NewSet(sharedaws.InitParameterStoreClient)

type AwsSet struct {
	Cognito        *sharedaws.CognitoClient
	ParameterStore *sharedaws.ParameterStoreClient
}

func Initialize(ctx context.Context) (*AwsSet, error) {
	wire.Build(cognitoSet, parameterStoreSet, wire.Struct(new(AwsSet), "*"))
	return nil, nil
}
