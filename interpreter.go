package main

import "fmt"

type Interpreter struct {
	current    int
	statements []Statement
	scope      *Scope
}

func MakeInterpreter() Interpreter {
	return Interpreter{}
}

func (interpreter *Interpreter) Interpret(statements []Statement) {
	for _, statement := range interpreter.statements {
		result := interpreter.Execute(statement)
		fmt.Println(result.GetValue())
	}
}

func (interpreter *Interpreter) Execute(stmt Statement) WokData {
	return stmt.Accept(interpreter)
}

func (interpreter *Interpreter) Evaluate(expr Expression) WokData {
	return expr.Accept(interpreter)
}

func (interpreter *Interpreter) VisitExpressionAssign(expr *ExpressionAssign) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionBinary(expr *ExpressionBinary) WokData {
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
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionVariable(expr *ExpressionVariable) WokData {
	return NewWokNull()
}

func (interpreter *Interpreter) VisitStatementExpression(stmt *StatementExpression) WokData {
	return NewWokNull()
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
