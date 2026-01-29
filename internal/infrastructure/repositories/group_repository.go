package repositories

import (
	"context"
	"fmt"
	"lovers/internal/domain/models/aggregates/group"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/logger"
	"strings"
)

type GroupRepositoryImpl struct {
	db *db.DbClient
}

func NewGroupRepository(d *db.DbClient) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{db: d}
}

func (g *GroupRepositoryImpl) Create(ctx context.Context, group group.GroupAggregate) error {
	l := logger.FromContext(ctx)
	c := g.db.GetClient()

	tx, txErr := c.BeginTx(ctx, nil)
	if txErr != nil {
		l.ErrorContext(ctx, "failed begin transaction", "error", txErr)
		return txErr
	}
	defer tx.Rollback()

	groupQuery := `insert into "group" (group_id, created_by, group_name, created_at, updated_at) values ($1, $2, $3, $4, $5)`
	_, groupErr := c.ExecContext(ctx, groupQuery,
		group.GetGroupId().GetValue(),
		group.GetCreatedBy().GetValue(),
		group.GetGroupName().GetValue(),
		group.GetCreatedAt().GetValue(),
		group.GetUpdatedAt().GetValue(),
	)
	if groupErr != nil {
		l.ErrorContext(ctx, "failed to create group", "error", groupErr)
		return groupErr
	}

	members := group.GetMembers()
	if len(members) == 0 {
		l.WarnContext(ctx, "メンバーが0圏のため、group_memberへのインサートをスキップします。")
		return nil
	}

	memberQuery := `insert into "group_member" (group_member_id, group_id, user_id, created_at) values `
	values := []interface{}{}
	placeholders := []string{}

	for i, m := range members {
		offset := i * 4
		placeholders = append(placeholders,
			fmt.Sprintf("($%d, $%d, $%d, $%d)", offset+1, offset+2, offset+3, offset+4),
		)
		values = append(values,
			m.GetGroupMemberId().GetValue(),
			m.GetGroupId().GetValue(),
			m.GetUserId().GetValue(),
			group.GetCreatedAt(),
		)
	}

	memberQuery += strings.Join(placeholders, ", ")

	_, memberErr := c.ExecContext(ctx, memberQuery, values...)
	if memberErr != nil {
		l.ErrorContext(ctx, "failed to add group member", "error", memberErr)
		return memberErr
	}

	return tx.Commit()
}
