package woklang

import "strconv"

type WokFloat struct {
	Value float64
	DType int
}

func NewWokFloat(value float64) *WokFloat {
	return &WokFloat{Value: value, DType: WokTypeFloat}
}

func (data *WokFloat) ToString() string {
	return strconv.FormatFloat(data.Value, 'E', -1, 64)
}

func (data *WokFloat) ToBoolean() bool {
	return data.Value != 0
}

func (data *WokFloat) ToInteger() int64 {
	return int64(data.Value)
}

func (data *WokFloat) ToFloat() float64 {
	return data.Value
}

func (data *WokFloat) GetType() int {
	return data.DType
}

func (data *WokFloat) GetTypeName() string {
	return "float"
}

func (data *WokFloat) GetValue() interface{} {
	return data.Value
}

func (data *WokFloat) Equals(other WokData) bool {
	return data.GetType() == other.GetType() && data.Value == other.GetValue()
}
