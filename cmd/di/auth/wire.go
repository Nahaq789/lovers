//go:build wireinject
// +build wireinject

package authDi

import (
	domainRepos "lovers/internal/domain/repositories"
	"lovers/internal/infrastructure/repositories"
	authController "lovers/internal/presentation/auth"
	"lovers/internal/shared/config"
	"lovers/internal/shared/infrastructure/sharedAws"
	use_case "lovers/internal/use_cases/auth"

	"github.com/google/wire"
)

func ProvideAuthRepository(client *sharedAws.CognitoClient, cognitoConfig *config.CognitoConfig) *repositories.AuthRepositoryImpl {
	repository := repositories.NewAuthRepositoryImpl(client, cognitoConfig)
	return repository
}

var authRepositorySet = wire.NewSet(
	ProvideAuthRepository,
	wire.Bind(new(domainRepos.AuthRepository), new(*repositories.AuthRepositoryImpl)),
)

var signUpSet = wire.NewSet(use_case.NewSignUp)
var authControllerSet = wire.NewSet(authController.NewAuthController)

type AuthSet struct {
	AuthController *authController.AuthController
}

func Initialize(client *sharedAws.CognitoClient, cfg *config.CognitoConfig) *AuthSet {
	wire.Build(
		authRepositorySet,
		signUpSet,
		authControllerSet,
		wire.Struct(new(AuthSet), "*"),
	)
	return nil
}
