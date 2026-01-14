//go:build wireinject
// +build wireinject

package user

import (
	domainRepos "lovers/internal/domain/repositories"
	"lovers/internal/infrastructure/repositories"
	"lovers/internal/presentation/user"
	"lovers/internal/shared/infrastructure/db"
	userRegistration "lovers/internal/usecases/user"

	"github.com/google/wire"
)

func ProvideUserRepository(d *db.DbClient) *repositories.UserRepositoryImpl {
	repository := repositories.NewUserRepository(d)
	return repository
}

var userRepositorySet = wire.NewSet(
	ProvideUserRepository,
	wire.Bind(new(domainRepos.UserRepository), new(*repositories.UserRepositoryImpl)),
)

var registrationSet = wire.NewSet(userRegistration.NewUserRegistration)
var userControllerSet = wire.NewSet(user.NewUserController)

type UserSet struct {
	UserController *user.UserController
}

func Initialize(d *db.DbClient) *UserSet {
	wire.Build(
		userRepositorySet,
		registrationSet,
		userControllerSet,
		wire.Struct(new(UserSet), "*"),
	)
	return nil
}
