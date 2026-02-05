package paymentuser

import (
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/amount"
)

type PaymentUser struct {
	userId userid.UserId
	amount amount.Amount
}

func NewExpensePaymentDetail(userId userid.UserId, amount amount.Amount) *PaymentUser {
	return &PaymentUser{
		userId: userId,
		amount: amount,
	}
}

func (e *PaymentUser) GetUserId() userid.UserId {
	return e.userId
}

func (e *PaymentUser) GetAmount() amount.Amount {
	return e.amount
}
