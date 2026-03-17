package expense

import "lovers/internal/domain/models/expense/afterdata"

type AfterDataJson struct {
	Nominal string `json:"nominal"`
	Amount  int64  `json:"amount"`
}

func NewAfterDataJson(a afterdata.AfterData) AfterDataJson {
	return AfterDataJson{
		Nominal: a.Nominal().GetValue(),
		Amount:  a.Amount().GetValue(),
	}
}
