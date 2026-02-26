package events

import "reflect"

type EventSubscriber interface {
	EventType() reflect.Type
	HandleEvent(event Event) error
}
