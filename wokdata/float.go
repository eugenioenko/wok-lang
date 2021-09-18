package wokdata

import "strconv"

type WokFloat struct {
	value float64
	dtype int
}

func NewWokFloat(value float64) *WokFloat {
	return &WokFloat{value: value, dtype: WokTypeFloat}
}

func (data *WokFloat) ToString() string {
	return strconv.FormatFloat(data.value, 'E', -1, 64)
}

func (data *WokFloat) ToBoolean() bool {
	return data.value != 0
}

func (data *WokFloat) ToInteger() int64 {
	return int64(data.value)
}

func (data *WokFloat) ToFloat() float64 {
	return data.value
}

func (data *WokFloat) GetType() int {
	return data.dtype
}

func (data *WokFloat) GetTypeName() string {
	return "float"
}

func (data *WokFloat) GetValue() interface{} {
	return data.value
}

func (data *WokFloat) Equals(other WokData) bool {
	return data.GetType() == other.GetType() && data.value == other.GetValue()
}
