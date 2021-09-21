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

func Exec(filename string) WokData {
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromFile(filename)
	tokens := tokenizer.Tokenize()

	parser := MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := MakeInterpreter()
	return interpreter.Interpret(expressions)
}
