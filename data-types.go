package main

import (
	"strconv"
)

type WokData interface {
	ToString() string
	ToBoolean() bool
	ToInteger() int64
	ToFloat() float64
	Equals(other WokData) bool
	GetType() int
	GetValue() interface{}
}

const (
	WokTypeNull    = 0
	WokTypeBool    = 1
	WokTypeInteger = 2
	WokTypeFloat   = 3
	WokTypeString  = 4
)

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

func (data *WokString) GetValue() interface{} {
	return data.value
}

func (data *WokString) Equals(other WokData) bool {
	return other.GetType() == data.dtype && other.ToString() == data.value
}

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

func (data *WokInteger) GetValue() interface{} {
	return data.value
}

func (data *WokInteger) Equals(other WokData) bool {
	// TODO: check if type asertion other.(int64) is faster then interface{} comparison
	// https://golang.org/ref/spec#Comparison_operators
	return data.GetType() == other.GetType() && data.value == other.GetValue()
}

type WokBoolean struct {
	value bool
	dtype int
}

func NewWokBoolean(value bool) *WokBoolean {
	return &WokBoolean{value: value, dtype: WokTypeString}
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

func (data *WokBoolean) GetValue() interface{} {
	return data.value
}

func (data *WokBoolean) Equals(other WokData) bool {
	return data.GetType() == other.GetType() && data.value == other.GetValue()
}
