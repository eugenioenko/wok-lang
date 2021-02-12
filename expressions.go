package main

type Expression interface {
    Accept(visitor VisitorExpression) WokData
}

type VisitorExpression interface {
	VisitExpressionAssign(expr *ExpressionAssign) WokData
	VisitExpressionBinary(expr *ExpressionBinary) WokData
	VisitExpressionCall(expr *ExpressionCall) WokData
	VisitExpressionGrouping(expr *ExpressionGrouping) WokData
	VisitExpressionLiteral(expr *ExpressionLiteral) WokData
	VisitExpressionUnary(expr *ExpressionUnary) WokData
	VisitExpressionValue(expr *ExpressionValue) WokData
	VisitExpressionVariable(expr *ExpressionVariable) WokData
}

type ExpressionAssign struct {
    name Expression
    value Expression
}

func NewExpressionAssign(name Expression, value Expression) *ExpressionAssign {
	return &ExpressionAssign{name, value}
}

func (expr *ExpressionAssign) Accept (visitor VisitorExpression) WokData {
	return visitor.VisitExpressionAssign(expr)
}

type ExpressionBinary struct {
    left Expression
    operator Token
    right Expression
}

func NewExpressionBinary(left Expression, operator Token, right Expression) *ExpressionBinary {
	return &ExpressionBinary{left, operator, right}
}

func (expr *ExpressionBinary) Accept (visitor VisitorExpression) WokData {
	return visitor.VisitExpressionBinary(expr)
}

type ExpressionCall struct {
    callee Expression
    paren Token
    args []Expression
}

func NewExpressionCall(callee Expression, paren Token, args []Expression) *ExpressionCall {
	return &ExpressionCall{callee, paren, args}
}

func (expr *ExpressionCall) Accept (visitor VisitorExpression) WokData {
	return visitor.VisitExpressionCall(expr)
}

type ExpressionGrouping struct {
    expression Expression
}

func NewExpressionGrouping(expression Expression) *ExpressionGrouping {
	return &ExpressionGrouping{expression}
}

func (expr *ExpressionGrouping) Accept (visitor VisitorExpression) WokData {
	return visitor.VisitExpressionGrouping(expr)
}

type ExpressionLiteral struct {
    value WokData
}

func NewExpressionLiteral(value WokData) *ExpressionLiteral {
	return &ExpressionLiteral{value}
}

func (expr *ExpressionLiteral) Accept (visitor VisitorExpression) WokData {
	return visitor.VisitExpressionLiteral(expr)
}

type ExpressionUnary struct {
    operator Token
    right Expression
}

func NewExpressionUnary(operator Token, right Expression) *ExpressionUnary {
	return &ExpressionUnary{operator, right}
}

func (expr *ExpressionUnary) Accept (visitor VisitorExpression) WokData {
	return visitor.VisitExpressionUnary(expr)
}

type ExpressionValue struct {
    value Token
}

func NewExpressionValue(value Token) *ExpressionValue {
	return &ExpressionValue{value}
}

func (expr *ExpressionValue) Accept (visitor VisitorExpression) WokData {
	return visitor.VisitExpressionValue(expr)
}

type ExpressionVariable struct {
    value Token
}

func NewExpressionVariable(value Token) *ExpressionVariable {
	return &ExpressionVariable{value}
}

func (expr *ExpressionVariable) Accept (visitor VisitorExpression) WokData {
	return visitor.VisitExpressionVariable(expr)
}

