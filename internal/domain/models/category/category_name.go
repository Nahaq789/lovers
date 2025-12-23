package category

import "errors"

type CategoryName struct {
	value string
}

func NewCategoryName(v string) (CategoryName, error) {
	if len(v) > 10 {
		return CategoryName{}, errors.New("カテゴリー名は20文字以内にしてください。")
	}

	return CategoryName{value: v}, nil
}

func (c CategoryName) GetValue() string {
	return c.value
}
