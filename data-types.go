package main

type WokData interface {
	ToString() string
	ToBoolean() bool
	ToNumber() int
	ToFloat() float64
	GetType() int
	GetValue()
}

const (
	// parser tokensA
	WokDataNull = iota
	WokDataBool
	WokDataNumber
	WokDataFloat
	WokDataString
)

type DataString struct {
	value string
	class int
}

func (data *DataString) ToString() string {
	return data.value
}
