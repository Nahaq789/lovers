package services

import (
	"context"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/group/member"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/logger"
)

type GroupQueryServiceImpl struct {
	db *db.DbClient
}

func NewGroupQueryService(d *db.DbClient) *GroupQueryServiceImpl {
	return &GroupQueryServiceImpl{db: d}
}

func (gq *GroupQueryServiceImpl) FindMemberById(ctx context.Context, groupId groupid.GroupId) (*member.MemberUserIds, error) {
	c := gq.db.GetClient()
	l := logger.FromContext(ctx)
	query := `
			SELECT
    			gm.user_id
			FROM
    			"group" AS g
			INNER JOIN
    			group_member AS gm
			ON g.group_id = gm.group_id
			WHERE
    			g.group_id = $1
		`

	rows, err := c.QueryContext(ctx, query, groupId.GetValue())
	if err != nil {
		l.ErrorContext(ctx, "failed to select group members", "error", err)
		return nil, err
	}

	defer rows.Close()

	var (
		userIds []userid.UserId
		id      string
	)

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		userId, err := userid.NewUserIdFromString(id)
		if err != nil {
			return nil, err
		}
		userIds = append(userIds, userId)
	}

	if err := rows.Err(); err != nil {
		l.ErrorContext(ctx, "error occurred during iteration", "error", err)
		return nil, err
	}

	memberIds := member.NewMemberUserIds(userIds)
	return memberIds, nil
}
