package category

import (
	"github.com/google/uuid"
	valueObjectUuid "lovers/internal/domain/models/valueobjects/uuid"
)

type CategoryId struct {
	value uuid.UUID
}

func NewCategoryId() (CategoryId, error) {
	return newCategoryIdWithGenerator(valueObjectUuid.DefaultGenerator)
}

func newCategoryIdWithGenerator(generator valueObjectUuid.UUIDGenerator) (CategoryId, error) {
	u, err := generator()
	if err != nil {
		return CategoryId{}, err
	}

	return CategoryId{value: u}, nil
}

func NewCategoryIdFromString(s string) (CategoryId, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return CategoryId{}, err
	}

	return CategoryId{value: v}, nil
}

func (c CategoryId) GetValue() string {
	return c.value.String()
}

