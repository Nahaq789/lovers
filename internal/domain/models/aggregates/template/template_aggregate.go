package template

import (
	"errors"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/template/detail"
	"lovers/internal/domain/models/template/detail/detailid"
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
	details       []detail.TemplateDetail
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
		details:       []detail.TemplateDetail{},
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

func (ta *TemplateAggregate) GetDetails() []detail.TemplateDetail {
	return ta.details
}

func (ta *TemplateAggregate) AddDetail(d detail.TemplateDetail) error {
	for _, exist := range ta.details {
		if exist.Equal(d.GetTemplateDetailId()) {
			return errors.New("同じ明細がすでに存在します。")
		}
	}

	ta.details = append(ta.details, d)
	return nil
}

func (ta *TemplateAggregate) RemoveDetail(d detailid.TemplateDetailId) error {
	for i, exist := range ta.details {
		if exist.Equal(d) {
			ta.details = append(ta.details[:i], ta.details[i+1:]...)
			return nil
		}
	}

	return errors.New("削除対象の明細が見つかりませんでした。")
}
