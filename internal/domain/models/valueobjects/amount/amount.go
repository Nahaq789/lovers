package amount

import "errors"

type Amount struct {
	value int64
}

func NewAmount(v int64) (Amount, error) {
	if err := validateAmount(v); err != nil {
		return Amount{}, err
	}
	return Amount{value: v}, nil
}

func validateAmount(v int64) error {
	if v < 0 {
		return errors.New("金額は0円以上にしてください。")
	}
	return nil
}

func (a Amount) GetValue() int64 {
	return a.value
}

func (a Amount) Add(v int64) (Amount, error) {
	if err := validateAmount(v); err != nil {
		return Amount{}, err
	}
	newValue := a.value + v
	return Amount{value: newValue}, nil
}

func (a Amount) Subtract(v int64) (Amount, error) {
	if err := validateAmount(v); err != nil {
		return Amount{}, err
	}
	newValue := a.value - v
	if err := validateAmount(newValue); err != nil {
		return Amount{}, errors.New("引き算の結果が負の値になります。")
	}
	return Amount{value: newValue}, nil
}
