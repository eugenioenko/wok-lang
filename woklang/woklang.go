package woklang

func Eval(source string) WokData {
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromString(source)
	tokens := tokenizer.Tokenize()

	parser := MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := MakeInterpreter()
	result := interpreter.Interpret(expressions)
	return result
}
