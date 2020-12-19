package main

type DataType interface {
	ToString() string
	ToBoolean() bool
	ToNumber() int
	ToFloat() float64
	GetType() int
	GetValue()
}

const (
	// parser tokensA
	DataTypeNull = iota
	DataTypeBool
	DataTypeNumber
	DataTypeFloat
	DataTypeString
)

type DataString struct {
	value string
	class int
}

func (data *DataString) ToString() string {
	return data.value
}
