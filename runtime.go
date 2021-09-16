package main

import "fmt"

var RuntimeScope = map[string]WokData{
	"print": WF("print", RuntimePrint),
	"+":     WF("+", RuntimeAddition),
	"*":     WF("*", RuntimeMultiplication),
	"-":     WF("-", RuntimeSubstraction),
}

func WF(name string, function Function) *WokFunction {
	return &WokFunction{dtype: WokTypeFunction, name: name, function: function}
}

func EvalParams(interpreter *Interpreter, expressions []Expression) []WokData {
	params := make([]WokData, len(expressions)-1)
	for index, expression := range expressions[1:] {
		params[index] = interpreter.Evaluate(expression)
	}
	return params
}

func RuntimePrint(interpreter *Interpreter, expressions []Expression) WokData {
	for _, expr := range expressions[1:] {
		fmt.Println(interpreter.Evaluate(expr).ToString())
	}
	return NewWokNull()
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
