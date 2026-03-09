package events

import "reflect"

type EventPublisher struct {
	subscribers []EventSubscriber
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{
		subscribers: []EventSubscriber{},
	}
}

func (ep *EventPublisher) Subscribe(s EventSubscriber) {
	ep.subscribers = append(ep.subscribers, s)
}

func (ep *EventPublisher) Publish(event Event) error {
	for _, subscriber := range ep.subscribers {
		if subscriber.EventType() == reflect.TypeOf(event) {
			err := subscriber.HandleEvent(event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
