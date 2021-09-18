package wokdata

import "strconv"

type WokString struct {
	value string
	dtype int
}

func NewWokString(value string) *WokString {
	return &WokString{value: value, dtype: WokTypeString}
}
func (data *WokString) ToString() string {
	return data.value
}

func (data *WokString) ToBoolean() bool {
	return len(data.value) > 0
}

func (data *WokString) ToInteger() int64 {
	value, err := strconv.ParseInt(data.value, 10, 64)
	if err != nil {
		panic("Cant convert string to int")
	}
	return value
}

func (data *WokString) ToFloat() float64 {
	value, err := strconv.ParseFloat(data.value, 64)
	if err != nil {
		panic("Cant convert string to float")
	}
	return value
}

func (data *WokString) GetType() int {
	return data.dtype
}

func (data *WokString) GetTypeName() string {
	return "string"
}

func (data *WokString) GetValue() interface{} {
	return data.value
}

func (data *WokString) Equals(other WokData) bool {
	return other.GetType() == data.dtype && other.ToString() == data.value
}
