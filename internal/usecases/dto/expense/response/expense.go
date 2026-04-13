package response

import (
	"lovers/internal/domain/models/aggregates/expense"
	"time"
)

type ExpenseResponse struct {
	ExpenseId    string         `json:"expense_id"`
	GroupId      string         `json:"group_id"`
	CategoryId   string         `json:"category_id"`
	Total        int64          `json:"total"`
	Nominal      string         `json:"nominal"`
	PaymentDate  time.Time      `json:"payment_date"`
	Description  string         `json:"description"`
	DeletedAt    *time.Time     `json:"deleted_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	PaymentUsers []PaymentUsers `json:"payment_users"`
}

type PaymentUsers struct {
	UserId string `json:"user_id"`
	Amount int64  `json:"amount"`
}

func NewExpenseResponse(agg *expense.ExpenseAggregate) ExpenseResponse {
	paymentUsers := make([]PaymentUsers, 0, len(agg.GetPaymentUsers().GetPaymentUsers()))
	for _, pu := range agg.GetPaymentUsers().GetPaymentUsers() {
		paymentUsers = append(paymentUsers, PaymentUsers{
			UserId: pu.GetUserId().GetValue(),
			Amount: pu.GetAmount().GetValue(),
		})
	}

	var deletedAt *time.Time
	if agg.GetDeletedAt() != nil {
		v := agg.GetDeletedAt().GetValue()
		deletedAt = &v
	}

	return ExpenseResponse{
		ExpenseId:    agg.GetExpenseId().GetValue(),
		GroupId:      agg.GetGroupId().GetValue(),
		CategoryId:   agg.GetCategoryId().GetValue(),
		Total:        agg.GetTotal().GetValue(),
		Nominal:      agg.GetNominal().GetValue(),
		PaymentDate:  agg.GetPaymentDate().GetValue(),
		Description:  agg.GetDescription().GetValue(),
		DeletedAt:    deletedAt,
		CreatedAt:    agg.GetCreatedAt().GetValue(),
		UpdatedAt:    agg.GetUpdatedAt().GetValue(),
		PaymentUsers: paymentUsers,
	}
}
