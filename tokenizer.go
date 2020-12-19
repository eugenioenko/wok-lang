package main

import (
	"io/ioutil"
	"unicode"
)

// Token and TokenTypes

type TokenType int

const (
	// parser tokens
	TokenTypeEof = iota
	// single character tokens
	TokenTypeLeftBrace
	TokenTypeLeftBracket
	TokenTypeLeftParen
	TokenTypePercent
	TokenTypeRightBrace
	TokenTypeRightBracket
	TokenTypeRightParen
	TokenTypeSlash
	TokenTypeStar
	TokenTypeComma
	TokenTypeDot
	TokenTypeSemicolon

	// one or two character tokens
	TokenTypeArrow
	TokenTypeBang
	TokenTypeBangEqual
	TokenTypeColon
	TokenTypeEqual
	TokenTypeEqualEqual
	TokenTypeGreater
	TokenTypeGreaterEqual
	TokenTypeLess
	TokenTypeLessEqual
	TokenTypeMinus
	TokenTypeMinusEqual
	TokenTypePercentEqual
	TokenTypePlus
	TokenTypePlusEqual
	TokenTypeSlashEqual
	TokenTypeStarEqual

	// literals
	TokenTypeIdentifier
	TokenTypeString
	TokenTypeNull
	TokenTypeTrue
	TokenTypeFalse
	TokenTypeNumber
)

type Token struct {
	ttype   TokenType
	literal string
}

func MakeToken(ttype TokenType, literal string) Token {
	return Token{ttype, literal}
}

// Tokenizer
// Converts a source file into an array of tokens
type Tokenizer struct {
	source  []byte
	current int
	start   int
	tokens  []Token
}

func MakeTokenizer() Tokenizer {
	return Tokenizer{}
}

func (tokenizer *Tokenizer) LoadFromFile(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	tokenizer.source = content
	tokenizer.tokens = make([]Token, 0)
}

func (tokenizer *Tokenizer) Eof() bool {
	return tokenizer.current >= len(tokenizer.source)
}

func (tokenizer *Tokenizer) Advance() byte {
	var current byte = tokenizer.source[tokenizer.current]
	tokenizer.current += 1
	return current
}

func (tokenizer *Tokenizer) Match(expected byte) bool {
	if tokenizer.Eof() {
		return false
	}
	if tokenizer.source[tokenizer.current] != expected {
		return false
	}
	tokenizer.current += 1
	return true
}

func (tokenizer *Tokenizer) Peek() rune {
	if tokenizer.Eof() {
		return 0
	}
	return rune(tokenizer.source[tokenizer.current])
}

func (tokenizer *Tokenizer) PeekNext() rune {
	if tokenizer.current+1 >= len(tokenizer.source) {
		return 0
	}
	return rune(tokenizer.source[tokenizer.current+1])
}

func (tokenizer *Tokenizer) AddToken(ttype TokenType, literal string) {
	tokenizer.tokens = append(tokenizer.tokens, MakeToken(ttype, literal))

}

func (tokenizer *Tokenizer) Tokenize() {
	tokenizer.current = 0
	tokenizer.start = 0

	for !tokenizer.Eof() {
		tokenizer.start = tokenizer.current
		tokenizer.ScanToken()
	}
	tokenizer.AddToken(TokenTypeEof, "")
}

func (tokenizer *Tokenizer) Comment() {
	for tokenizer.Peek() != '\n' && !tokenizer.Eof() {
		tokenizer.Advance()
	}
}

func (tokenizer *Tokenizer) String() {
	for tokenizer.Peek() != '"' && !tokenizer.Eof() {
		tokenizer.Advance()
	}
	if tokenizer.Eof() {
		panic("Unterminated string, expecting clsoing quote")
	}
	// the closing quote
	tokenizer.Advance()
	tokenizer.AddToken(TokenTypeString, string(tokenizer.source[tokenizer.start+1:tokenizer.current-1]))
}

func (tokenizer *Tokenizer) Identifier() {
	for unicode.IsLetter(tokenizer.Peek()) ||
		unicode.IsDigit(tokenizer.Peek()) ||
		tokenizer.Peek() == '_' {
		tokenizer.Advance()
	}
	tokenizer.AddToken(TokenTypeIdentifier, string(tokenizer.source[tokenizer.start:tokenizer.current]))
}

func (tokenizer *Tokenizer) Number() {
	for unicode.IsDigit(tokenizer.Peek()) {
		tokenizer.Advance()
	}
	tokenizer.AddToken(TokenTypeNumber, string(tokenizer.source[tokenizer.start:tokenizer.current]))
}

func (tokenizer *Tokenizer) ScanToken() {
	var char rune = rune(tokenizer.Advance())
	switch {
	case char == '(':
		tokenizer.AddToken(TokenTypeLeftParen, "(")
	case char == ')':
		tokenizer.AddToken(TokenTypeRightParen, ")")
	case char == '[':
		tokenizer.AddToken(TokenTypeLeftBracket, "[")
	case char == ']':
		tokenizer.AddToken(TokenTypeRightBracket, "]")
	case char == '{':
		tokenizer.AddToken(TokenTypeLeftBrace, "{")
	case char == '}':
		tokenizer.AddToken(TokenTypeRightBrace, "}")
	case char == ',':
		tokenizer.AddToken(TokenTypeComma, ",")
	case char == ':':
		tokenizer.AddToken(TokenTypeColon, ":")
	case char == ';':
		tokenizer.AddToken(TokenTypeSemicolon, ";")
	case char == '.':
		tokenizer.AddToken(TokenTypeDot, ".")
	case char == '*' && tokenizer.Match('='):
		tokenizer.AddToken(TokenTypeStarEqual, "*=")
	case char == '*':
		tokenizer.AddToken(TokenTypeStar, "*")
	case char == '<' && tokenizer.Match('='):
		tokenizer.AddToken(TokenTypeLessEqual, "<=")
	case char == '<':
		tokenizer.AddToken(TokenTypeLess, "<")
	case char == '>' && tokenizer.Match('='):
		tokenizer.AddToken(TokenTypeGreaterEqual, ">=")
	case char == '>':
		tokenizer.AddToken(TokenTypeGreater, ">")
	case char == '!' && tokenizer.Match('='):
		tokenizer.AddToken(TokenTypeBangEqual, "!=")
	case char == '!':
		tokenizer.AddToken(TokenTypeBang, "!")
	case char == '=' && tokenizer.Match('='):
		tokenizer.AddToken(TokenTypeEqualEqual, "==")
	case char == '=':
		tokenizer.AddToken(TokenTypeEqual, "=")
	case char == '/' && tokenizer.Match('/'):
		tokenizer.Comment()
	case char == '/' && tokenizer.Match('='):
		tokenizer.AddToken(TokenTypeSlashEqual, "/=")
	case char == '"':
		tokenizer.String()
	case unicode.IsDigit(char):
		tokenizer.Number()
	case unicode.IsLetter(char):
		tokenizer.Identifier()
	case char == ' ':
	case char == '\t':
	case char == '\n':
	case char == '\r':
	default:
		panic("Unexpected character: " + string(char))

	}

}
