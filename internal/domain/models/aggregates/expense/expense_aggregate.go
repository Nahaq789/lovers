package expense

import (
	"context"
	"fmt"
	"lovers/internal/domain/events"
	"lovers/internal/domain/events/expense"
	"lovers/internal/domain/models/category/categoryid"
	"lovers/internal/domain/models/expense/afterdata"
	"lovers/internal/domain/models/expense/expenseid"
	paymentdetail "lovers/internal/domain/models/expense/paymentuser"
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
	expenseId    expenseid.ExpenseId
	groupId      groupid.GroupId
	categoryId   categoryid.CategoryId
	amount       amount.Amount
	paymentUsers *paymentdetail.PaymentUsers
	nominal      nominal.Nominal
	paymentDate  paymentdate.PaymentDate
	description  description.Description
	deletedAt    *deletedat.DeletedAt
	createdAt    createdat.CreatedAt
	updatedAt    updatedat.UpdatedAt
	events       []expense.ExpenseDomainEvent
}

func NewExpenseAggregate(
	expenseId expenseid.ExpenseId,
	groupId groupid.GroupId,
	categoryId categoryid.CategoryId,
	paymentUsers *paymentdetail.PaymentUsers,
	nom nominal.Nominal,
	paymentDate paymentdate.PaymentDate,
	desc description.Description,
	createdAt createdat.CreatedAt,
	updatedAt updatedat.UpdatedAt,
) (*ExpenseAggregate, error) {
	amount, err := paymentUsers.TotalAmount()
	if err != nil {
		return nil, err
	}
	return &ExpenseAggregate{
		expenseId:    expenseId,
		groupId:      groupId,
		categoryId:   categoryId,
		amount:       amount,
		paymentUsers: paymentUsers,
		nominal:      nom,
		paymentDate:  paymentDate,
		description:  desc,
		deletedAt:    nil,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
		events:       []expense.ExpenseDomainEvent{},
	}, nil
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

func (ea *ExpenseAggregate) GetPaymentUsers() *paymentdetail.PaymentUsers {
	return ea.paymentUsers
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

	return fmt.Errorf("expense %s not found", e.GetValue())
}

func (ea *ExpenseAggregate) PublishExpenseAdded(ctx context.Context, subscriber events.EventSubscriber, userId userid.UserId) error {

	// ドメインイベント作成
	afterDataList := make([]afterdata.AfterData, 0)
	for _, payment := range ea.paymentUsers.GetPaymentUsers() {
		a := payment.GetAmount()
		afterData := afterdata.NewAfterData(payment.GetUserId(), ea.nominal, a)
		afterDataList = append(afterDataList, *afterData)
	}
	event, err := expense.NewExpenseAdded(ea.expenseId, ea.groupId, userId, afterDataList)
	if err != nil {
		return err
	}
	publisher := events.NewEventPublisher()

	// サブスクライバ登録
	publisher.Subscribe(subscriber)

	// イベント発行
	domainEventErr := publisher.Publish(ctx, event)
	if domainEventErr != nil {
		return domainEventErr
	}
	return nil
}
