package expense

import (
	"errors"
	"lovers/internal/domain/models/category/categoryid"
	"lovers/internal/domain/models/expense/expenseid"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/amount"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/deletedat"
	"lovers/internal/domain/models/valueobjects/description"
	"lovers/internal/domain/models/valueobjects/nominal"
	"lovers/internal/domain/models/valueobjects/paymentdate"
	"lovers/internal/domain/models/valueobjects/updatedat"
)

type ExpenseAggregate struct {
	expense_id   expenseid.ExpenseId
	group_id     groupid.GroupId
	payment_by   userid.UserId
	category_id  categoryid.CategoryId
	amount       amount.Amount
	nominal      nominal.Nominal
	payment_date paymentdate.PaymentDate
	description  description.Description
	deleted_at   *deletedat.DeletedAt
	created_at   createdat.CreatedAt
	updated_at   updatedat.UpdatedAt
}

func NewExpenseAggregate(
	expenseId expenseid.ExpenseId,
	groupId groupid.GroupId,
	paymentBy userid.UserId,
	categoryId categoryid.CategoryId,
	amt amount.Amount,
	nom nominal.Nominal,
	paymentDate paymentdate.PaymentDate,
	desc description.Description,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt,
) *ExpenseAggregate {
	return &ExpenseAggregate{
		expense_id:   expenseId,
		group_id:     groupId,
		payment_by:   paymentBy,
		category_id:  categoryId,
		amount:       amt,
		nominal:      nom,
		payment_date: paymentDate,
		description:  desc,
		deleted_at:   nil,
		created_at:   createdAt,
		updated_at:   updatedAt,
	}
}

func (ea *ExpenseAggregate) GetExpenseId() expenseid.ExpenseId {
	return ea.expense_id
}

func (ea *ExpenseAggregate) GetGroupId() groupid.GroupId {
	return ea.group_id
}

func (ea *ExpenseAggregate) GetPaymentBy() userid.UserId {
	return ea.payment_by
}

func (ea *ExpenseAggregate) GetCategoryId() categoryid.CategoryId {
	return ea.category_id
}

func (ea *ExpenseAggregate) GetAmount() amount.Amount {
	return ea.amount
}

func (ea *ExpenseAggregate) GetNominal() nominal.Nominal {
	return ea.nominal
}

func (ea *ExpenseAggregate) GetPaymentDate() paymentdate.PaymentDate {
	return ea.payment_date
}

func (ea *ExpenseAggregate) GetDescription() description.Description {
	return ea.description
}

func (ea *ExpenseAggregate) GetDeletedAt() *deletedat.DeletedAt {
	return ea.deleted_at
}

func (ea *ExpenseAggregate) GetCreatedAt() createdat.CreatedAt {
	return ea.created_at
}

func (ea *ExpenseAggregate) GetUpdatedAt() updatedat.UpdatedAt {
	return ea.updated_at
}

func (ea *ExpenseAggregate) Delete(e expenseid.ExpenseId) error {
	if ea.expense_id.GetValue() == e.GetValue() {
		now := deletedat.NewDeletedAt()
		ea.deleted_at = &now
		return nil
	}

	return errors.New("削除対象の支出が見つかりませんでした。")
}
