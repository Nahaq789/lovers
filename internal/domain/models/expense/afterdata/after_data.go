package afterdata

import (
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/amount"
	"lovers/internal/domain/models/valueobjects/nominal"
)

type AfterData struct {
	userId  userid.UserId
	nominal nominal.Nominal
	amount  amount.Amount
}

func NewAfterData(u userid.UserId, n nominal.Nominal, a amount.Amount) *AfterData {
	return &AfterData{
		userId:  u,
		nominal: n,
		amount:  a,
	}
}

func (a *AfterData) UserId() userid.UserId {
	return a.userId
}

func (a *AfterData) Nominal() nominal.Nominal {
	return a.nominal
}

func (a *AfterData) Amount() amount.Amount {
	return a.amount
}
