package expense

import "lovers/internal/domain/repositories"

type ExpenseCreate struct {
	expenseRepository repositories.ExpenseRepository
}

func NewExpenseRepository(er repositories.ExpenseRepository) *ExpenseCreate {
	return &ExpenseCreate{expenseRepository: er}
}
