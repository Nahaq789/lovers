package expense

import (
	"lovers/internal/domain/events"
	"lovers/internal/domain/models/expense/expenseid"
)

type ExpenseDomainEvent interface {
	events.Event
	ExpenseId() expenseid.ExpenseId
}
