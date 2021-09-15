package main

import "fmt"

func runtimePrint(interpreter *Interpreter, expressions []Expression) WokData {
	for _, expr := range expressions[1:] {
		fmt.Println(interpreter.Evaluate(expr).ToString())
	}
	return NewWokNull()
}
