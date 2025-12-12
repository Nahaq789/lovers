//go:build wireinject
// +build wireinject

package aws

import (
	"context"
	"lovers/internal/shared/infrastructure/sharedAws"

	"github.com/google/wire"
)

var cognitoSet = wire.NewSet(sharedAws.InitCognitoClient)
var parameterStoreSet = wire.NewSet(sharedAws.InitParameterStoreClient)

type AwsSet struct {
	Cognito        *sharedAws.CognitoClient
	ParameterStore *sharedAws.ParameterStoreClient
}

func Initialize(ctx context.Context) (*AwsSet, error) {
	wire.Build(cognitoSet, parameterStoreSet, wire.Struct(new(AwsSet), "*"))
	return nil, nil
}
