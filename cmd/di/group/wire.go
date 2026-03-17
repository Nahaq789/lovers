//go:build wireinject
// +build wireinject

package group

import (
	domainRepos "lovers/internal/domain/repositories"
	"lovers/internal/presentation/group"
	"lovers/internal/usecases/port"

	infraPort "lovers/internal/infrastructure/port"
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

func ProvideTransactionManager(d *db.DbClient) *infraPort.TransactionManagerImpl {
	manager := infraPort.NewTransactionManager(d)
	return manager
}

var transactionManagerSet = wire.NewSet(
	ProvideTransactionManager,
	wire.Bind(new(port.TransactionManager), new(*infraPort.TransactionManagerImpl)),
)

var createSet = wire.NewSet(groupCreate.NewGroupCreate)
var groupControllerSet = wire.NewSet(group.NewGroupController)

type GroupSet struct {
	GroupController *group.GroupController
}

func Initialize(d *db.DbClient) *GroupSet {
	wire.Build(
		groupRepositorySet,
		transactionManagerSet,
		createSet,
		groupControllerSet,
		wire.Struct(new(GroupSet), "*"),
	)
	return nil
}
