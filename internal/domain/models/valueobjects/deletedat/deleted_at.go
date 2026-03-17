package deletedat

import "time"

type DeletedAt struct {
	value time.Time
}

func NewDeletedAt() DeletedAt {
	now := time.Now().UTC()
	return DeletedAt{value: now}
}

func (c DeletedAt) GetValue() time.Time {
	return c.value.UTC()
}

func (c DeletedAt) ToString() string {
	return c.value.UTC().Format(time.RFC3339)
}
