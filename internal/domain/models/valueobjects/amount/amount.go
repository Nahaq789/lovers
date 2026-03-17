package amount

import (
	"fmt"
)

type Amount struct {
	value int64
}

func NewAmount(v int64) (Amount, error) {
	if err := validateAmount(v); err != nil {
		return Amount{}, err
	}
	return Amount{value: v}, nil
}

func NewAmountZero() Amount {
	return Amount{value: 0}
}

func validateAmount(v int64) error {
	if v < 0 {
		return fmt.Errorf("amount must be non-negative, got %d", v)
	}
	return nil
}

func (a Amount) GetValue() int64 {
	return a.value
}

func (a Amount) Add(v Amount) (Amount, error) {
	if err := validateAmount(v.GetValue()); err != nil {
		return Amount{}, err
	}
	newValue := a.value + v.GetValue()
	return Amount{value: newValue}, nil
}

func (a Amount) Subtract(v Amount) (Amount, error) {
	if err := validateAmount(v.GetValue()); err != nil {
		return Amount{}, err
	}
	newValue := a.value - v.GetValue()
	if err := validateAmount(newValue); err != nil {
		return Amount{}, fmt.Errorf("subtraction result is negative: %d - %d = %d", a.value, v.GetValue(), newValue)
	}
	return Amount{value: newValue}, nil
}
