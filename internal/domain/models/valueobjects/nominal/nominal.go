package nominal

import "errors"

type Nominal struct {
	value string
}

func NewNominal(v string) (Nominal, error) {
	if len(v) > 15 {
		return Nominal{}, errors.New("名目は15文字以内にしてください。")
	}

	return Nominal{value: v}, nil
}

func (n Nominal) GetValue() string {
	return n.value
}
