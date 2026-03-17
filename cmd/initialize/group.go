package initialize

import (
	"context"
	"lovers/cmd/di/group"
	"lovers/internal/shared/infrastructure/db"
)

func InitGroup(ctx context.Context, d *db.DbClient) *group.GroupSet {
	groupSet := group.Initialize(d)
	return groupSet
}
