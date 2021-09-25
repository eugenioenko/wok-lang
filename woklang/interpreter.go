package woklang

import (
	"fmt"
	"os"
)

type Interpreter struct {
	Root    *Scope
	Scope   *Scope
	Runtime *Scope
}

func MakeInterpreter() Interpreter {
	interpreter := Interpreter{}
	interpreter.Runtime = &Scope{parent: nil, values: RuntimeScope}
	interpreter.Root = NewScope(nil)
	interpreter.Scope = NewScope(interpreter.Root)
	return interpreter
}

func (interpreter *Interpreter) Interpret(statements []Expression) WokData {
	var result WokData = NewWokNull()
	for _, statement := range statements {
		result = interpreter.Evaluate(statement)
	}
	return result
}

func (interpreter *Interpreter) Evaluate(expr Expression) WokData {
	return expr.Accept(interpreter)
}

func (interpreter *Interpreter) Error(errorMessage string) {
	fmt.Println("[Runtime Error] " + errorMessage)
	os.Exit(1)
}

func (interpreter *Interpreter) VisitExpressionList(expr *ExpressionList) WokData {
	if len(expr.List) == 0 {
		return NewWokNull()
	}
	callee := interpreter.Evaluate(expr.List[0])
	if callee.GetType() == WokTypeCallable {
		function := callee.GetValue().(Callable)
		return function(interpreter, expr.List[1:])
	}
	if callee.GetType() == WokTypeFunction {
		return interpreter.FunctionCall(callee.(*WokFunction), expr.List[1:])
	}
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionAtom(expr *ExpressionAtom) WokData {
	literal := expr.Atom.Literal

	switch expr.Atom.Type {
	case TokenTypeNull:
		return NewWokNull()
	case TokenTypeTrue:
		return NewWokBoolean(true)
	case TokenTypeFalse:
		return NewWokBoolean(false)
	case TokenTypeString:
		return NewWokString(literal)
	case TokenTypeInteger:
		return NewWokInteger(NewWokString(literal).ToInteger())
	case TokenTypeFloat:
		return NewWokFloat(NewWokString(literal).ToFloat())
	case TokenTypeBoolean:
		return NewWokBoolean(NewWokString(literal).ToBoolean())
	case TokenTypeIdentifier:
		scopeValue, ok := interpreter.Scope.Get(literal)
		if ok {
			return scopeValue
		}
		interpreter.Error(fmt.Sprintf("Undefined '%s'", literal))
		return NewWokNull()
	case TokenTypeReserved:
		runtimeValue, ok := interpreter.Runtime.Get(literal)
		if ok {
			return runtimeValue
		}
		interpreter.Error(fmt.Sprintf("Undefined predicate '%s'", literal))
		return NewWokNull()
	}
	return NewWokNull()
}
