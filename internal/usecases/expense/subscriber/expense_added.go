package eventhandler

import (
	"fmt"
	"lovers/internal/domain/events"
	"lovers/internal/domain/events/expense"
	"lovers/internal/domain/repositories"
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
	return reflect.TypeOf(expense.ExpenseAdded{})
}

func (ea *ExpenseAddedSubscriber) HandleEvent(event events.Event) error {
	// 明細追加ドメインイベントにキャストしないといけない
	_, ok := event.(*expense.ExpenseAdded)
	if !ok {
		return fmt.Errorf("unexpected event type: want *expense.ExpenseAdded, got %T", event)
	}
	return nil
}
