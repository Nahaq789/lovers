package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"lovers/internal/domain/models/aggregates/expense/log"
	"lovers/internal/infrastructure/json/expense"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/logger"
	"lovers/internal/shared/infrastructure/transaction"
	"strings"
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
	l := logger.FromContext(ctx)
	tx := transaction.FromContext(ctx)

	s := make([]any, 0, len(e)*8)
	placeHolders := make([]string, 0, len(e))

	for i, log := range e {
		afterDataJson, err := json.Marshal(expense.NewAfterDataJson(*log.GetAfterData()))
		fmt.Printf("%+v \n", afterDataJson)
		if err != nil {
			l.ErrorContext(ctx, "failed to parse after data", "error", err)
			return err
		}

		s = append(s, log.GetExpenseLogId().GetValue())
		s = append(s, log.GetExpenseId().GetValue())
		s = append(s, log.GetGroupId().GetValue())
		s = append(s, log.GetUserId().GetValue())
		s = append(s, log.GetOperation())
		s = append(s, nil)
		s = append(s, afterDataJson)
		s = append(s, log.GetCreatedAt().GetValue())
		placeHolders = append(placeHolders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*8+1, i*8+2, i*8+3, i*8+4, i*8+5, i*8+6, i*8+7, i*8+8))
	}

	query := `insert into expense_log (expense_log_id, expense_id, group_id, user_id, operation, before_data, after_data, created_at) values `
	query += strings.Join(placeHolders, ",")
	_, err := tx.ExecContext(ctx, query, s...)
	if err != nil {
		l.ErrorContext(ctx, "failed to insert expense_log", "error", err)
		return err
	}
	return nil
}
