package main

import (
	"fmt"
	"strconv"
)

type Interpreter struct {
	current    int
	statements []Statement
	scope      *Scope
}

func MakeInterpreter() Interpreter {
	return Interpreter{}
}

func (interpreter *Interpreter) Interpret(statements []Statement) {
	for i, statement := range statements {
		result := interpreter.Execute(statement)
		fmt.Print(":")
		fmt.Print(i)
		fmt.Print(" - ")
		fmt.Println(result.ToString())
	}
}

func (interpreter *Interpreter) Execute(stmt Statement) WokData {
	return stmt.Accept(interpreter)
}

func (interpreter *Interpreter) Evaluate(expr Expression) WokData {
	return expr.Accept(interpreter)
}

func (interpreter *Interpreter) Error(errorMessage string) {
	panic(errorMessage)
}

func (interpreter *Interpreter) VisitExpressionAssign(expr *ExpressionAssign) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionBinary(expr *ExpressionBinary) WokData {
	left := interpreter.Evaluate(expr.left)
	right := interpreter.Evaluate(expr.right)
	switch expr.operator.ttype {
	case TokenTypePlus:
		return NewWokInteger(left.ToInteger() + right.ToInteger())
	case TokenTypeMinus:
		return NewWokInteger(left.ToInteger() - right.ToInteger())
	case TokenTypeStar:
		return NewWokInteger(left.ToInteger() * right.ToInteger())
	case TokenTypeSlash:
		return NewWokInteger(left.ToInteger() / right.ToInteger())
	default:
		interpreter.Error("Unknown binary operator: " + expr.operator.literal)
		return NewWokNull()
	}
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
	value, err := strconv.ParseInt(expr.value.literal, 10, 64)
	if err != nil {
		interpreter.Error("string to int failed")
	}
	return NewWokInteger(value)
}

func (interpreter *Interpreter) VisitExpressionVariable(expr *ExpressionVariable) WokData {
	return NewWokNull()
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
