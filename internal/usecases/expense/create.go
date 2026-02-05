package expense

import (
	"context"
	"lovers/internal/domain/models/category/categoryid"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/valueobjects/description"
	"lovers/internal/domain/models/valueobjects/nominal"
	"lovers/internal/domain/models/valueobjects/paymentdate"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	expenseDto "lovers/internal/usecases/dto/expense"
	"lovers/internal/usecases/port"
	"lovers/internal/usecases/port/query"
)

type ExpenseCreate struct {
	expenseRepository repositories.ExpenseRepository
	groupQueryService query.GroupQueryService
	txManager         port.TransactionManager
}

func NewExpenseRepository(er repositories.ExpenseRepository, gq query.GroupQueryService, tm port.TransactionManager) *ExpenseCreate {
	return &ExpenseCreate{expenseRepository: er, groupQueryService: gq, txManager: tm}
}

func (ec *ExpenseCreate) Execute(ctx context.Context, d *expenseDto.ExpenseCreateDto) error {
	l := logger.FromContext(ctx)
	l.InfoContext(ctx, "明細作成処理を開始します。")

	groupId, err := groupid.NewGroupIdFromString(d.GroupId)
	if err != nil {
		l.ErrorContext(ctx, "グループIDの取得に失敗しました。", "error", err)
		return err
	}

	categoryId, err := categoryid.NewCategoryIdFromString(d.CategoryId)
	if err != nil {
		l.ErrorContext(ctx, "カテゴリIDの取得に失敗しました。", "error", err)
		return err
	}

	nominal, err := nominal.NewNominal(d.Nominal)
	if err != nil {
		l.ErrorContext(ctx, "項目の取得に失敗しました。", "error", err)
		return err
	}

	description := description.NewDescription(d.Description)
	paymentDate, err := paymentdate.NewPaymentDateFromString(d.PaymentDate)
	if err != nil {
		l.ErrorContext(ctx, "支払日の取得に失敗しました。", "error", err)
		return err
	}

	groupMembers, err := ec.groupQueryService.FindMemberById(ctx, groupId)
	if err != nil {
		l.ErrorContext(ctx, "グループメンバーの取得に失敗しました。", "error", err)
		return err
	}

	return nil
}
