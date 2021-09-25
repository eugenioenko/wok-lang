package woklang

func RuntimeWhile(interpreter *Interpreter, expressions []Expression) WokData {
	var result WokData = NewWokNull()
	for interpreter.Evaluate(expressions[0]).ToBoolean() {
		result = interpreter.Evaluate(expressions[1])
	}
	return result
}
