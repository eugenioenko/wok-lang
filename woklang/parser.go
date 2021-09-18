package woklang

import (
	"fmt"
	"os"
)

type Parser struct {
	current     int
	tokens      []Token
	expressions []Expression
}

func (parser *Parser) Parse(tokens []Token) []Expression {
	parser.current = 0
	parser.expressions = make([]Expression, 0)
	parser.tokens = tokens
	for !parser.Eof() {
		stmt := parser.Statement()
		parser.expressions = append(parser.expressions, stmt)
	}
	return parser.expressions
}

func MakeParser() Parser {
	return Parser{}
}

func (parser *Parser) Match(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		if parser.Peek().ttype == tokenType {
			parser.Advance()
			return true
		}
	}
	return false
}

func (parser *Parser) Check(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		currentType := parser.Peek().ttype
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
	parser.Error(parser.Peek(), errorMessage)
	return parser.Peek()
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
	return parser.Peek().ttype == TokenTypeEof ||
		parser.current >= len(parser.tokens)
}

func (parser *Parser) Error(token Token, errorMessage string) {
	fmt.Println("[Syntax Error] " + errorMessage)
	os.Exit(1)
}

//------------------------------------------------------------------------------
// AST STARTS HERE
//------------------------------------------------------------------------------
func (parser *Parser) Statement() Expression {
	if !parser.Check(TokenTypeLeftParen) {
		parser.Error(parser.Peek(), "Expected opening '(' before expression call")
	}
	return parser.ListExpression()
}

func (parser *Parser) ListExpression() Expression {
	if parser.Match(TokenTypeLeftParen) {
		args := make([]Expression, 0)
		for !parser.Check(TokenTypeRightParen) {
			args = append(args, parser.ListExpression())
		}
		parser.Consume("Expected closing ')' after list", TokenTypeRightParen)
		return NewExpressionList(args)
	} else {
		return parser.AtomExpression()
	}
}

func (parser *Parser) AtomExpression() Expression {
	atom := parser.Peek()
	if atom.ttype != TokenTypeEof {
		parser.Advance()
		return NewExpressionAtom(atom)
	}
	parser.Error(parser.Previous(), "Unexpected end of file")
	return nil
}
