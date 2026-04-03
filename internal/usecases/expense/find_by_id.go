package expense

import (
	"context"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
)

type ExpenseFindById struct {
	expenseRepository repositories.ExpenseRepository
}

func NewExpenseFindById(er repositories.ExpenseRepository) *ExpenseFindById {
	return &ExpenseFindById{
		expenseRepository: er,
	}
}

func (ef *ExpenseFindById) Execute(ctx context.Context, expenseId) error {
	l := logger.FromContext(ctx)
	l.InfoContext(ctx, "明細取得処理を開始します。")

	return nil
}
