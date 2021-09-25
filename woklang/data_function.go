package woklang

type WokFunction struct {
	name  string
	args  []string
	body  []Expression
	DType int
}

func NewWokFunction(name string, args []string, body []Expression) *WokFunction {
	return &WokFunction{DType: WokTypeFunction, name: name, args: args, body: body}
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
	return data.DType
}

func (data *WokFunction) GetTypeName() string {
	return "function"
}

func (data *WokFunction) GetValue() interface{} {
	return data.body
}

func (data *WokFunction) Equals(other WokData) bool {
	return other.GetType() == WokTypeFunction && data.GetValue() == other.GetValue()
}
