package woklang

type WokBoolean struct {
	Value bool
	DType int
}

func NewWokBoolean(value bool) *WokBoolean {
	return &WokBoolean{Value: value, DType: WokTypeBoolean}
}
func (data *WokBoolean) ToString() string {
	if data.Value {
		return "true"
	}
	return "false"
}

func (data *WokBoolean) ToBoolean() bool {
	return data.Value
}

func (data *WokBoolean) ToInteger() int64 {
	// TODO: confirm if its better to return error here instead of a value
	if data.Value {
		return 1
	}
	return 0
}

func (data *WokBoolean) ToFloat() float64 {
	// TODO: confirm if its better to return error here instead of a value
	if data.Value {
		return 1
	}
	return 0
}

func (data *WokBoolean) GetType() int {
	return data.DType
}

func (data *WokBoolean) GetTypeName() string {
	return "boolean"
}

func (data *WokBoolean) GetValue() interface{} {
	return data.Value
}

func (data *WokBoolean) Equals(other WokData) bool {
	return data.GetType() == other.GetType() && data.Value == other.GetValue()
}
