package wokdata

import "strconv"

type WokInteger struct {
	value int64
	dtype int
}

func NewWokInteger(value int64) *WokInteger {
	return &WokInteger{value: value, dtype: WokTypeInteger}
}

func (data *WokInteger) ToString() string {
	return strconv.FormatInt(data.value, 10)
}

func (data *WokInteger) ToBoolean() bool {
	return data.value != 0
}

func (data *WokInteger) ToInteger() int64 {
	return data.value
}

func (data *WokInteger) ToFloat() float64 {
	return float64(data.value)
}

func (data *WokInteger) GetType() int {
	return data.dtype
}

func (data *WokInteger) GetTypeName() string {
	return "integer"
}

func (data *WokInteger) GetValue() interface{} {
	return data.value
}

func (data *WokInteger) Equals(other WokData) bool {
	return data.GetType() == other.GetType() && data.value == other.GetValue()
}
