package expenselogid

import (
	"github.com/google/uuid"
	valueObjectUuid "lovers/internal/domain/models/valueobjects/uuid"
)

type ExpenseLogId struct {
	value uuid.UUID
}

func NewExpenseLogId() (ExpenseLogId, error) {
	return newExpenseLogIdWithGenerator(valueObjectUuid.DefaultGenerator)
}

func newExpenseLogIdWithGenerator(generator valueObjectUuid.UUIDGenerator) (ExpenseLogId, error) {
	u, err := generator()
	if err != nil {
		return ExpenseLogId{}, err
	}
	return ExpenseLogId{value: u}, nil
}

func NewExpenseLogIdFromString(s string) (ExpenseLogId, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return ExpenseLogId{}, err
	}
	return ExpenseLogId{value: v}, nil
}

func (e ExpenseLogId) GetValue() string {
	return e.value.String()
}

func (e ExpenseLogId) Equal(n ExpenseLogId) bool {
	return e.value == n.value
}
