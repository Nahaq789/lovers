//go:build wireinject
// +build wireinject

package group

import (
	domainRepos "lovers/internal/domain/repositories"
	"lovers/internal/presentation/group"

	"lovers/internal/infrastructure/repositories"
	"lovers/internal/shared/infrastructure/db"
	groupCreate "lovers/internal/usecases/group"

	"github.com/google/wire"
)

func ProvideGroupRepository(d *db.DbClient) *repositories.GroupRepositoryImpl {
	repository := repositories.NewGroupRepository(d)
	return repository
}

var groupRepositorySet = wire.NewSet(
	ProvideGroupRepository,
	wire.Bind(new(domainRepos.GroupRepository), new(*repositories.GroupRepositoryImpl)),
)

var createSet = wire.NewSet(groupCreate.NewGroupCreate)
var groupControllerSet = wire.NewSet(group.NewGroupController)

type GroupSet struct {
	GroupController *group.GroupController
}

func Initialize(d *db.DbClient) *GroupSet {
	wire.Build(
		groupRepositorySet,
		createSet,
		groupControllerSet,
		wire.Struct(new(GroupSet), "*"),
	)
	return nil
}
