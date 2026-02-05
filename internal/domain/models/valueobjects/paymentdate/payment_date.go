package paymentdate

import (
	"fmt"
	"time"
)

const PaymentDateFormat = "2006-01-02"

type PaymentDate struct {
	value time.Time
}

func NewPaymentDate() PaymentDate {
	now := time.Now().UTC()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return PaymentDate{value: date}
}

func NewPaymentDateFromString(dateStr string) (PaymentDate, error) {
	t, err := time.Parse(PaymentDateFormat, dateStr)
	if err != nil {
		return PaymentDate{}, fmt.Errorf("invalid payment date format (expected YYYY-MM-DD): %w", err)
	}
	return PaymentDate{value: t.UTC()}, nil
}

func (c PaymentDate) GetValue() time.Time {
	return c.value.UTC()
}

func (c PaymentDate) ToString() string {
	return c.value.UTC().Format(PaymentDateFormat)
}
