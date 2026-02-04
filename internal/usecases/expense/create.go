package expense

import (
	"context"
	"lovers/internal/domain/repositories"
	expenseDto "lovers/internal/usecases/dto/expense"
	"lovers/internal/usecases/port"
	"lovers/internal/usecases/port/query"
)

type ExpenseCreate struct {
	expenseRepository repositories.ExpenseRepository
	groupQueryService query.GroupQueryService
	txManager         port.TransactionManager
}

func NewExpenseRepository(er repositories.ExpenseRepository, gq query.GroupQueryService, tm port.TransactionManager) *ExpenseCreate {
	return &ExpenseCreate{expenseRepository: er, groupQueryService: gq, txManager: tm}
}

func (ec *ExpenseCreate) Execute(ctx context.Context, d *expenseDto.ExpenseCreateDto) error {
	return nil
}
