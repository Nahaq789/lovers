package expense

type ExpenseCreateDto struct {
	GroupId        string          `json:"group_id"`
	CategoryId     string          `json:"category_id"`
	Nominal        string          `json:"nominal"`
	Description    string          `json:"description"`
	PaymentDetails []PaymentDetail `json:"payment_details"`
	PaymentDate    string          `json:"payment_date"`
}

type PaymentDetail struct {
	UserId string `json:"user_id"`
	Amount int    `json:"amount"`
}
