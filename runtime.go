package main

import "fmt"

var RuntimeScope = map[string]WokData{
	"print": NWF("print", RuntimePrint),
}

func NWF(name string, function Function) *WokFunction {
	return &WokFunction{dtype: WokTypeFunction, name: name, function: function}
}

func RuntimePrint(interpreter *Interpreter, expressions []Expression) WokData {
	for _, expr := range expressions[1:] {
		fmt.Println(interpreter.Evaluate(expr).ToString())
	}
	return NewWokNull()
}
