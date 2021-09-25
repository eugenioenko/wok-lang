package woklang

var RuntimeTokens = map[string]string{
	"if":    "if",
	"while": "while",
	"print": "print",
	"write": "print",
	"cond":  "cond",
	"debug": "debug",
	"defun": "defun",
}

var ReservedTokens = map[string]TokenType{
	"null":  TokenTypeNull,
	"true":  TokenTypeTrue,
	"false": TokenTypeFalse,
}
