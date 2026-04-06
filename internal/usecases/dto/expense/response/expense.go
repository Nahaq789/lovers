package response

import "time"

type ExpenseResponse struct {
	ExpenseId    string         `json:"expense_id"`
	GroupId      string         `json:"group_id"`
	CategoryId   string         `json:"category_id"`
	Total        int64          `json:"total"`
	Nominal      string         `json:"nominal"`
	PaymentDate  time.Time      `json:"payment_date"`
	Description  string         `json:"description"`
	DeletedAt    time.Time      `json:"deleted_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	PaymentUsers []PaymentUsers `json:"payment_users"`
}

type PaymentUsers struct {
	UserId string `json:"user_id"`
	Amount int64  `json:"amount"`
}
