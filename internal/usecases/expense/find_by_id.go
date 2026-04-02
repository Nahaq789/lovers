package expense

import (
	"context"
	"lovers/internal/domain/repositories"
)

type ExpenseFindById struct {
	expenseRepository repositories.ExpenseRepository
}

func NewExpenseFindById(er repositories.ExpenseRepository) *ExpenseFindById {
	return &ExpenseFindById{
		expenseRepository: er,
	}
}

func (ef *ExpenseFindById) Execute(ctx context.Context) error {
	return nil
}
