package initialize

import (
	"context"
	"lovers/cmd/di/user"
	"lovers/internal/shared/infrastructure/db"
)

func InitUser(ctx context.Context, d *db.DbClient) *user.UserSet {
	userSet := user.Initialize(d)
	return userSet
}
