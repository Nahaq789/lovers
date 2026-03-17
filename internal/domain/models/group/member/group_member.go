package member

import (
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/group/member/memberid"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/createdat"
)

type GroupMember struct {
	groupMemberId memberid.GroupMemberId
	groupId       groupid.GroupId
	userId        userid.UserId
	createdAt     createdat.CreatedAt
}

func NewGroupMember(
	groupMemberId memberid.GroupMemberId,
	groupId groupid.GroupId,
	userId userid.UserId,
	createdAt createdat.CreatedAt,
) *GroupMember {
	return &GroupMember{
		groupMemberId: groupMemberId,
		groupId:       groupId,
		userId:        userId,
		createdAt:     createdAt,
	}
}

func (gm *GroupMember) GetGroupMemberId() memberid.GroupMemberId {
	return gm.groupMemberId
}

func (gm *GroupMember) GetGroupId() groupid.GroupId {
	return gm.groupId
}

func (gm *GroupMember) GetUserId() userid.UserId {
	return gm.userId
}

func (gm *GroupMember) GetCreatedAt() createdat.CreatedAt {
	return gm.createdAt
}

func (gm *GroupMember) isSameGroup(g groupid.GroupId) bool {
	return gm.groupId.GetValue() == g.GetValue()
}

func (gm *GroupMember) Equals(m GroupMember) bool {
	if gm.isSameGroup(m.groupId) {
		return gm.userId.GetValue() == m.userId.GetValue()
	}
	return false
}
