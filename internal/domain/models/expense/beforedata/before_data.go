package beforedata

import (
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/amount"
	"lovers/internal/domain/models/valueobjects/nominal"
)

type BeforeData struct {
	userId  userid.UserId
	nominal nominal.Nominal
	amount  amount.Amount
}

func NewBeforeData(u userid.UserId, n nominal.Nominal, a amount.Amount) *BeforeData {
	return &BeforeData{
		userId:  u,
		nominal: n,
		amount:  a,
	}
}

func (b *BeforeData) UserId() userid.UserId {
	return b.userId
}

func (b *BeforeData) Nominal() nominal.Nominal {
	return b.nominal
}

func (b *BeforeData) Amount() amount.Amount {
	return b.amount
}
