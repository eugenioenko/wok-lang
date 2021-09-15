package main

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
	fmt.Println("[Syntax Error] " + errorMessage)
	os.Exit(1)
}

//------------------------------------------------------------------------------
// AST STARTS HERE
//------------------------------------------------------------------------------
func (parser *Parser) Statement() Expression {
	return parser.ListExpression()
}

func (parser *Parser) ListExpression() Expression {
	if parser.Match(TokenTypeLeftParen) {
		args := make([]Expression, 0)
		for !parser.Check(TokenTypeRightParen) {
			args = append(args, parser.ListExpression())
		}
		parser.Consume("Expected closing ')' after function call", TokenTypeRightParen)
		return NewExpressionList(args)
	} else {
		return parser.AtomExpression()
	}
}

func (parser *Parser) AtomExpression() Expression {
	switch {
	case parser.Match(TokenTypeInteger),
		parser.Match(TokenTypeFloat),
		parser.Match(TokenTypeString),
		parser.Match(TokenTypeIdentifier):
		return NewExpressionAtom(parser.Previous())
	}
	parser.Error(parser.Peek(), "Unexpected token '"+parser.Peek().literal+"'")
	return nil
}
