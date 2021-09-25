package woklang

type WokBoolean struct {
	value bool
	dtype int
}

func NewWokBoolean(value bool) *WokBoolean {
	return &WokBoolean{value: value, dtype: WokTypeBoolean}
}
func (data *WokBoolean) ToString() string {
	if data.value {
		return "true"
	}
	return "false"
}

func (data *WokBoolean) ToBoolean() bool {
	return data.value
}

func (data *WokBoolean) ToInteger() int64 {
	// TODO: confirm if its better to return error here instead of a value
	if data.value {
		return 1
	}
	return 0
}

func (data *WokBoolean) ToFloat() float64 {
	// TODO: confirm if its better to return error here instead of a value
	if data.value {
		return 1
	}
	return 0
}

func (data *WokBoolean) GetType() int {
	return data.dtype
}

func (data *WokBoolean) GetTypeName() string {
	return "boolean"
}

func (data *WokBoolean) GetValue() interface{} {
	return data.value
}

func (data *WokBoolean) Equals(other WokData) bool {
	return data.GetType() == other.GetType() && data.value == other.GetValue()
}
