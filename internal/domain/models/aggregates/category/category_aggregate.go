package category

import (
	"lovers/internal/domain/models/category/categoryid"
	"lovers/internal/domain/models/category/categoryname"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/updatedat"
)

type CategoryAggregate struct {
	category_id   categoryid.CategoryId
	group_id      groupid.GroupId
	created_by    userid.UserId
	category_name categoryname.CategoryName
	created_at    createdat.CreatedAt
	updated_at    updatedat.UpdatedAt
}

func NewCategoryAggregate(
	categoryId categoryid.CategoryId,
	groupId groupid.GroupId,
	createdBy userid.UserId,
	categoryName categoryname.CategoryName,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt,
) *CategoryAggregate {
	return &CategoryAggregate{
		category_id:   categoryId,
		group_id:      groupId,
		created_by:    createdBy,
		category_name: categoryName,
		created_at:    createdAt,
		updated_at:    updatedAt,
	}
}

func (ca *CategoryAggregate) GetCategoryId() categoryid.CategoryId {
	return ca.category_id
}

func (ca *CategoryAggregate) GetGroupId() groupid.GroupId {
	return ca.group_id
}

func (ca *CategoryAggregate) GetCreatedBy() userid.UserId {
	return ca.created_by
}

func (ca *CategoryAggregate) GetCategoryName() categoryname.CategoryName {
	return ca.category_name
}

func (ca *CategoryAggregate) GetCreatedAt() createdat.CreatedAt {
	return ca.created_at
}

func (ca *CategoryAggregate) GetUpdatedAt() updatedat.UpdatedAt {
	return ca.updated_at
}