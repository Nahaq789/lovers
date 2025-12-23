package groupname

import "errors"

type GroupName struct {
	value string
}

func NewGroupName(v string) (GroupName, error) {
	if len(v) > 20 {
		return GroupName{}, errors.New("グループ名は20文字以内にしてください。")
	}

	return GroupName{value: v}, nil
}

func (g GroupName) GetValue() string {
	return g.value
}
