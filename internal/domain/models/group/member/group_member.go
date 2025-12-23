package member

import (
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/group/member/memberid"
	"lovers/internal/domain/models/user/userid"
)

type GroupMember struct {
	group_member_id memberid.GroupMemberId
	group_id        groupid.GroupId
	user_id         userid.UserId
}

func NewGroupMember(
	groupMemberId memberid.GroupMemberId,
	groupId groupid.GroupId,
	userId userid.UserId,
) *GroupMember {
	return &GroupMember{
		group_member_id: groupMemberId,
		group_id:        groupId,
		user_id:         userId,
	}
}

func (gm *GroupMember) GetGroupMemberId() memberid.GroupMemberId {
	return gm.group_member_id
}

func (gm *GroupMember) GetGroupId() groupid.GroupId {
	return gm.group_id
}

func (gm *GroupMember) GetUserId() userid.UserId {
	return gm.user_id
}

func (gm *GroupMember) isSameGroup(g groupid.GroupId) bool {
	return gm.group_id.GetValue() == g.GetValue()
}

func (gm *GroupMember) Equals(m GroupMember) bool {
	if gm.isSameGroup(m.group_id) {
		return gm.user_id.GetValue() == m.user_id.GetValue()
	}
	return false
}
