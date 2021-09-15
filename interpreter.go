package main

import (
	"fmt"
	"os"
)

type Interpreter struct {
	scope *Scope
	root  *Scope
}

func MakeInterpreter() Interpreter {
	interpreter := Interpreter{}
	interpreter.root = NewScope(nil)
	interpreter.scope = NewScope(interpreter.root)
	interpreter.root.Set("print", NewWokFunction("print", runtimePrint))
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
	if len(expr.value) == 0 {
		return NewWokNull()
	}
	function := interpreter.Evaluate(expr.value[0]).GetValue().(Function)
	function(interpreter, expr.value)
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionAtom(expr *ExpressionAtom) WokData {
	literal := expr.value.literal
	switch expr.value.ttype {
	case TokenTypeString:
		return NewWokString(literal)
	case TokenTypeInteger:
		return NewWokInteger(NewWokString(literal).ToInteger())
	case TokenTypeFloat:
		return NewWokFloat(NewWokString(literal).ToFloat())
	case TokenTypeBoolean:
		return NewWokBoolean(NewWokString(literal).ToBoolean())
	case TokenTypeIdentifier:
		rootValue, ok := interpreter.root.Get(literal)
		if ok {
			return rootValue
		}
		scopeValue, ok := interpreter.scope.Get(literal)
		if ok {
			return scopeValue
		}
		return NewWokNull()
	}
	return NewWokNull()
}
