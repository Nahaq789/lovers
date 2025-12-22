package detail

import (
	"github.com/google/uuid"
	valueObjectUuid "lovers/internal/domain/models/valueobjects/uuid"
)

type TemplateDetailId struct {
	value uuid.UUID
}

func NewTemplateDetailId() (TemplateDetailId, error) {
	return newTemplateDetailIdWithGenerator(valueObjectUuid.DefaultGenerator)
}

func newTemplateDetailIdWithGenerator(generator valueObjectUuid.UUIDGenerator) (TemplateDetailId, error) {
	u, err := generator()
	if err != nil {
		return TemplateDetailId{}, err
	}

	return TemplateDetailId{value: u}, nil
}

func NewTemplateDetailIdFromString(s string) (TemplateDetailId, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return TemplateDetailId{}, err
	}

	return TemplateDetailId{value: v}, nil
}

func (t TemplateDetailId) GetValue() string {
	return t.value.String()
}
