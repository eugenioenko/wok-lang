package woklang

type Callable func(*Interpreter, []Expression) WokData

type WokCallable struct {
	function Callable
	name     string
	dtype    int
}

func NewWokCallable(name string, function Callable) *WokCallable {
	return &WokCallable{dtype: WokTypeFunction, name: name, function: function}
}

func (data *WokCallable) ToString() string {
	return data.name
}

func (data *WokCallable) ToBoolean() bool {
	return true
}

func (data *WokCallable) ToInteger() int64 {
	return 0
}

func (data *WokCallable) ToFloat() float64 {
	return 0
}

func (data *WokCallable) GetType() int {
	return data.dtype
}

func (data *WokCallable) GetTypeName() string {
	return "function"
}

func (data *WokCallable) GetValue() interface{} {
	return data.function
}

func (data *WokCallable) Equals(other WokData) bool {
	return other.GetType() == WokTypeFunction && data.GetValue() == other.GetValue()
}
