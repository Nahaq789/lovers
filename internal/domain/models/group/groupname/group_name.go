package groupname

import "errors"

type GroupName struct {
	value string
}

func NewGroupName(v string) (GroupName, error) {
	if validation(v) {
		return GroupName{}, errors.New("グループ名は0文字以上20文字以内にしてください。")
	}

	return GroupName{value: v}, nil
}

func (g GroupName) GetValue() string {
	return g.value
}

func validation(v string) bool {
	if len(v) > 20 || 1 > len(v) {
		return true
	}

	return false
}
