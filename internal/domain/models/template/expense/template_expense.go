package expense

import (
	"lovers/internal/domain/models/category/categoryid"
	"lovers/internal/domain/models/template/expense/expenseid"
	"lovers/internal/domain/models/template/templateid"
	"lovers/internal/domain/models/valueobjects/amount"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/description"
	"lovers/internal/domain/models/valueobjects/nominal"
	"lovers/internal/domain/models/valueobjects/paymentdate"
	"lovers/internal/domain/models/valueobjects/updatedat"
)

type TemplateExpense struct {
	template_expense_id expenseid.TemplateExpenseId
	template_id         templateid.TemplateId
	category_id         categoryid.CategoryId
	amount              amount.Amount
	nominal             nominal.Nominal
	payment_date        paymentdate.PaymentDate
	description         description.Description
	created_at          createdat.CreatedAt
	updated_at          updatedat.UpdatedAt
}

func NewTemplateDetail(
	templateExpenseId expenseid.TemplateExpenseId,
	templateId templateid.TemplateId,
	categoryId categoryid.CategoryId,
	amt amount.Amount,
	nom nominal.Nominal,
	paymentDate paymentdate.PaymentDate,
	desc description.Description,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt,
) *TemplateExpense {
	return &TemplateExpense{
		template_expense_id: templateExpenseId,
		template_id:         templateId,
		category_id:         categoryId,
		amount:              amt,
		nominal:             nom,
		payment_date:        paymentDate,
		description:         desc,
		created_at:          createdAt,
		updated_at:          updatedAt,
	}
}

func (td *TemplateExpense) GetTemplateExpenseId() expenseid.TemplateExpenseId {
	return td.template_expense_id
}

func (td *TemplateExpense) GetTemplateId() templateid.TemplateId {
	return td.template_id
}

func (td *TemplateExpense) GetCategoryId() categoryid.CategoryId {
	return td.category_id
}

func (td *TemplateExpense) GetAmount() amount.Amount {
	return td.amount
}

func (td *TemplateExpense) GetNominal() nominal.Nominal {
	return td.nominal
}

func (td *TemplateExpense) GetPaymentDate() paymentdate.PaymentDate {
	return td.payment_date
}

func (td *TemplateExpense) GetDescription() description.Description {
	return td.description
}

func (td *TemplateExpense) GetCreatedAt() createdat.CreatedAt {
	return td.created_at
}

func (td *TemplateExpense) GetUpdatedAt() updatedat.UpdatedAt {
	return td.updated_at
}

func (td *TemplateExpense) isSameTemplate(t templateid.TemplateId) bool {
	return td.template_id.GetValue() == t.GetValue()
}

func (td *TemplateExpense) Equal(d expenseid.TemplateExpenseId) bool {
	return td.template_expense_id.GetValue() == d.GetValue()
}
