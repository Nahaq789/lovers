package group

import (
	"errors"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/group/groupname"
	"lovers/internal/domain/models/group/member"
	"lovers/internal/domain/models/group/member/memberid"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/updatedat"
	"slices"
)

type GroupAggregate struct {
	groupId    groupid.GroupId
	createdBy  userid.UserId
	group_name groupname.GroupName
	createdAt  createdat.CreatedAt
	updatedAt  updatedat.UpdatedAt
	member     []*member.GroupMember
}

func NewGroupAggregate(
	groupId groupid.GroupId,
	createdBy userid.UserId,
	groupName groupname.GroupName,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt,
) *GroupAggregate {
	return &GroupAggregate{
		groupId:    groupId,
		createdBy:  createdBy,
		group_name: groupName,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
		member:     []*member.GroupMember{},
	}
}

func (ga *GroupAggregate) GetGroupId() groupid.GroupId {
	return ga.groupId
}

func (ga *GroupAggregate) GetCreatedBy() userid.UserId {
	return ga.createdBy
}

func (ga *GroupAggregate) GetGroupName() groupname.GroupName {
	return ga.group_name
}

func (ga *GroupAggregate) GetCreatedAt() createdat.CreatedAt {
	return ga.createdAt
}

func (ga *GroupAggregate) GetUpdatedAt() updatedat.UpdatedAt {
	return ga.updatedAt
}

func (ga *GroupAggregate) GetMembers() []*member.GroupMember {
	return ga.member
}

func (ga *GroupAggregate) CreateMember(u userid.UserId) (*member.GroupMember, error) {
	memberId, err := memberid.NewGroupMemberId()
	if err != nil {
		return nil, err
	}

	createdAt := createdat.NewCreatedAt()
	member := member.NewGroupMember(memberId, ga.groupId, u, createdAt)
	return member, nil
}

func (ga *GroupAggregate) AddMember(u userid.UserId) error {
	for _, exist := range ga.member {
		if exist.GetUserId().Equal(u) {
			return errors.New("ユーザーはすでに追加されています。")
		}
	}

	m, err := ga.CreateMember(u)
	if err != nil {
		return err
	}

	ga.member = append(ga.member, m)
	return nil
}

func (ga *GroupAggregate) RemoveMember(m member.GroupMember) error {
	for i, exist := range ga.member {
		if exist.Equals(m) {
			ga.member = slices.Delete(ga.member, i, i+1)
			return nil
		}
	}
	return errors.New("削除対象のユーザーが見つかりませんでした。")
}
