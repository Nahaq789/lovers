package paymentdate

import "time"

type PaymentDate struct {
	value time.Time
}

func NewPaymentDate() PaymentDate {
	now := time.Now().UTC()
	return PaymentDate{value: now}
}

func (c PaymentDate) GetValue() time.Time {
	return c.value.UTC()
}

func (c PaymentDate) ToString() string {
	return c.value.UTC().Format(time.RFC3339)
}
