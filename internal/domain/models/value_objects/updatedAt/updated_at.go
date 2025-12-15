package updatedat

import "time"

type UpdatedAt struct {
	value time.Time
}

func NewUpdatedAt() UpdatedAt {
	now := time.Now().UTC()
	return UpdatedAt{value: now}
}

func (u UpdatedAt) GetValue() time.Time {
	return u.value.UTC()
}

func (u UpdatedAt) ToString() string {
	return u.value.UTC().Format(time.RFC3339)
}
