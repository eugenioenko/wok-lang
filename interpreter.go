package main

import (
	"fmt"
	"os"
	"strconv"
)

type Interpreter struct {
	current    int
	statements []Statement
	scope      *Scope
}

func MakeInterpreter() Interpreter {
	interpreter := Interpreter{}
	interpreter.scope = NewScope(nil)
	return interpreter
}

func (interpreter *Interpreter) Interpret(statements []Statement) {
	for _, statement := range statements {
		interpreter.Execute(statement)
	}
}

func (interpreter *Interpreter) Execute(stmt Statement) WokData {
	return stmt.Accept(interpreter)
}

func (interpreter *Interpreter) Evaluate(expr Expression) WokData {
	return expr.Accept(interpreter)
}

func (interpreter *Interpreter) Error(errorMessage string) {
	fmt.Println("[Runtime Error] " + errorMessage)
	os.Exit(1)
}

func (interpreter *Interpreter) VisitExpressionAssign(expr *ExpressionAssign) WokData {
	value := interpreter.Evaluate(expr.value)
	name := expr.name.literal
	interpreter.scope.Set(name, value)
	return value
}

func (interpreter *Interpreter) VisitExpressionBinary(expr *ExpressionBinary) WokData {

	left := interpreter.Evaluate(expr.left)
	right := interpreter.Evaluate(expr.right)

	if left.GetType() != right.GetType() {
		error := fmt.Sprintf("Invalid binary expression. No implicit casting allowed. %s %s %s", left.GetTypeName(), expr.operator.literal, right.GetTypeName())
		interpreter.Error(error)
		return NewWokNull()
	}
	if left.GetType() == WokTypeInteger {
		switch expr.operator.ttype {
		case TokenTypePlus:
			return NewWokInteger(left.ToInteger() + right.ToInteger())
		case TokenTypeMinus:
			return NewWokInteger(left.ToInteger() - right.ToInteger())
		case TokenTypeStar:
			return NewWokInteger(left.ToInteger() * right.ToInteger())
		case TokenTypeSlash:
			return NewWokInteger(left.ToInteger() / right.ToInteger())
		}
	}

	if left.GetType() == WokTypeFloat {
		switch expr.operator.ttype {
		case TokenTypePlus:
			return NewWokFloat(left.ToFloat() + right.ToFloat())
		case TokenTypeMinus:
			return NewWokFloat(left.ToFloat() - right.ToFloat())
		case TokenTypeStar:
			return NewWokFloat(left.ToFloat() * right.ToFloat())
		case TokenTypeSlash:
			return NewWokFloat(left.ToFloat() / right.ToFloat())
		}
	}

	if left.GetType() == WokTypeString {
		switch expr.operator.ttype {
		case TokenTypePlus:
			return NewWokString(left.ToString() + right.ToString())
		}
		interpreter.Error("Only addition can be performed on strings")
		return NewWokNull()
	}

	interpreter.Error("Unknown binary operator: " + expr.operator.literal)
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionCall(expr *ExpressionCall) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionGrouping(expr *ExpressionGrouping) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionLiteral(expr *ExpressionLiteral) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionUnary(expr *ExpressionUnary) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionValue(expr *ExpressionValue) WokData {
	switch expr.value.ttype {
	case TokenTypeInteger:
		value, err := strconv.ParseInt(expr.value.literal, 10, 64)
		if err != nil {
			interpreter.Error(expr.value.literal + "is not a valid integer")
		}
		return NewWokInteger(value)
	case TokenTypeFloat:
		value, err := strconv.ParseFloat(expr.value.literal, 64)
		if err != nil {
			interpreter.Error(expr.value.literal + "is not a valid integer")
		}
		return NewWokFloat(value)
	case TokenTypeString:
		return NewWokString(expr.value.literal)
	}

	interpreter.Error(expr.value.literal + "is not a valid value")
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionVariable(expr *ExpressionVariable) WokData {
	value, ok := interpreter.scope.Get(expr.name.literal)
	if ok {
		return value
	}
	interpreter.Error(expr.name.literal + " is not defined")
	return nil
}

func (interpreter *Interpreter) VisitStatementExpression(stmt *StatementExpression) WokData {
	return interpreter.Evaluate(stmt.expr)
}

func (interpreter *Interpreter) VisitStatementFunc(stmt *StatementFunc) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitStatementIf(stmt *StatementIf) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitStatementReturn(stmt *StatementReturn) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitStatementVar(stmt *StatementVar) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitStatementWhile(stmt *StatementWhile) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitStatementPrint(stmt *StatementPrint) WokData {
	result := interpreter.Evaluate(stmt.value)
	fmt.Println(result.ToString())
	return result
}
