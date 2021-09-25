package woklang

func RuntimeNegation(interpreter *Interpreter, expressions []Expression) WokData {
	result := interpreter.Evaluate(expressions[0])
	return NewWokBoolean(!result.ToBoolean())
}
