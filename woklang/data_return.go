package woklang

type WokReturn struct {
	From  string
	Value WokData
	DType int
}

func NewWokReturn(from string, value WokData) *WokReturn {
	return &WokReturn{DType: WokTypeReturn, From: from, Value: value}
}

func (data *WokReturn) ToString() string {
	return data.Value.ToString()
}

func (data *WokReturn) ToBoolean() bool {
	return data.Value.ToBoolean()
}

func (data *WokReturn) ToInteger() int64 {
	return data.Value.ToInteger()
}

func (data *WokReturn) ToFloat() float64 {
	return data.Value.ToFloat()
}

func (data *WokReturn) GetType() int {
	return data.DType
}

func (data *WokReturn) GetTypeName() string {
	return "return"
}

func (data *WokReturn) GetValue() interface{} {
	return data.Value
}

func (data *WokReturn) Equals(other WokData) bool {
	return other.GetType() == WokTypeFunction && data.GetValue() == other.GetValue()
}
