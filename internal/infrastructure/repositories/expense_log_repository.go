package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/expense/log"
	"lovers/internal/shared/infrastructure/db"
)

type ExpenseLogRepositoryImpl struct {
	db *db.DbClient
}

func NewExpenseLogRepository(d *db.DbClient) *ExpenseLogRepositoryImpl {
	return &ExpenseLogRepositoryImpl{
		db: d,
	}
}

func (elr *ExpenseLogRepositoryImpl) Add(ctx context.Context, e []*log.ExpenseLog) error {
	return nil
}
