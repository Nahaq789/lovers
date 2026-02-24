//go:build wireinject
// +build wireinject

package expense

import (
	domainRepos "lovers/internal/domain/repositories"
	infraPort "lovers/internal/infrastructure/port"
	"lovers/internal/infrastructure/repositories"
	"lovers/internal/infrastructure/services"
	"lovers/internal/presentation/expense"
	"lovers/internal/shared/infrastructure/db"
	expenseCreate "lovers/internal/usecases/expense"
	"lovers/internal/usecases/port/query"

	"lovers/internal/usecases/port"

	"github.com/google/wire"
)

func ProvideExpenseRepository(d *db.DbClient) *repositories.ExpenseRepositoryImpl {
	repository := repositories.NewExpenseRepository(d)
	return repository
}

var expenseRepositorySet = wire.NewSet(
	ProvideExpenseRepository,
	wire.Bind(new(domainRepos.ExpenseRepository), new(*repositories.ExpenseRepositoryImpl)),
)

func ProvideGroupQueryService(d *db.DbClient) *services.GroupQueryServiceImpl {
	service := services.NewGroupQueryService(d)
	return service
}

var groupQueryServiceSet = wire.NewSet(
	ProvideGroupQueryService,
	wire.Bind(new(query.GroupQueryService), new(*services.GroupQueryServiceImpl)),
)

func ProvideTransactionManager(d *db.DbClient) *infraPort.TransactionManagerImpl {
	manager := infraPort.NewTransactionManager(d)
	return manager
}

var transactionManagerSet = wire.NewSet(
	ProvideTransactionManager,
	wire.Bind(new(port.TransactionManager), new(*infraPort.TransactionManagerImpl)),
)

var createSet = wire.NewSet(expenseCreate.NewExpenseCreate)
var expenseControllerSet = wire.NewSet(expense.NewExpenseController)

type ExpenseSet struct {
	ExpenseController *expense.ExpenseController
}

func Initialize(d *db.DbClient) *ExpenseSet {
	wire.Build(
		expenseRepositorySet,
		groupQueryServiceSet,
		transactionManagerSet,
		createSet,
		expenseControllerSet,
		wire.Struct(new(ExpenseSet), "*"),
	)
	return nil
}
