package repositories

import (
	"context"
	"fmt"
	"lovers/internal/domain/models/aggregates/expense"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/logger"
	"strings"
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

	paymentUsers := expense.GetPaymentUsers()
	s := make([]any, 0, len(paymentUsers.GetPaymentUsers())*10)
	placeHolders := make([]string, 0, len(paymentUsers.GetPaymentUsers()))

	for i, u := range paymentUsers.GetPaymentUsers() {
		s = append(s, expense.GetExpenseId())
		s = append(s, expense.GetGroupId())
		s = append(s, u.GetUserId())
		s = append(s, expense.GetCategoryId())
		s = append(s, u.GetAmount())
		s = append(s, expense.GetNominal())
		s = append(s, expense.GetPaymentDate())
		s = append(s, expense.GetDescription())
		s = append(s, expense.GetCreatedAt())
		s = append(s, expense.GetUpdatedAt())
		placeHolders = append(placeHolders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*10+1, i*10+2, i*10+3, i*10+4, i*10+5, i*10+6, i*10+7, i*10+8, i*10+9, i*10+10))
	}

	expenseQuery := `insert into expense (expense_id, group_id, payment_by, category_id, amount, nominal, payment_date, description, created_at, updated_at) values `
	expenseQuery += strings.Join(placeHolders, ",")
	_, err := tx.ExecContext(ctx, expenseQuery, s...)
	if err != nil {
		l.ErrorContext(ctx, "failed to insert expense", "error", err)
		return err
	}

	return tx.Commit()
}
