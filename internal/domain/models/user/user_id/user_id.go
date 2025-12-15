package userid

import "github.com/google/uuid"

type UserId struct {
	value uuid.UUID
}

type UUIDGenerator func() (uuid.UUID, error)

var defaultGenerator UUIDGenerator = uuid.NewV7

func NewUserId() (UserId, error) {
	return newUserIdWithGenerator(defaultGenerator)
}

func newUserIdWithGenerator(generator UUIDGenerator) (UserId, error) {
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
