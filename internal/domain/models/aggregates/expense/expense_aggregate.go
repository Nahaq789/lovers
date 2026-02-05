package expense

import (
	"errors"
	"lovers/internal/domain/models/category/categoryid"
	"lovers/internal/domain/models/expense/expenseid"
	paymentdetail "lovers/internal/domain/models/expense/paymentuser"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/deletedat"
	"lovers/internal/domain/models/valueobjects/description"
	"lovers/internal/domain/models/valueobjects/nominal"
	"lovers/internal/domain/models/valueobjects/paymentdate"
	"lovers/internal/domain/models/valueobjects/updatedat"
)

type ExpenseAggregate struct {
	expenseId   expenseid.ExpenseId
	groupId     groupid.GroupId
	categoryId  categoryid.CategoryId
	paymentUser paymentdetail.PaymentUser
	nominal     nominal.Nominal
	paymentDate paymentdate.PaymentDate
	description description.Description
	deletedAt   *deletedat.DeletedAt
	createdAt   createdat.CreatedAt
	updatedAt   updatedat.UpdatedAt
}

func NewExpenseAggregate(
	expenseId expenseid.ExpenseId,
	groupId groupid.GroupId,
	categoryId categoryid.CategoryId,
	paymentUser paymentdetail.PaymentUser,
	nom nominal.Nominal,
	paymentDate paymentdate.PaymentDate,
	desc description.Description,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt,
) *ExpenseAggregate {
	return &ExpenseAggregate{
		expenseId:   expenseId,
		groupId:     groupId,
		categoryId:  categoryId,
		paymentUser: paymentUser,
		nominal:     nom,
		paymentDate: paymentDate,
		description: desc,
		deletedAt:   nil,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func (ea *ExpenseAggregate) GetExpenseId() expenseid.ExpenseId {
	return ea.expenseId
}

func (ea *ExpenseAggregate) GetGroupId() groupid.GroupId {
	return ea.groupId
}

func (ea *ExpenseAggregate) GetCategoryId() categoryid.CategoryId {
	return ea.categoryId
}

func (ea *ExpenseAggregate) GetPaymentUser() paymentdetail.PaymentUser {
	return ea.paymentUser
}

func (ea *ExpenseAggregate) GetNominal() nominal.Nominal {
	return ea.nominal
}

func (ea *ExpenseAggregate) GetPaymentDate() paymentdate.PaymentDate {
	return ea.paymentDate
}

func (ea *ExpenseAggregate) GetDescription() description.Description {
	return ea.description
}

func (ea *ExpenseAggregate) GetDeletedAt() *deletedat.DeletedAt {
	return ea.deletedAt
}

func (ea *ExpenseAggregate) GetCreatedAt() createdat.CreatedAt {
	return ea.createdAt
}

func (ea *ExpenseAggregate) GetUpdatedAt() updatedat.UpdatedAt {
	return ea.updatedAt
}

func (ea *ExpenseAggregate) Delete(e expenseid.ExpenseId) error {
	if ea.expenseId.GetValue() == e.GetValue() {
		now := deletedat.NewDeletedAt()
		ea.deletedAt = &now
		return nil
	}

	return errors.New("削除対象の支出が見つかりませんでした。")
}
