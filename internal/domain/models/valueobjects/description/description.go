package description

type Description struct {
	value string
}

func NewDescription(v string) Description {
	return Description{value: v}
}

func (d Description) GetValue() string {
	return d.value
}
