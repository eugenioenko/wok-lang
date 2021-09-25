package woklang

import "strconv"

type WokString struct {
	Value string
	DType int
}

func NewWokString(value string) *WokString {
	return &WokString{Value: value, DType: WokTypeString}
}
func (data *WokString) ToString() string {
	return data.Value
}

func (data *WokString) ToBoolean() bool {
	return len(data.Value) > 0
}

func (data *WokString) ToInteger() int64 {
	value, err := strconv.ParseInt(data.Value, 10, 64)
	if err != nil {
		panic("Cant convert string to int")
	}
	return value
}

func (data *WokString) ToFloat() float64 {
	value, err := strconv.ParseFloat(data.Value, 64)
	if err != nil {
		panic("Cant convert string to float")
	}
	return value
}

func (data *WokString) GetType() int {
	return data.DType
}

func (data *WokString) GetTypeName() string {
	return "string"
}

func (data *WokString) GetValue() interface{} {
	return data.Value
}

func (data *WokString) Equals(other WokData) bool {
	return other.GetType() == data.DType && other.ToString() == data.Value
}
