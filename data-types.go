package main

import (
	"strconv"
)

type WokData interface {
	ToString() (string, error)
	ToBoolean() (bool, error)
	ToInteger() (int64, error)
	ToFloat() (float64, error)
	Equals(other WokData) bool
	GetType() int
	GetValue() interface{}
}

const (
	WokDataNull    = 0
	WokDataBool    = 1
	WokDataInteger = 2
	WokDataFloat   = 3
	WokDataString  = 4
)

type WokString struct {
	value string
	dtype int
}

func (data *WokString) ToString() (string, error) {
	return data.value, nil
}

func (data *WokString) ToBoolean() (bool, error) {
	return len(data.value) > 0, nil
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
