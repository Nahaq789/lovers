package query

import (
	"context"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/user/userid"
)

type GroupQueryService interface {
	FindMemberById(ctx context.Context, groupId groupid.GroupId) ([]userid.UserId, error)
}
