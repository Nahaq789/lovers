package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/group"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/logger"
)

type GroupRepositoryImpl struct {
	db *db.DbClient
}

func NewGroupRepository(d *db.DbClient) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{db: d}
}

func (g *GroupRepositoryImpl) Create(ctx context.Context, group group.GroupAggregate) error {
	l := logger.FromContext(ctx)
	query := `insert into "group" (group_id, created_by, group_name, created_at, updated_at) values ($1, $2, $3, $4, $5)`
	c := g.db.GetClient()
	_, err := c.ExecContext(ctx, query,
		group.GetGroupId().GetValue(),
		group.GetCreatedBy().GetValue(),
		group.GetGroupName().GetValue(),
		group.GetCreatedAt().GetValue(),
		group.GetUpdatedAt().GetValue(),
	)
	if err != nil {
		l.ErrorContext(ctx, "failed to create group", "error", err)
		return err
	}

	return nil
}
