package events

import "lovers/internal/domain/models/event"

type Event interface {
	EventId() event.EventId
	OccurredAt() event.OccurredAt
}
