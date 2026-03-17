package event

import "time"

type OccurredAt struct {
	value time.Time
}

func NewOccurredAt() OccurredAt {
	now := time.Now().UTC()
	return OccurredAt{value: now}
}

func (o OccurredAt) GetValue() time.Time {
	return o.value.UTC()
}

func (o OccurredAt) ToString() string {
	return o.value.UTC().Format(time.RFC3339)
}
