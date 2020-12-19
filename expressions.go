package main

type Expression interface {
	Accept(visitor VisitorExpression) DataType
}

type VisitorExpression interface {
	VisitExpressionAssign(expr *ExpressionAssign) DataType
	VisitExpressionBinary(expr *ExpressionBinary) DataType
	VisitExpressionCall(expr *ExpressionCall) DataType
	VisitExpressionGrouping(expr *ExpressionGrouping) DataType
	VisitExpressionLiteral(expr *ExpressionLiteral) DataType
	VisitExpressionUnary(expr *ExpressionUnary) DataType
	VisitExpressionVariable(expr *ExpressionVariable) DataType
}

type ExpressionAssign struct {
	name  Token
	value *Expression
}

func NewExpressionAssign(name Token, value *Expression) *ExpressionAssign {
	return &ExpressionAssign{name, value}
}

func (expr *ExpressionAssign) Accept(visitor VisitorExpression) DataType {
	return visitor.VisitExpressionAssign(expr)
}

type ExpressionBinary struct {
	left     *Expression
	operator Token
	right    *Expression
}

func NewExpressionBinary(left *Expression, operator Token, right *Expression) *ExpressionBinary {
	return &ExpressionBinary{left, operator, right}
}

func (expr *ExpressionBinary) Accept(visitor VisitorExpression) DataType {
	return visitor.VisitExpressionBinary(expr)
}

type ExpressionCall struct {
	callee *Expression
	paren  Token
	args   []*Expression
}

func NewExpressionCall(callee *Expression, paren Token, args []*Expression) *ExpressionCall {
	return &ExpressionCall{callee, paren, args}
}

func (expr *ExpressionCall) Accept(visitor VisitorExpression) DataType {
	return visitor.VisitExpressionCall(expr)
}

type ExpressionGrouping struct {
	expression *Expression
}

func NewExpressionGrouping(expression *Expression) *ExpressionGrouping {
	return &ExpressionGrouping{expression}
}

func (expr *ExpressionGrouping) Accept(visitor VisitorExpression) DataType {
	return visitor.VisitExpressionGrouping(expr)
}

type ExpressionLiteral struct {
	value *DataType
}

func NewExpressionLiteral(value *DataType) *ExpressionLiteral {
	return &ExpressionLiteral{value}
}

func (expr *ExpressionLiteral) Accept(visitor VisitorExpression) DataType {
	return visitor.VisitExpressionLiteral(expr)
}

type ExpressionUnary struct {
	operator Token
	right    *Expression
}

func NewExpressionUnary(operator Token, right *Expression) *ExpressionUnary {
	return &ExpressionUnary{operator, right}
}

func (expr *ExpressionUnary) Accept(visitor VisitorExpression) DataType {
	return visitor.VisitExpressionUnary(expr)
}

type ExpressionVariable struct {
	name Token
}

func NewExpressionVariable(name Token) *ExpressionVariable {
	return &ExpressionVariable{name}
}

func (expr *ExpressionVariable) Accept(visitor VisitorExpression) DataType {
	return visitor.VisitExpressionVariable(expr)
}
