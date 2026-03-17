package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/expense"
)

type ExpenseRepository interface {
	Add(ctx context.Context, expense *expense.ExpenseAggregate) error
}
