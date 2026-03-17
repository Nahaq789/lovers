package userid

import (
	"github.com/google/uuid"
	valueObjectUuid "lovers/internal/domain/models/valueobjects/uuid"
)

type UserId struct {
	value uuid.UUID
}

func NewUserId() (UserId, error) {
	return newUserIdWithGenerator(valueObjectUuid.DefaultGenerator)
}

func newUserIdWithGenerator(generator valueObjectUuid.UUIDGenerator) (UserId, error) {
	u, err := generator()
	if err != nil {
		return UserId{}, err
	}

	return UserId{value: u}, nil
}

func NewUserIdFromString(s string) (UserId, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return UserId{}, err
	}

	return UserId{value: v}, nil
}

func (u UserId) GetValue() string {
	return u.value.String()
}

func (u UserId) Equal(id UserId) bool {
	return u.GetValue() == id.GetValue()
}
