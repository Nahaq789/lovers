package eventhandler

import (
	"context"
	"fmt"
	"lovers/internal/domain/events"
	"lovers/internal/domain/events/expense"
	"lovers/internal/domain/models/aggregates/expense/log"
	"lovers/internal/domain/models/expense/expenselogid"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	"reflect"
)

type ExpenseAddedSubscriber struct {
	repository repositories.ExpenseLogRepository
}

func NewExpenseAddedSubscriber(elr repositories.ExpenseLogRepository) ExpenseAddedSubscriber {
	return ExpenseAddedSubscriber{
		repository: elr,
	}
}

func (ea *ExpenseAddedSubscriber) EventType() reflect.Type {
	return reflect.TypeOf(&expense.ExpenseAdded{})
}

func (ea *ExpenseAddedSubscriber) HandleEvent(ctx context.Context, event events.Event) error {
	l := logger.FromContext(ctx)
	l.InfoContext(ctx, "明細ログ追加処理を開始します。")
	e, ok := event.(*expense.ExpenseAdded)
	if !ok {
		return fmt.Errorf("unexpected event type: want *expense.ExpenseAdded, got %T", event)
	}

	createdAt := createdat.NewCreatedAt()

	addedList := make([]*log.ExpenseLog, 0)
	for _, afterData := range e.AfterDataList() {
		logId, err := expenselogid.NewExpenseLogId()
		if err != nil {
			return err
		}
		added := log.NewExpenseLog(logId, e.ExpenseId(), e.GroupId(), e.UserId(), e.Operation(), nil, &afterData, createdAt)
		addedList = append(addedList, added)
	}
	dbErr := ea.repository.Add(ctx, addedList)
	if dbErr != nil {
		return dbErr
	}

	l.InfoContext(ctx, "明細ログ追加処理を終了します。")
	return nil
}
