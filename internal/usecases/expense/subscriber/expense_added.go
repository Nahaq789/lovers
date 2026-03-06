package eventhandler

import (
	"lovers/internal/domain/events/expense"
	"reflect"
)

type ExpenseAddedSubscriber struct {
}

func NewExpenseAddedSubscriber() *ExpenseAddedSubscriber {
	return &ExpenseAddedSubscriber{}
}

func (ea *ExpenseAddedSubscriber) EventType() reflect.Type {
	return reflect.TypeOf(expense.ExpenseAdded{})
}

func (ea *ExpenseAddedSubscriber) HandleEvent(event expense.ExpenseAdded) error {
	ここに処理をついかする
	repositoryもさくせいしてね
	return nil
}
