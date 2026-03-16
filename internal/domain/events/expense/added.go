package expense

import (
	"lovers/internal/domain/models/event"
	"lovers/internal/domain/models/expense/afterdata"
	"lovers/internal/domain/models/expense/beforedata"
	"lovers/internal/domain/models/expense/expenseid"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/user/userid"
)

type ExpenseAdded struct {
	eventId    event.EventId
	occurredAt event.OccurredAt
	expenseId  expenseid.ExpenseId
	groupId    groupid.GroupId
	userId     userid.UserId
	operation  string
	beforeData *beforedata.BeforeData
	afterData  *afterdata.AfterData
}

func NewExpenseAdded(
	expenseId expenseid.ExpenseId,
	groupId groupid.GroupId,
	userId userid.UserId,
	operation string,
	afterdata *afterdata.AfterData,
) (*ExpenseAdded, error) {
	id, err := event.NewEventId()
	if err != nil {
		return nil, err
	}

	occ := event.NewOccurredAt()
	return &ExpenseAdded{
		eventId:    id,
		occurredAt: occ,
		expenseId:  expenseId,
		groupId:    groupId,
		userId:     userId,
		operation:  operation,
		beforeData: nil,
		afterData:  afterdata,
	}, nil
}

func (a *ExpenseAdded) EventId() event.EventId {
	return a.eventId
}

func (a *ExpenseAdded) OccurredAt() event.OccurredAt {
	return a.occurredAt
}

func (a *ExpenseAdded) ExpenseId() expenseid.ExpenseId {
	return a.expenseId
}

func (a *ExpenseAdded) GroupId() groupid.GroupId {
	return a.groupId
}

func (a *ExpenseAdded) UserId() userid.UserId {
	return a.userId
}

func (a *ExpenseAdded) Operation() string {
	return a.operation
}

func (a *ExpenseAdded) BeforeData() *beforedata.BeforeData {
	return a.beforeData
}

func (a *ExpenseAdded) AfterData() *afterdata.AfterData {
	return a.afterData
}
