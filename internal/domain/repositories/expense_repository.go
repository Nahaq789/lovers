package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/expense"
	"lovers/internal/domain/models/expense/expenseid"
	"lovers/internal/domain/models/user/userid"
)

type ExpenseRepository interface {
	Add(ctx context.Context, expense *expense.ExpenseAggregate) error
	FindById(ctx context.Context, expenseId expenseid.ExpenseId, userId userid.UserId) (*expense.ExpenseAggregate, error)
}
