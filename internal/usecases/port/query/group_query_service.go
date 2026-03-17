package query

import (
	"context"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/group/member"
)

type GroupQueryService interface {
	FindMemberById(ctx context.Context, groupId groupid.GroupId) (*member.MemberUserIds, error)
}
