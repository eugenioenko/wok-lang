package woklang

type Expression interface {
	Accept(visitor VisitorExpression) WokData
}

type VisitorExpression interface {
	VisitExpressionList(expr *ExpressionList) WokData
	VisitExpressionAtom(expr *ExpressionAtom) WokData
}

type ExpressionList struct {
	value []Expression
}

func NewExpressionList(value []Expression) *ExpressionList {
	return &ExpressionList{value}
}

func (expr *ExpressionList) Accept(visitor VisitorExpression) WokData {
	return visitor.VisitExpressionList(expr)
}

type ExpressionAtom struct {
	value Token
}

func NewExpressionAtom(value Token) *ExpressionAtom {
	return &ExpressionAtom{value}
}

func (expr *ExpressionAtom) Accept(visitor VisitorExpression) WokData {
	return visitor.VisitExpressionAtom(expr)
}
