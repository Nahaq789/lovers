package expense

import (
	"context"
	"lovers/internal/domain/models/aggregates/expense"
	"lovers/internal/domain/models/category/categoryid"
	"lovers/internal/domain/models/expense/expenseid"
	"lovers/internal/domain/models/expense/paymentuser"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/amount"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/description"
	"lovers/internal/domain/models/valueobjects/nominal"
	"lovers/internal/domain/models/valueobjects/paymentdate"
	"lovers/internal/domain/models/valueobjects/updatedat"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	expenseDto "lovers/internal/usecases/dto/expense"
	"lovers/internal/usecases/port"
	"lovers/internal/usecases/port/query"
)

type ExpenseAdd struct {
	expenseRepository repositories.ExpenseRepository
	groupQueryService query.GroupQueryService
	txManager         port.TransactionManager
}

func NewExpenseAdd(er repositories.ExpenseRepository, gq query.GroupQueryService, tm port.TransactionManager) *ExpenseAdd {
	return &ExpenseAdd{expenseRepository: er, groupQueryService: gq, txManager: tm}
}

func (ec *ExpenseAdd) Execute(ctx context.Context, d *expenseDto.ExpenseCreateDto) error {
	return ec.add(ctx, d)
}

func (ec *ExpenseAdd) add(ctx context.Context, d *expenseDto.ExpenseCreateDto) error {
	l := logger.FromContext(ctx)
	l.InfoContext(ctx, "明細作成処理を開始します。")

	expenseId, err := expenseid.NewExpenseId()
	if err != nil {
		l.ErrorContext(ctx, "明細IDの生成に失敗しました。", "error", err)
		return err
	}

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

	paymentUsers := make([]*paymentuser.PaymentUser, 0, len(d.PaymentDetails))
	for _, p := range d.PaymentDetails {
		userId, err := userid.NewUserIdFromString(p.UserId)
		if err != nil {
			l.ErrorContext(ctx, "ユーザIDの生成に失敗しました。", "error", err)
			return err
		}

		amount, err := amount.NewAmount(int64(p.Amount))
		if err != nil {
			l.ErrorContext(ctx, "金額の生成に失敗しました。", "error", err)
			return err
		}
		detail := paymentuser.NewExpensePaymentDetail(userId, amount)
		paymentUsers = append(paymentUsers, detail)
	}

	p := paymentuser.NewExpensePaymentUsers(paymentUsers)
	err = groupMembers.ValidateExpensePayments(p)
	if err != nil {
		return err
	}

	createdAt := createdat.NewCreatedAt()
	updatedAt := updatedat.NewUpdatedAt()

	expense, err := expense.NewExpenseAggregate(
		expenseId, groupId, categoryId, p, nominal, paymentDate, description, createdAt, updatedAt,
	)
	err = ec.expenseRepository.Add(ctx, expense)
	if err != nil {
		l.ErrorContext(ctx, "明細の作成に失敗しました。", "error", err)
		return err
	}

	l.InfoContext(ctx, "明細作成処理を終了します。")
	return nil
}
