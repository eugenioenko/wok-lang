package woklang

// Token and TokenTypes

type TokenType int

// Token definitions
const (
	// parser tokens
	TokenTypeEof  = -1
	TokenTypeNull = 0

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
	TokenTypeBoolean    = 34
	TokenTypeTrue       = 35
	TokenTypeFalse      = 36
	TokenTypeInteger    = 37
	TokenTypeFloat      = 38
)

type Token struct {
	Type    TokenType
	literal string
}

func MakeToken(Type TokenType, literal string) Token {
	return Token{Type, literal}
}
