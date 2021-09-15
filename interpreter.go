package main

import (
	"fmt"
	"os"
)

type Interpreter struct {
	statements []Expression
	scope      *Scope
}

func MakeInterpreter() Interpreter {
	interpreter := Interpreter{}
	interpreter.scope = NewScope(nil)
	return interpreter
}

func (interpreter *Interpreter) Interpret(statements []Expression) {
	for _, statement := range statements {
		interpreter.Evaluate(statement)
	}
}

func (interpreter *Interpreter) Evaluate(expr Expression) WokData {
	return expr.Accept(interpreter)
}

func (interpreter *Interpreter) Error(errorMessage string) {
	fmt.Println("[Runtime Error] " + errorMessage)
	os.Exit(1)
}

func (interpreter *Interpreter) VisitExpressionList(expr *ExpressionList) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionAtom(expr *ExpressionAtom) WokData {
	return NewWokNull()
}
