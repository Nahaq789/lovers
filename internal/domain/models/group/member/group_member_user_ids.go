package member

import (
	"fmt"
	paymentuser "lovers/internal/domain/models/expense/paymentuser"
	"lovers/internal/domain/models/user/userid"
)

type MemberUserIds struct {
	members []userid.UserId
}

func NewMemberUserIds(m []userid.UserId) *MemberUserIds {
	return &MemberUserIds{members: m}
}

func (gm *MemberUserIds) ValidateExpensePayments(details *paymentuser.PaymentUsers) error {
	paymentUserIds := make(map[userid.UserId]bool)
	for _, detail := range details.GetPaymentUsers() {
		paymentUserIds[detail.GetUserId()] = true
	}

	memberIds := make(map[userid.UserId]bool)
	for _, member := range gm.members {
		memberIds[member] = true
	}

	// 1. すべてのグループメンバーがPaymentDetailsに含まれているか
	for _, member := range gm.members {
		if !paymentUserIds[member] {
			return fmt.Errorf("member %s is not included in payment users", member.GetValue())
		}
	}

	// 2. PaymentDetailsのすべてのUserIdがグループメンバーか
	for _, detail := range details.GetPaymentUsers() {
		if !memberIds[detail.GetUserId()] {
			return fmt.Errorf("user %s is not a group member", detail.GetUserId().GetValue())
		}
	}

	return nil
}
