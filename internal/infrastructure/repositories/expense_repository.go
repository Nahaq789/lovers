package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/expense"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/logger"
)

type ExpenseRepositoryImpl struct {
	db *db.DbClient
}

func NewExpenseRepository(d *db.DbClient) *ExpenseRepositoryImpl {
	return &ExpenseRepositoryImpl{db: d}
}

func (e *ExpenseRepositoryImpl) Add(ctx context.Context, expense expense.ExpenseAggregate) error {
	l := logger.FromContext(ctx)
	c := e.db.GetClient()

	tx, txErr := c.BeginTx(ctx, nil)
	if txErr != nil {
		l.ErrorContext(ctx, "failed begin transaction", "error", txErr)
		return txErr
	}
	defer tx.Rollback()

	expenseQuery := `insert into expense (expense_id, group_id, payment_by, category_id, amount, nominal, payment_date, description, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	paymentUser := expense.GetPaymentUser()
	_, err := tx.ExecContext(ctx, expenseQuery,
		expense.GetExpenseId(),
		expense.GetGroupId(),
		paymentUser.GetUserId(),
		expense.GetCategoryId(),
		paymentUser.GetAmount(),
		expense.GetNominal(),
		expense.GetPaymentDate(),
		expense.GetDescription(),
		expense.GetCreatedAt(),
		expense.GetUpdatedAt())

	if err != nil {
		l.ErrorContext(ctx, "failed to insert expense", "error", err)
		return err
	}

	return tx.Commit()
}
