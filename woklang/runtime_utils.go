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

func RuntimeDefun(interpreter *Interpreter, expressions []Expression) WokData {
	name := expressions[0].(*ExpressionAtom).Atom.Literal
	args := make([]string, len(expressions[1].(*ExpressionList).List))
	for index, token := range expressions[1].(*ExpressionList).List {
		args[index] = token.(*ExpressionAtom).Atom.Literal
	}
	function := NewWokFunction(name, args, expressions[2:])
	interpreter.Scope.Set(name, function)
	return function
}
