package woklang

func RuntimeEquality(interpreter *Interpreter, expressions []Expression) WokData {
	params := EvalParams(interpreter, expressions)
	result := Every(params, func(item WokData, index int) bool {
		return item.GetType() == params[0].GetType() && item.GetValue() == params[0].GetValue()
	})
	return NewWokBoolean(result)
}

func RuntimeInequality(interpreter *Interpreter, expressions []Expression) WokData {
	result := RuntimeEquality(interpreter, expressions)
	return NewWokBoolean(!result.ToBoolean())
}

func RuntimeNegation(interpreter *Interpreter, expressions []Expression) WokData {
	result := interpreter.Evaluate(expressions[0])
	return NewWokBoolean(!result.ToBoolean())
}
