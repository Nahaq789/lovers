package expense

import (
	valueObjectUuid "lovers/internal/domain/models/valueobjects/uuid"

	"github.com/google/uuid"
)

type ExpenseId struct {
	value uuid.UUID
}

func NewExpenseId() (ExpenseId, error) {
	return newExpenseIdWithGenerator(valueObjectUuid.DefaultGenerator)
}

func newExpenseIdWithGenerator(generator valueObjectUuid.UUIDGenerator) (ExpenseId, error) {
	u, err := generator()
	if err != nil {
		return ExpenseId{}, err
	}

	return ExpenseId{value: u}, nil
}

func NewExpenseIdFromString(s string) (ExpenseId, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return ExpenseId{}, err
	}

	return ExpenseId{value: v}, nil
}

func (e ExpenseId) GetValue() string {
	return e.value.String()
}