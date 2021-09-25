package woklang

import "strconv"

type WokInteger struct {
	Value int64
	DType int
}

func NewWokInteger(value int64) *WokInteger {
	return &WokInteger{Value: value, DType: WokTypeInteger}
}

func (data *WokInteger) ToString() string {
	return strconv.FormatInt(data.Value, 10)
}

func (data *WokInteger) ToBoolean() bool {
	return data.Value != 0
}

func (data *WokInteger) ToInteger() int64 {
	return data.Value
}

func (data *WokInteger) ToFloat() float64 {
	return float64(data.Value)
}

func (data *WokInteger) GetType() int {
	return data.DType
}

func (data *WokInteger) GetTypeName() string {
	return "integer"
}

func (data *WokInteger) GetValue() interface{} {
	return data.Value
}

func (data *WokInteger) Equals(other WokData) bool {
	return data.GetType() == other.GetType() && data.Value == other.GetValue()
}
