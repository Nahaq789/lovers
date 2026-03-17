package expenseid

import (
	"github.com/google/uuid"
	valueObjectUuid "lovers/internal/domain/models/valueobjects/uuid"
)

type TemplateExpenseId struct {
	value uuid.UUID
}

func NewTemplateDetailId() (TemplateExpenseId, error) {
	return newTemplateDetailIdWithGenerator(valueObjectUuid.DefaultGenerator)
}

func newTemplateDetailIdWithGenerator(generator valueObjectUuid.UUIDGenerator) (TemplateExpenseId, error) {
	u, err := generator()
	if err != nil {
		return TemplateExpenseId{}, err
	}

	return TemplateExpenseId{value: u}, nil
}

func NewTemplateDetailIdFromString(s string) (TemplateExpenseId, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return TemplateExpenseId{}, err
	}

	return TemplateExpenseId{value: v}, nil
}

func (t TemplateExpenseId) GetValue() string {
	return t.value.String()
}
