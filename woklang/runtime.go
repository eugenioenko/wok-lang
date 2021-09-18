package woklang

import "fmt"

var RuntimeScope = map[string]WokData{
	"print": WF("print", RuntimePrint),
	"cond":  WF("cond", RuntimeCond),
	"if":    WF("if", RuntimeIf),
	"+":     WF("+", RuntimeAddition),
	"*":     WF("*", RuntimeMultiplication),
	"-":     WF("-", RuntimeSubstraction),
}

func WF(name string, function Function) *WokFunction {
	return &WokFunction{dtype: WokTypeFunction, name: name, function: function}
}

func EvalParams(interpreter *Interpreter, expressions []Expression) []WokData {
	params := make([]WokData, len(expressions))
	for index, expression := range expressions {
		params[index] = interpreter.Evaluate(expression)
	}
	return params
}

func RuntimePrint(interpreter *Interpreter, expressions []Expression) WokData {
	var result WokData
	for _, expr := range expressions {
		result = interpreter.Evaluate(expr)
		fmt.Println(result.ToString())
	}
	return result
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

func RuntimeCond(interpreter *Interpreter, expressions []Expression) WokData {
	for _, expression := range expressions {
		condition := expression.(*ExpressionList).value
		if interpreter.Evaluate(condition[0]).ToBoolean() {
			return interpreter.Evaluate(condition[1])
		}
	}
	return NewWokNull()
}

func RuntimeIf(interpreter *Interpreter, expressions []Expression) WokData {
	expressionsCount := len(expressions)
	if expressionsCount > 0 {
		if interpreter.Evaluate(expressions[0]).ToBoolean() {
			if expressionsCount > 1 {
				return interpreter.Evaluate(expressions[1])
			}
		} else {
			if expressionsCount > 2 {
				return interpreter.Evaluate(expressions[2])
			}
		}
	}
	return NewWokNull()
}
