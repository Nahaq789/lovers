package username

import "errors"

type UserName struct {
	value string
}

func NewUserName(v string) (*UserName, error) {
	if len(v) > 20 {
		return nil, errors.New("ユーザー名は20文字以内にしてください。")
	}
	return &UserName{
		value: v,
	}, nil
}
