package main

type Parser struct {
	current    int
	tokens     []Token
	statements []Statement
}

func (parser *Parser) Parse(tokens []Token) []Statement {
	parser.current = 0
	parser.statements = make([]Statement, 0)
	parser.tokens = tokens
	for !parser.Eof() {
		stmt := parser.DeclarationStatement()
		parser.statements = append(parser.statements, stmt)
	}
	return parser.statements
}

func MakeParser() Parser {
	return Parser{}
}

func (parser *Parser) Match(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		if parser.tokens[parser.current].ttype == tokenType {
			parser.Advance()
			return true
		}
	}
	return false
}

func (parser *Parser) Check(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		currentType := parser.tokens[parser.current].ttype
		if currentType == tokenType {
			return true
		}
	}
	return false
}

func (parser *Parser) Consume(errorMessage string, tokenTypes ...TokenType) Token {
	if parser.Check(tokenTypes...) {
		return parser.Advance()
	}
	panic(errorMessage)
}

func (parser *Parser) Advance() Token {
	if !parser.Eof() {
		parser.current += 1
	}
	return parser.Previous()
}

func (parser *Parser) Previous() Token {
	return parser.tokens[parser.current-1]
}

func (parser *Parser) Peek() Token {
	return parser.tokens[parser.current]
}

func (parser *Parser) Eof() bool {
	return parser.tokens[parser.current].ttype == TokenTypeEof ||
		parser.current >= len(parser.tokens)
}

func (parser *Parser) Error(token Token, errorMessage string) {
	panic(errorMessage)
}

//------------------------------------------------------------------------------
// AST STARTS HERE
//------------------------------------------------------------------------------
func (parser *Parser) DeclarationStatement() Statement {
	if parser.Match(TokenTypePrint) {
		return parser.PrintStatement()
	}
	return parser.ExpressionStatement()
}

func (parser *Parser) PrintStatement() Statement {
	expr := parser.AssignmentExpression()
	return NewStatementPrint(expr)
}

func (parser *Parser) ExpressionStatement() Statement {
	expr := parser.AssignmentExpression()
	return NewStatementExpression(expr)
}

func (parser *Parser) AssignmentExpression() Expression {
	expr := parser.EqualityExpression()
	if parser.Match(TokenTypeEqual) {
		value := parser.AssignmentExpression()
		expr = NewExpressionAssign(expr, value)
	}
	return expr
}

func (parser *Parser) EqualityExpression() Expression {
	expr := parser.ComparisonExpression()
	for parser.Match(TokenTypeEqualEqual) {
		oprtr := parser.Previous()
		right := parser.ComparisonExpression()
		expr = NewExpressionBinary(expr, oprtr, right)
	}
	return expr
}

func (parser *Parser) ComparisonExpression() Expression {
	expr := parser.AdditionExpression()
	for parser.Match(
		TokenTypeGreater, TokenTypeLess,
		TokenTypeGreaterEqual, TokenTypeLessEqual) {
		oprtr := parser.Previous()
		right := parser.AdditionExpression()
		expr = NewExpressionBinary(expr, oprtr, right)
	}
	return expr
}

func (parser *Parser) AdditionExpression() Expression {
	expr := parser.MultiplicationExpression()
	for parser.Match(TokenTypePlus, TokenTypeMinus) {
		oprtr := parser.Previous()
		right := parser.MultiplicationExpression()
		expr = NewExpressionBinary(expr, oprtr, right)
	}
	return expr
}

func (parser *Parser) MultiplicationExpression() Expression {
	expr := parser.UnaryExpression()
	for parser.Match(TokenTypeSlash, TokenTypeStar) {
		oprtr := parser.Previous()
		right := parser.UnaryExpression()
		expr = NewExpressionBinary(expr, oprtr, right)
	}
	return expr
}

func (parser *Parser) UnaryExpression() Expression {
	if parser.Match(TokenTypeBang, TokenTypeMinus, TokenTypePlus) {
		oprtr := parser.Previous()
		right := parser.PrimaryExpression()
		return NewExpressionUnary(oprtr, right)
	}
	return parser.PrimaryExpression()
}

func (parser *Parser) PrimaryExpression() Expression {
	switch {
	case parser.Match(TokenTypeInteger),
		parser.Match(TokenTypeFloat),
		parser.Match(TokenTypeString):
		return NewExpressionValue(parser.Previous())
	case parser.Match(TokenTypeIdentifier):
		return NewExpressionVariable(parser.Previous())
	}
	parser.Error(parser.Peek(), "Unexpected token")
	return nil
}
