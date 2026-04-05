package expense

import (
	"context"
	"lovers/internal/domain/models/expense/expenseid"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	"lovers/internal/usecases/dto/expense/response"
)

type ExpenseFindById struct {
	expenseRepository repositories.ExpenseRepository
}

func NewExpenseFindById(er repositories.ExpenseRepository) *ExpenseFindById {
	return &ExpenseFindById{
		expenseRepository: er,
	}
}

func (ef *ExpenseFindById) Execute(ctx context.Context, expenseId, targetUserId string) (*response.ExpenseResponse, error) {
	l := logger.FromContext(ctx)
	l.InfoContext(ctx, "明細取得処理を開始します。")

	e, err := expenseid.NewExpenseIdFromString(expenseId)
	if err != nil {
		l.ErrorContext(ctx, "明細IDの生成に失敗しました。", "error", err)
		return nil, err
	}

	t, err := userid.NewUserIdFromString(targetUserId)
	if err != nil {
		l.ErrorContext(ctx, "ユーザIDの生成に失敗しました。", "error", err)
		return nil, err
	}

	aggregate, err := ef.expenseRepository.FindById(ctx, e, t)
	if err != nil {
		l.ErrorContext(ctx, "明細の取得に失敗しました。", "error", err)
		return nil, err
	}

	return nil, nil
}
