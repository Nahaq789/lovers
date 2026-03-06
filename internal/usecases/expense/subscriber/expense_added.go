package eventhandler

import "reflect"

type ExpenseAddedSubscriber struct {
}

func NewExpenseAddedSubscriber() *ExpenseAddedSubscriber {
	return &ExpenseAddedSubscriber{}
}

func (ea *ExpenseAddedSubscriber) EventType() reflect.Type {
	return nil
}

func (ea *ExpenseAddedSubscriber) HandleEvent() error {
	return nil
}
