package main

func WokEvaluate(source string) WokData {
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromString(source)
	tokenizer.Tokenize()

	parser := MakeParser()
	parser.Parse(tokenizer.tokens)

	interpreter := MakeInterpreter()
	return interpreter.Interpret(parser.expressions)
}

func main() {

	tokenizer := MakeTokenizer()
	tokenizer.LoadFromFile("demo.lisp")
	tokenizer.Tokenize()

	parser := MakeParser()
	parser.Parse(tokenizer.tokens)

	interpreter := MakeInterpreter()
	interpreter.Interpret(parser.expressions)
}
