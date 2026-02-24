package event

import (
	valueObjectUuid "lovers/internal/domain/models/valueobjects/uuid"

	"github.com/google/uuid"
)

type EventId struct {
	value uuid.UUID
}

func NewEventId() (EventId, error) {
	return newEventIdWithGenerator(valueObjectUuid.DefaultGenerator)
}

func newEventIdWithGenerator(generator valueObjectUuid.UUIDGenerator) (EventId, error) {
	u, err := generator()
	if err != nil {
		return EventId{}, err
	}

	return EventId{value: u}, nil
}

func NewEventIdFromString(s string) (EventId, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return EventId{}, err
	}

	return EventId{value: v}, nil
}

func (e EventId) GetValue() string {
	return e.value.String()
}

func (e EventId) Equal(n EventId) bool {
	return e.value == n.value
}
