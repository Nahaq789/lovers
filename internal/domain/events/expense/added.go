package expense

import (
	"lovers/internal/domain/models/event"
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
}

func NewExpenseAdded(expenseId expenseid.ExpenseId) (*ExpenseAdded, error) {
	id, err := event.NewEventId()
	if err != nil {
		return nil, err
	}

	occ := event.NewOccurredAt()
	return &ExpenseAdded{
		eventId:    id,
		occurredAt: occ,
		expenseId:  expenseId,
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
