package beforedata

import (
	"lovers/internal/domain/models/valueobjects/amount"
	"lovers/internal/domain/models/valueobjects/nominal"
)

type BeforeData struct {
	nominal nominal.Nominal
	amount  amount.Amount
}

func NewBeforeData(n nominal.Nominal, a amount.Amount) *BeforeData {
	return &BeforeData{
		nominal: n,
		amount:  a,
	}
}

func (b *BeforeData) Nominal() nominal.Nominal {
	return b.nominal
}

func (b *BeforeData) Amount() amount.Amount {
	return b.amount
}
