package main

type Parser struct {
	current     int
	tokens      []Token
	expressions []*Expression
}

func (parser *Parser) Parse(tokens []Token) []*Expression {
	parser.current = 0
	parser.expressions = make([]*Expression, 0)
	parser.tokens = tokens
	for !parser.Eof() {
		expr := parser.DeclarationStatement()
		parser.expressions = append(parser.expressions, &expr)
	}
	return parser.expressions
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

func (parser *Parser) DeclarationStatement() Expression {
	return parser.ExpressionStatement()
}

func (parser *Parser) ExpressionStatement() Expression {
	return parser.EqualityExpression()
}

func (parser *Parser) EqualityExpression() Expression {
	expr := parser.ComparisonExpression()
	for parser.Match(TokenTypeEqualEqual) {
		oprtr := parser.Previous()
		right := parser.ComparisonExpression()
		expr = MakeExpressionBinary(expr, oprtr, right)

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
		expr = MakeExpressionBinary(expr, oprtr, right)
	}
	return expr
}

func (parser *Parser) AdditionExpression() Expression {
	expr := parser.MultiplicationExpression()
	for parser.Match(TokenTypePlus, TokenTypeMinus) {
		oprtr := parser.Previous()
		right := parser.MultiplicationExpression()
		expr = MakeExpressionBinary(expr, oprtr, right)
	}
	return expr
}

func (parser *Parser) MultiplicationExpression() Expression {
	expr := parser.PrimaryExpression()
	for parser.Match(TokenTypeSlash, TokenTypeStar) {
		oprtr := parser.Previous()
		right := parser.PrimaryExpression()
		expr = MakeExpressionBinary(expr, oprtr, right)
	}
	return expr
}

func (parser *Parser) PrimaryExpression() Expression {
	token := parser.Consume("Identifier or value expected", TokenTypeNumber)
	return MakeExpressionValue(token)
}