package templateid

import (
	"github.com/google/uuid"
	valueObjectUuid "lovers/internal/domain/models/valueobjects/uuid"
)

type TemplateId struct {
	value uuid.UUID
}

func NewTemplateId() (TemplateId, error) {
	return newTemplateIdWithGenerator(valueObjectUuid.DefaultGenerator)
}

func newTemplateIdWithGenerator(generator valueObjectUuid.UUIDGenerator) (TemplateId, error) {
	u, err := generator()
	if err != nil {
		return TemplateId{}, err
	}

	return TemplateId{value: u}, nil
}

func NewTemplateIdFromString(s string) (TemplateId, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return TemplateId{}, err
	}

	return TemplateId{value: v}, nil
}

func (t TemplateId) GetValue() string {
	return t.value.String()
}
