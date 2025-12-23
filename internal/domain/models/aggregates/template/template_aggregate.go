package template

import (
	"errors"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/template/expense"
	"lovers/internal/domain/models/template/expense/expenseid"
	"lovers/internal/domain/models/template/templateid"
	"lovers/internal/domain/models/template/templatename"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/updatedat"
)

type TemplateAggregate struct {
	template_id   templateid.TemplateId
	group_id      groupid.GroupId
	created_by    userid.UserId
	template_name templatename.TemplateName
	created_at    createdat.CreatedAt
	updated_at    updatedat.UpdatedAt
	details       []expense.TemplateExpense
}

func NewTemplateAggregate(
	templateId templateid.TemplateId,
	groupId groupid.GroupId,
	createdBy userid.UserId,
	templateName templatename.TemplateName,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt,
) *TemplateAggregate {
	return &TemplateAggregate{
		template_id:   templateId,
		group_id:      groupId,
		created_by:    createdBy,
		template_name: templateName,
		created_at:    createdAt,
		updated_at:    updatedAt,
		details:       []expense.TemplateExpense{},
	}
}

func (ta *TemplateAggregate) GetTemplateId() templateid.TemplateId {
	return ta.template_id
}

func (ta *TemplateAggregate) GetGroupId() groupid.GroupId {
	return ta.group_id
}

func (ta *TemplateAggregate) GetCreatedBy() userid.UserId {
	return ta.created_by
}

func (ta *TemplateAggregate) GetTemplateName() templatename.TemplateName {
	return ta.template_name
}

func (ta *TemplateAggregate) GetCreatedAt() createdat.CreatedAt {
	return ta.created_at
}

func (ta *TemplateAggregate) GetUpdatedAt() updatedat.UpdatedAt {
	return ta.updated_at
}

func (ta *TemplateAggregate) GetDetails() []expense.TemplateExpense {
	return ta.details
}

func (ta *TemplateAggregate) AddDetail(d expense.TemplateExpense) error {
	for _, exist := range ta.details {
		if exist.Equal(d.GetTemplateExpenseId()) {
			return errors.New("同じ支出がすでに存在します。")
		}
	}

	ta.details = append(ta.details, d)
	return nil
}

func (ta *TemplateAggregate) RemoveDetail(d expenseid.TemplateExpenseId) error {
	for i, exist := range ta.details {
		if exist.Equal(d) {
			ta.details = append(ta.details[:i], ta.details[i+1:]...)
			return nil
		}
	}

	return errors.New("削除対象の支出が見つかりませんでした。")
}
