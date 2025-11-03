//go:build wireinject
// +build wireinject

package authDi

import (
	"log/slog"
	domainRepos "lovers/internal/domain/repositories"
	"lovers/internal/infrastructure/repositories"
	authController "lovers/internal/presentations/auth"
	"lovers/internal/shared/config"
	"lovers/internal/shared/infrastructure/aws"
	use_case "lovers/internal/use_cases/auth"

	"github.com/google/wire"
)

func ProvideAuthRepository(logger *slog.Logger, client *aws.CognitoClient, cognitoConfig *config.CognitoConfig) *repositories.AuthRepositoryImpl {
	repository := repositories.NewAuthRepositoryImpl(logger, client, cognitoConfig)
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

func Initialize(logger *slog.Logger, client *aws.CognitoClient, cfg *config.CognitoConfig) *AuthSet {
	wire.Build(
		authRepositorySet,
		signUpSet,
		authControllerSet,
		wire.Struct(new(AuthSet), "*"),
	)
	return nil
}
