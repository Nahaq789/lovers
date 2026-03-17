package paymentuser

import (
	"fmt"
	"lovers/internal/domain/models/valueobjects/amount"
)

type PaymentUsers struct {
	paymentUsers []*PaymentUser
}

func NewExpensePaymentUsers(paymentUsers []*PaymentUser) *PaymentUsers {
	return &PaymentUsers{paymentUsers: paymentUsers}
}

func (pu *PaymentUsers) GetPaymentUsers() []*PaymentUser {
	return pu.paymentUsers
}

func (pu *PaymentUsers) TotalAmount() (amount.Amount, error) {
	total := amount.NewAmountZero()
	for _, user := range pu.paymentUsers {
		var err error
		total, err = total.Add(user.GetAmount())
		if err != nil {
			return amount.Amount{}, fmt.Errorf("支払いユーザー %v の金額の加算に失敗しました: %w", user.GetUserId().GetValue(), err)
		}
	}
	return total, nil
}
