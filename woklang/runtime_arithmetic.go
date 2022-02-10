package woklang

func RuntimeAssignment(interpreter *Interpreter, expressions []Expression) WokData {
	value := interpreter.Evaluate(expressions[1])
	token := expressions[0].(*ExpressionAtom).Atom.Literal
	interpreter.Scope.Set(token, value)
	return value
}

func RuntimeAddition(interpreter *Interpreter, expressions []Expression) WokData {
	params := EvalParams(interpreter, expressions)
	count := MathReduce(params, func(total int64, item WokData, index int) int64 {
		total += item.ToInteger()
		return total
	}, 0)
	return NewWokInteger(count)
}

func RuntimeMultiplication(interpreter *Interpreter, expressions []Expression) WokData {
	params := EvalParams(interpreter, expressions)
	count := MathReduce(params, func(total int64, item WokData, index int) int64 {
		total *= item.ToInteger()
		return total
	}, 0)
	return NewWokInteger(count)
}

func RuntimeSubstraction(interpreter *Interpreter, expressions []Expression) WokData {
	params := EvalParams(interpreter, expressions)
	if len(params) == 1 {
		return NewWokInteger(-params[0].ToInteger())
	}
	count := MathReduce(params[1:], func(total int64, item WokData, index int) int64 {
		total -= item.ToInteger()
		return total
	}, params[0].ToInteger())
	return NewWokInteger(count)
}
