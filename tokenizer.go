package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

// Token and TokenTypes

type TokenType int

// Token definitions
const (
	// parser tokens
	TokenTypeEof = 0

	// single character tokens
	TokenTypeLeftBrace    = 1
	TokenTypeLeftBracket  = 2
	TokenTypeLeftParen    = 3
	TokenTypePercent      = 4
	TokenTypeRightBrace   = 5
	TokenTypeRightBracket = 6
	TokenTypeRightParen   = 7
	TokenTypeSlash        = 8
	TokenTypeStar         = 9
	TokenTypeComma        = 10
	TokenTypeDot          = 11
	TokenTypeSemicolon    = 12

	// one or two character tokens
	TokenTypeArrow        = 13
	TokenTypeBang         = 14
	TokenTypeBangEqual    = 15
	TokenTypeColon        = 16
	TokenTypeEqual        = 17
	TokenTypeEqualEqual   = 18
	TokenTypeGreater      = 19
	TokenTypeGreaterEqual = 20
	TokenTypeLess         = 21
	TokenTypeLessEqual    = 22
	TokenTypeMinus        = 23
	TokenTypeMinusEqual   = 24
	TokenTypePercentEqual = 25
	TokenTypePlus         = 26
	TokenTypePlusEqual    = 27
	TokenTypeSlashEqual   = 28
	TokenTypeStarEqual    = 29

	// literals
	TokenTypeReserved   = 30
	TokenTypeIdentifier = 31
	TokenTypeString     = 32
	TokenTypeNull       = 33
	TokenTypeBoolean    = 34
	TokenTypeTrue       = 35
	TokenTypeFalse      = 36
	TokenTypeInteger    = 37
	TokenTypeFloat      = 38
)

// Reserved words dictionary
var ReservedTokens = map[string]string{
	"if":    "if",
	"while": "while",
	"print": "print",
	"write": "print",
}

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
		tokenizer.Error(err.Error())
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

func (tokenizer *Tokenizer) Error(errorMessage string) {
	fmt.Println("[Scan Error] " + errorMessage)
	os.Exit(1)
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

func (tokenizer *Tokenizer) String(quote rune) {

	for tokenizer.Peek() != quote && !tokenizer.Eof() {
		tokenizer.Advance()
	}
	if tokenizer.Eof() {
		tokenizer.Error("Unterminated string, expecting closing quote")
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
	token := string(tokenizer.source[tokenizer.start:tokenizer.current])
	reserved, ok := ReservedTokens[token]
	if ok {
		tokenizer.AddToken(TokenTypeReserved, reserved)
	} else {
		tokenizer.AddToken(TokenTypeIdentifier, token)
	}
}

func (tokenizer *Tokenizer) Number() {
	isFloat := false

	for unicode.IsDigit(tokenizer.Peek()) {
		tokenizer.Advance()
	}

	if tokenizer.Match('.') && unicode.IsDigit(tokenizer.Peek()) {
		isFloat = true
		for unicode.IsDigit(tokenizer.Peek()) {
			tokenizer.Advance()
		}
	}

	var tokenType TokenType
	if isFloat {
		tokenType = TokenTypeFloat
	} else {
		tokenType = TokenTypeInteger
	}
	tokenizer.AddToken(tokenType, string(tokenizer.source[tokenizer.start:tokenizer.current]))
}

func (tokenizer *Tokenizer) twoChar(char rune) bool {
	return (char == '<' && tokenizer.Match('=')) ||
		(char == '>' && tokenizer.Match('=')) ||
		(char == '!' && tokenizer.Match('=')) ||
		(char == '=' && tokenizer.Match('=')) ||
		(char == ':' && tokenizer.Match('=')) ||
		(char == '/' && tokenizer.Match('='))
}

func (tokenizer *Tokenizer) oneChar(char rune) bool {
	return char == '*' || char == '+' || char == '-' || char == '<' ||
		char == '>' || char == '!' || char == '=' || char == '/'
}

func (tokenizer *Tokenizer) ignoreChar(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func (tokenizer *Tokenizer) ScanToken() {
	char := rune(tokenizer.Advance())
	switch {
	case char == '(':
		tokenizer.AddToken(TokenTypeLeftParen, "(")
	case char == ')':
		tokenizer.AddToken(TokenTypeRightParen, ")")
	case tokenizer.twoChar(char):
		tokenizer.AddToken(TokenTypeReserved, string(char)+"=")
	case tokenizer.oneChar(char):
		tokenizer.AddToken(TokenTypeReserved, string(char))
	case char == '"' || char == '\'':
		tokenizer.String(char)
	case unicode.IsDigit(char):
		tokenizer.Number()
	case unicode.IsLetter(char):
		tokenizer.Identifier()
	case char == ';':
		tokenizer.Comment()
	case tokenizer.ignoreChar(char):
		break
	default:
		tokenizer.Error("[Tokenizer] Unexpected character: " + string(char))

	}
}
