package password

import (
	"errors"
	"regexp"
)

type Password struct {
	value string
}

func NewPassword(v string) (Password, error) {
	if err := validatePassword(v); err != nil {
		return Password{}, err
	}
	return Password{value: v}, nil
}

func (p Password) GetValue() string {
	return p.value
}

func validatePassword(v string) error {
	if len(v) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	matched, _ := regexp.MatchString(`[a-z]`, v)
	if !matched {
		return errors.New("password must contain at least one lowercase letter")
	}

	matched, _ = regexp.MatchString(`[0-9]`, v)
	if !matched {
		return errors.New("password must contain at least one number")
	}

	return nil
}