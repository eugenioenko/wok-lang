package woklang

type WokException struct {
	Value string
	DType int
}

func NewWokException(value string) *WokException {
	return &WokException{Value: value, DType: WokTypeException}
}
func (data *WokException) ToString() string {
	return data.Value
}

func (data *WokException) ToBoolean() bool {
	return true
}

func (data *WokException) ToInteger() int64 {
	return 0
}

func (data *WokException) ToFloat() float64 {
	return 0.0
}

func (data *WokException) GetType() int {
	return data.DType
}

func (data *WokException) GetTypeName() string {
	return "exception"
}

func (data *WokException) GetValue() interface{} {
	return data.Value
}

func (data *WokException) Equals(other WokData) bool {
	return other.GetType() == data.DType && other.ToString() == data.Value
}
