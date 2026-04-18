package password

import (
	"errors"
	"regexp"
	"strings"
)

var (
	lowerRegex  = regexp.MustCompile(`[a-z]`)
	numberRegex = regexp.MustCompile(`[0-9]`)
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
	if strings.ContainsRune(v, ' ') {
		return errors.New("password must not contain spaces")
	}

	if len([]rune(v)) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if !lowerRegex.MatchString(v) {
		return errors.New("password must contain at least one lowercase letter")
	}

	if !numberRegex.MatchString(v) {
		return errors.New("password must contain at least one number")
	}

	return nil
}
