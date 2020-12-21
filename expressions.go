package main

type Expression interface {
    Accept(visitor VisitorExpression) DataType
}

type VisitorExpression interface {
	VisitExpressionAssign(expr ExpressionAssign) DataType
	VisitExpressionBinary(expr ExpressionBinary) DataType
	VisitExpressionCall(expr ExpressionCall) DataType
	VisitExpressionGrouping(expr ExpressionGrouping) DataType
	VisitExpressionLiteral(expr ExpressionLiteral) DataType
	VisitExpressionUnary(expr ExpressionUnary) DataType
	VisitExpressionValue(expr ExpressionValue) DataType
	VisitExpressionVariable(expr ExpressionVariable) DataType
}

type ExpressionAssign struct {
    name Token
    value Expression
}

func MakeExpressionAssign(name Token, value Expression) ExpressionAssign {
	return ExpressionAssign{name, value}
}

func (expr ExpressionAssign) Accept (visitor VisitorExpression) DataType {
	return visitor.VisitExpressionAssign(expr)
}

type ExpressionBinary struct {
    left Expression
    operator Token
    right Expression
}

func MakeExpressionBinary(left Expression, operator Token, right Expression) ExpressionBinary {
	return ExpressionBinary{left, operator, right}
}

func (expr ExpressionBinary) Accept (visitor VisitorExpression) DataType {
	return visitor.VisitExpressionBinary(expr)
}

type ExpressionCall struct {
    callee Expression
    paren Token
    args []Expression
}

func MakeExpressionCall(callee Expression, paren Token, args []Expression) ExpressionCall {
	return ExpressionCall{callee, paren, args}
}

func (expr ExpressionCall) Accept (visitor VisitorExpression) DataType {
	return visitor.VisitExpressionCall(expr)
}

type ExpressionGrouping struct {
    expression Expression
}

func MakeExpressionGrouping(expression Expression) ExpressionGrouping {
	return ExpressionGrouping{expression}
}

func (expr ExpressionGrouping) Accept (visitor VisitorExpression) DataType {
	return visitor.VisitExpressionGrouping(expr)
}

type ExpressionLiteral struct {
    value DataType
}

func MakeExpressionLiteral(value DataType) ExpressionLiteral {
	return ExpressionLiteral{value}
}

func (expr ExpressionLiteral) Accept (visitor VisitorExpression) DataType {
	return visitor.VisitExpressionLiteral(expr)
}

type ExpressionUnary struct {
    operator Token
    right Expression
}

func MakeExpressionUnary(operator Token, right Expression) ExpressionUnary {
	return ExpressionUnary{operator, right}
}

func (expr ExpressionUnary) Accept (visitor VisitorExpression) DataType {
	return visitor.VisitExpressionUnary(expr)
}

type ExpressionValue struct {
    value Token
}

func MakeExpressionValue(value Token) ExpressionValue {
	return ExpressionValue{value}
}

func (expr ExpressionValue) Accept (visitor VisitorExpression) DataType {
	return visitor.VisitExpressionValue(expr)
}

type ExpressionVariable struct {
    value Token
}

func MakeExpressionVariable(value Token) ExpressionVariable {
	return ExpressionVariable{value}
}

func (expr ExpressionVariable) Accept (visitor VisitorExpression) DataType {
	return visitor.VisitExpressionVariable(expr)
}

