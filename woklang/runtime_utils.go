package woklang

func WF(name string, function Callable) *WokCallable {
	return &WokCallable{DType: WokTypeCallable, name: name, function: function}
}

func EvalParams(interpreter *Interpreter, expressions []Expression) []WokData {
	params := make([]WokData, len(expressions))
	for index, expression := range expressions {
		params[index] = interpreter.Evaluate(expression)
	}
	return params
}

func RuntimeDebug(interpreter *Interpreter, expressions []Expression) WokData {
	return interpreter.Evaluate(expressions[0])
}
