package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/group"
)

type GroupRepository interface {
	Create(ctx context.Context, group group.GroupAggregate) error
}
