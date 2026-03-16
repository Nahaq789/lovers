package afterdata

import (
	"lovers/internal/domain/models/valueobjects/amount"
	"lovers/internal/domain/models/valueobjects/nominal"
)

type AfterData struct {
	nominal nominal.Nominal
	amount  amount.Amount
}

func NewAfterData(n nominal.Nominal, a amount.Amount) *AfterData {
	return &AfterData{
		nominal: n,
		amount:  a,
	}
}

func (a *AfterData) Nominal() nominal.Nominal {
	return a.nominal
}

func (a *AfterData) Amount() amount.Amount {
	return a.amount
}
