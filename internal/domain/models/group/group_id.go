package group

import (
	valueObjectUuid "lovers/internal/domain/models/value_objects/uuid"

	"github.com/google/uuid"
)

type GroupId struct {
	value uuid.UUID
}

func NewGroupId() (GroupId, error) {
	return newGroupIdWithGenerator(valueObjectUuid.DefaultGenerator)
}

func newGroupIdWithGenerator(generator valueObjectUuid.UUIDGenerator) (GroupId, error) {
	u, err := generator()
	if err != nil {
		return GroupId{}, err
	}

	return GroupId{value: u}, nil
}

func NewGroupIdFromString(s string) (GroupId, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return GroupId{}, err
	}

	return GroupId{value: v}, nil
}

func (g GroupId) GetValue() string {
	return g.value.String()
}