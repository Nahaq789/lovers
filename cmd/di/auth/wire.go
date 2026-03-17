//go:build wireinject
// +build wireinject

package auth

import (
	domainRepos "lovers/internal/domain/repositories"
	"lovers/internal/infrastructure/repositories"
	authController "lovers/internal/presentation/auth"
	"lovers/internal/shared/config"
	"lovers/internal/shared/infrastructure/sharedaws"
	use_case "lovers/internal/usecases/auth"

	"github.com/google/wire"
)

func ProvideAuthRepository(client *sharedaws.CognitoClient, cognitoConfig *config.CognitoConfig) *repositories.AuthRepositoryImpl {
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

func Initialize(client *sharedaws.CognitoClient, cfg *config.CognitoConfig) *AuthSet {
	wire.Build(
		authRepositorySet,
		signUpSet,
		authControllerSet,
		wire.Struct(new(AuthSet), "*"),
	)
	return nil
}
