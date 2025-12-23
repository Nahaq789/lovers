package group

import (
	"errors"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/group/groupname"
	"lovers/internal/domain/models/group/member"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/updatedat"
)

type GroupAggregate struct {
	group_id   groupid.GroupId
	created_by userid.UserId
	group_name groupname.GroupName
	created_at createdat.CreatedAt
	updated_at updatedat.UpdatedAt
	member     []member.GroupMember
}

func NewGroupAggregate(
	groupId groupid.GroupId,
	createdBy userid.UserId,
	groupName groupname.GroupName,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt,
) *GroupAggregate {
	return &GroupAggregate{
		group_id:   groupId,
		created_by: createdBy,
		group_name: groupName,
		created_at: createdAt,
		updated_at: updatedAt,
		member:     []member.GroupMember{},
	}
}

func (ga *GroupAggregate) GetGroupId() groupid.GroupId {
	return ga.group_id
}

func (ga *GroupAggregate) GetCreatedBy() userid.UserId {
	return ga.created_by
}

func (ga *GroupAggregate) GetGroupName() groupname.GroupName {
	return ga.group_name
}

func (ga *GroupAggregate) GetCreatedAt() createdat.CreatedAt {
	return ga.created_at
}

func (ga *GroupAggregate) GetUpdatedAt() updatedat.UpdatedAt {
	return ga.updated_at
}

func (ga *GroupAggregate) GetMembers() []member.GroupMember {
	return ga.member
}

func (ga *GroupAggregate) AddMember(m member.GroupMember) error {
	for _, exist := range ga.member {
		if exist.Equals(m) {
			return errors.New("ユーザーはすでに追加されています。")
		}
	}

	ga.member = append(ga.member, m)
	return nil
}

func (ga *GroupAggregate) RemoveMember(m member.GroupMember) error {
	for i, exist := range ga.member {
		if exist.Equals(m) {
			ga.member = append(ga.member[:i], ga.member[i+1:]...)
			return nil
		}
	}
	return errors.New("削除対象のユーザーが見つかりませんでした。")
}
