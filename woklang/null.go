package woklang

type WokNull struct {
	dtype int
}

func NewWokNull() *WokNull {
	return &WokNull{dtype: WokTypeNull}
}

func (data *WokNull) ToString() string {
	return "null"
}

func (data *WokNull) ToBoolean() bool {
	return false
}

func (data *WokNull) ToInteger() int64 {
	return 0
}

func (data *WokNull) ToFloat() float64 {
	return 0
}

func (data *WokNull) GetType() int {
	return data.dtype
}

func (data *WokNull) GetTypeName() string {
	return "null"
}

func (data *WokNull) GetValue() interface{} {
	panic("Cant GetValue of Null")
}

func (data *WokNull) Equals(other WokData) bool {
	return data.GetType() == other.GetType()
}
