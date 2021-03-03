package main

type Statement interface {
    Accept(visitor VisitorStatement) WokData
}

type VisitorStatement interface {
	VisitStatementExpression(stmt *StatementExpression) WokData
	VisitStatementFunc(stmt *StatementFunc) WokData
	VisitStatementIf(stmt *StatementIf) WokData
	VisitStatementReturn(stmt *StatementReturn) WokData
	VisitStatementVar(stmt *StatementVar) WokData
	VisitStatementWhile(stmt *StatementWhile) WokData
	VisitStatementPrint(stmt *StatementPrint) WokData
}

type StatementExpression struct {
    expr Expression
}

func NewStatementExpression(expr Expression) *StatementExpression {
	return &StatementExpression{expr}
}

func (stmt *StatementExpression) Accept (visitor VisitorStatement) WokData {
	return visitor.VisitStatementExpression(stmt)
}

type StatementFunc struct {
    name Token
    params []Token
    body []Statement
}

func NewStatementFunc(name Token, params []Token, body []Statement) *StatementFunc {
	return &StatementFunc{name, params, body}
}

func (stmt *StatementFunc) Accept (visitor VisitorStatement) WokData {
	return visitor.VisitStatementFunc(stmt)
}

type StatementIf struct {
    condition Expression
    thenStmt Statement
    elseStmt Statement
}

func NewStatementIf(condition Expression, thenStmt Statement, elseStmt Statement) *StatementIf {
	return &StatementIf{condition, thenStmt, elseStmt}
}

func (stmt *StatementIf) Accept (visitor VisitorStatement) WokData {
	return visitor.VisitStatementIf(stmt)
}

type StatementReturn struct {
    keyword Token
    value Expression
}

func NewStatementReturn(keyword Token, value Expression) *StatementReturn {
	return &StatementReturn{keyword, value}
}

func (stmt *StatementReturn) Accept (visitor VisitorStatement) WokData {
	return visitor.VisitStatementReturn(stmt)
}

type StatementVar struct {
    name Token
    dtype Token
    initial Expression
    writable bool
}

func NewStatementVar(name Token, dtype Token, initial Expression, writable bool) *StatementVar {
	return &StatementVar{name, dtype, initial, writable}
}

func (stmt *StatementVar) Accept (visitor VisitorStatement) WokData {
	return visitor.VisitStatementVar(stmt)
}

type StatementWhile struct {
    condition Expression
    loop Statement
}

func NewStatementWhile(condition Expression, loop Statement) *StatementWhile {
	return &StatementWhile{condition, loop}
}

func (stmt *StatementWhile) Accept (visitor VisitorStatement) WokData {
	return visitor.VisitStatementWhile(stmt)
}

type StatementPrint struct {
    value Expression
}

func NewStatementPrint(value Expression) *StatementPrint {
	return &StatementPrint{value}
}

func (stmt *StatementPrint) Accept (visitor VisitorStatement) WokData {
	return visitor.VisitStatementPrint(stmt)
}

