package events

type Event interface {
	EventId() string
	OccurredAt() string
}
