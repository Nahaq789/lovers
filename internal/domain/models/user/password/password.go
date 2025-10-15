package password

import (
	"errors"
	"regexp"
)

type Password struct {
	value string
}

func NewPassword(v string) (*Password, error) {
	if err := validatePassword(v); err != nil {
		return nil, err
	}
	return &Password{value: v}, nil
}

func validatePassword(v string) error {
	if len(v) < 6 {
		return errors.New("パスワードは最低6文字必要です")
	}

	matched, _ := regexp.MatchString(`[a-z]`, v)
	if !matched {
		return errors.New("パスワードには小文字が必要です")
	}

	matched, _ = regexp.MatchString(`[0-9]`, v)
	if !matched {
		return errors.New("パスワードには数字が必要です")
	}

	return nil
}
