package wokdata

type Function func(*Interpreter, []Expression) WokData

type WokFunction struct {
	function Function
	name     string
	dtype    int
}

func NewWokFunction(name string, function Function) *WokFunction {
	return &WokFunction{dtype: WokTypeFunction, name: name, function: function}
}

func (data *WokFunction) ToString() string {
	return data.name
}

func (data *WokFunction) ToBoolean() bool {
	return true
}

func (data *WokFunction) ToInteger() int64 {
	return 0
}

func (data *WokFunction) ToFloat() float64 {
	return 0
}

func (data *WokFunction) GetType() int {
	return data.dtype
}

func (data *WokFunction) GetTypeName() string {
	return "function"
}

func (data *WokFunction) GetValue() interface{} {
	return data.function
}

func (data *WokFunction) Equals(other WokData) bool {
	return other.GetType() == WokTypeFunction && data.GetValue() == other.GetValue()
}
