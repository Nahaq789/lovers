package events

import (
	"context"
	"reflect"
)

type EventSubscriber interface {
	EventType() reflect.Type
	HandleEvent(ctx context.Context, event Event) error
}
