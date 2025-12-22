package template

import "errors"

type TemplateName struct {
	value string
}

func NewTemplateName(v string) (TemplateName, error) {
	if len(v) > 20 {
		return TemplateName{}, errors.New("テンプレート名は20文字以内にしてください。")
	}

	return TemplateName{
		value: v,
	}, nil
}

func (t TemplateName) GetValue() string {
	return t.value
}
