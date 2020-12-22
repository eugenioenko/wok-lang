package main

type WokData interface {
	ToString() (string, bool)
	ToBoolean() (bool, bool)
	ToNumber() (int, bool)
	ToFloat() (float64, bool)
	Equals(other WokData) bool
	GetType() int
	GetValue() interface{}
}

const (
	// parser tokensA
	WokDataNull   = 0
	WokDataBool   = 1
	WokDataNumber = 2
	WokDataFloat  = 3
	WokDataString = 4
)

type DataString struct {
	value string
	class int
}

func (data *DataString) ToString() string {
	return data.value
}
