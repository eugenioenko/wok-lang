package woklang

var RuntimeTokens = map[string]string{
	"if":          "if",
	"while":       "while",
	"print":       "print",
	"write":       "print",
	"cond":        "cond",
	"debug":       "debug",
	"func":        "func",
	"return":      "return",
	"return-from": "return-from",
}

var ReservedTokens = map[string]TokenType{
	"null":  TokenTypeNull,
	"true":  TokenTypeTrue,
	"false": TokenTypeFalse,
}

var RuntimeScope = map[string]WokData{
	"print":       WF("print", RuntimePrint),
	"cond":        WF("cond", RuntimeCond),
	"while":       WF("while", RuntimeWhile),
	"debug":       WF("debug", RuntimeDebug),
	"func":        WF("func", RuntimeFunc),
	"return-from": WF("return-from", RuntimeReturnFrom),
	"return":      WF("return", RuntimeReturn),
	"if":          WF("if", RuntimeIf),
	":=":          WF(":=", RuntimeAssignment),
	"==":          WF("==", RuntimeEquality),
	"!=":          WF("==", RuntimeInequality),
	"!":           WF("!", RuntimeNegation),
	"+":           WF("+", RuntimeAddition),
	"*":           WF("*", RuntimeMultiplication),
	"-":           WF("-", RuntimeSubstraction),
}
