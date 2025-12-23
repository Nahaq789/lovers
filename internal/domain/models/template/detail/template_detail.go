package detail

import (
	"lovers/internal/domain/models/category/categoryid"
	"lovers/internal/domain/models/template/detail/detailid"
	"lovers/internal/domain/models/template/templateid"
	"lovers/internal/domain/models/valueobjects/amount"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/description"
	"lovers/internal/domain/models/valueobjects/nominal"
	"lovers/internal/domain/models/valueobjects/paymentdate"
	"lovers/internal/domain/models/valueobjects/updatedat"
)

type TemplateDetail struct {
	template_detail_id detailid.TemplateDetailId
	template_id        templateid.TemplateId
	category_id        categoryid.CategoryId
	amount             amount.Amount
	nominal            nominal.Nominal
	payment_date       paymentdate.PaymentDate
	description        description.Description
	created_at         createdat.CreatedAt
	updated_at         updatedat.UpdatedAt
}

func NewTemplateDetail(
	templateDetailId detailid.TemplateDetailId,
	templateId templateid.TemplateId,
	categoryId categoryid.CategoryId,
	amt amount.Amount,
	nom nominal.Nominal,
	paymentDate paymentdate.PaymentDate,
	desc description.Description,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt,
) *TemplateDetail {
	return &TemplateDetail{
		template_detail_id: templateDetailId,
		template_id:        templateId,
		category_id:        categoryId,
		amount:             amt,
		nominal:            nom,
		payment_date:       paymentDate,
		description:        desc,
		created_at:         createdAt,
		updated_at:         updatedAt,
	}
}

func (td *TemplateDetail) GetTemplateDetailId() detailid.TemplateDetailId {
	return td.template_detail_id
}

func (td *TemplateDetail) GetTemplateId() templateid.TemplateId {
	return td.template_id
}

func (td *TemplateDetail) GetCategoryId() categoryid.CategoryId {
	return td.category_id
}

func (td *TemplateDetail) GetAmount() amount.Amount {
	return td.amount
}

func (td *TemplateDetail) GetNominal() nominal.Nominal {
	return td.nominal
}

func (td *TemplateDetail) GetPaymentDate() paymentdate.PaymentDate {
	return td.payment_date
}

func (td *TemplateDetail) GetDescription() description.Description {
	return td.description
}

func (td *TemplateDetail) GetCreatedAt() createdat.CreatedAt {
	return td.created_at
}

func (td *TemplateDetail) GetUpdatedAt() updatedat.UpdatedAt {
	return td.updated_at
}

func (td *TemplateDetail) isSameTemplate(t templateid.TemplateId) bool {
	return td.template_id.GetValue() == t.GetValue()
}

func (td *TemplateDetail) Equal(d detailid.TemplateDetailId) bool {
	return td.template_detail_id.GetValue() == d.GetValue()
}
