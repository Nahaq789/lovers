package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/expense/log"
)

type ExpenseLogRepository interface {
	Add(ctx context.Context, e []*log.ExpenseLog) error
}
