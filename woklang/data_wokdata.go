package woklang

type WokData interface {
	ToString() string
	ToBoolean() bool
	ToInteger() int64
	ToFloat() float64
	Equals(other WokData) bool
	GetType() int
	GetTypeName() string
	GetValue() interface{}
}

const (
	WokTypeNull      = 0
	WokTypeBoolean   = 1
	WokTypeInteger   = 2
	WokTypeFloat     = 3
	WokTypeString    = 4
	WokTypeCallable  = 5
	WokTypeFunction  = 6
	WokTypeReturn    = 7
	WokTypeException = 8
)
