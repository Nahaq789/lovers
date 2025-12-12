package createdat

import "time"

type CreatedAt struct {
	value time.Time
}

func NewCreatedAt() *CreatedAt {
	now := time.Now().UTC()
	return &CreatedAt{value: now}
}

func (c CreatedAt) GetValue() time.Time {
	return c.value.UTC()
}

func (c CreatedAt) ToString() string {
	return c.value.UTC().Format(time.RFC3339)
}
