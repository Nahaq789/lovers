package member

import (
	valueObjectUuid "lovers/internal/domain/models/value_objects/uuid"

	"github.com/google/uuid"
)

type GroupMemberId struct {
	value uuid.UUID
}

func NewGroupMemberId() (GroupMemberId, error) {
	return newGroupMemberIdWithGenerator(valueObjectUuid.DefaultGenerator)
}

func newGroupMemberIdWithGenerator(generator valueObjectUuid.UUIDGenerator) (GroupMemberId, error) {
	u, err := generator()
	if err != nil {
		return GroupMemberId{}, err
	}

	return GroupMemberId{value: u}, nil
}

func NewGroupMemberIdFromString(s string) (GroupMemberId, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return GroupMemberId{}, err
	}

	return GroupMemberId{value: v}, nil
}

func (gm GroupMemberId) GetValue() string {
	return gm.value.String()
}