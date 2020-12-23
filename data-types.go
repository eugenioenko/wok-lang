package main

import (
	"strconv"
)

type WokData interface {
	ToString() string
	ToBoolean() bool
	ToInteger() (int64, error)
	ToFloat() (float64, error)
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
	return data.value, nil
}

func (data *WokString) ToBoolean() bool {
	return len(data.value) > 0
}

func (data *WokString) ToInteger() (int64, error) {
	return strconv.ParseInt(data.value, 10, 64)
}

func (data *WokString) ToFloat() (float64, error) {
	return strconv.ParseFloat(data.value, 64)
}

func (data *WokString) GetType() int {
	return data.dtype
}

func (data *WokString) GetValue() string {
	return data.value
}

func (data *WokString) Equals(other WokData) bool {
	return other.GetType() == data.dtype && other.GetValue() == data.value
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

func (data *WokInteger) ToInteger() (int64, error) {
	return data.value, nil
}

func (data *WokInteger) ToFloat() (float64, error) {
	return float64(data.value), nil
}

func (data *WokInteger) GetType() int {
	return data.dtype
}

func (data *WokInteger) GetValue() int64 {
	return data.value
}

func (data *WokInteger) Equals(other WokData) bool {
	return other.GetType() == data.dtype && other.GetValue() == data.value
}
