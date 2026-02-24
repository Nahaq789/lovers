package initialize

import (
	"context"
	"lovers/cmd/di/expense"
	"lovers/internal/shared/infrastructure/db"
)

func InitExpense(ctx context.Context, d *db.DbClient) *expense.ExpenseSet {
	expenseSet := expense.Initialize(d)
	return expenseSet
}
