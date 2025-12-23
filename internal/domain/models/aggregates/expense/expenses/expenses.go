package expenses

import (
	"errors"
	"lovers/internal/domain/models/aggregates/expense"
	"lovers/internal/domain/models/expense/expenseid"
)

type Expenses struct {
	expenses []expense.ExpenseAggregate
}

func NewExpenses() *Expenses {
	return &Expenses{
		expenses: []expense.ExpenseAggregate{},
	}
}

func FromArray(l []expense.ExpenseAggregate) *Expenses {
	return &Expenses{
		expenses: l,
	}
}

func (es *Expenses) GetExpenses() []expense.ExpenseAggregate {
	return es.expenses
}

func (es *Expenses) AddExpense(e expense.ExpenseAggregate) {
	for _, expense := range es.expenses {
		if !expense.GetExpenseId().Equal(e.GetExpenseId()) {
			return
		}
	}

	es.expenses = append(es.expenses, e)
}

func (es *Expenses) RemoveExpense(e expenseid.ExpenseId) error {
	for i, expense := range es.expenses {
		if expense.GetExpenseId().Equal(e) {
			es.expenses = append(es.expenses[:i], es.expenses[i+1:]...)
			return nil
		}
	}

	return errors.New("削除対象の支出が見つかりませんでした。")
}
