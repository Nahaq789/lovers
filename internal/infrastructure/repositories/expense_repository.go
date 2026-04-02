package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"lovers/internal/domain/models/aggregates/expense"
	"lovers/internal/domain/models/expense/expenseid"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/logger"
	"lovers/internal/shared/infrastructure/transaction"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ExpenseRepositoryImpl struct {
	db *db.DbClient
}

func NewExpenseRepository(d *db.DbClient) *ExpenseRepositoryImpl {
	return &ExpenseRepositoryImpl{db: d}
}

func (e *ExpenseRepositoryImpl) Add(ctx context.Context, expense *expense.ExpenseAggregate) error {
	l := logger.FromContext(ctx)
	tx := transaction.FromContext(ctx)

	paymentUsers := expense.GetPaymentUsers()
	s := make([]any, 0, len(paymentUsers.GetPaymentUsers())*10)
	placeHolders := make([]string, 0, len(paymentUsers.GetPaymentUsers()))

	for i, u := range paymentUsers.GetPaymentUsers() {
		s = append(s, expense.GetExpenseId().GetValue())
		s = append(s, expense.GetGroupId().GetValue())
		s = append(s, u.GetUserId().GetValue())
		s = append(s, expense.GetCategoryId().GetValue())
		s = append(s, u.GetAmount().GetValue())
		s = append(s, expense.GetNominal().GetValue())
		s = append(s, expense.GetPaymentDate().GetValue())
		s = append(s, expense.GetDescription().GetValue())
		s = append(s, expense.GetCreatedAt().GetValue())
		s = append(s, expense.GetUpdatedAt().GetValue())
		placeHolders = append(placeHolders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*10+1, i*10+2, i*10+3, i*10+4, i*10+5, i*10+6, i*10+7, i*10+8, i*10+9, i*10+10))
	}

	expenseQuery := `insert into expense (expense_id, group_id, user_id, category_id, amount, nominal, payment_date, description, created_at, updated_at) values `
	expenseQuery += strings.Join(placeHolders, ",")
	_, err := tx.ExecContext(ctx, expenseQuery, s...)
	if err != nil {
		l.ErrorContext(ctx, "failed to insert expense", "error", err)
		return err
	}

	return nil
}

func (e *ExpenseRepositoryImpl) FindById(ctx context.Context, expenseId expenseid.ExpenseId, userId userid.UserId) (*expense.ExpenseAggregate, error) {
	l := logger.FromContext(ctx)
	query := `
		SELECT
    		expense_id,
    		user_id,
    		group_id,
    		category_id,
    		amount,
    		nominal,
    		payment_date,
    		description,
    		deleted_at,
    		created_at,
    		updated_at
		FROM 
			expense
		WHERE 
			expense_id = $1
			AND user_id = $2`

	var (
		rawExpenseId   uuid.UUID
		rawUserId      uuid.UUID
		rawGroupId     uuid.UUID
		rawCategoryId  uuid.UUID
		rawAmount      int64
		rawNominal     string
		rawPaymentDate time.Time
		rawDescription sql.NullString
		rawDeletedAt   sql.NullTime
		rawCreatedAt   time.Time
		rawUpdatedAt   time.Time
	)

	c := e.db.GetClient()
	err := c.QueryRowContext(ctx, query, expenseId.GetValue(), userId.GetValue()).Scan(
		&rawExpenseId,
		&rawUserId,
		&rawGroupId,
		&rawCategoryId,
		&rawAmount,
		&rawNominal,
		&rawPaymentDate,
		&rawDescription,
		&rawDeletedAt,
		&rawCreatedAt,
		&rawUpdatedAt,
	)
	if err == sql.ErrNoRows {
		l.ErrorContext(ctx, "expense not found", "error", err)
		return nil, nil
	}

	if err != nil {
		l.ErrorContext(ctx, "failed to find expense", "error", err)
		return nil, err
	}
	return nil, nil
}
