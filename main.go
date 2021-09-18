package main

import woklang "wok/src"

func WokEvaluate(source string) woklang.WokData {
	tokenizer := woklang.MakeTokenizer()
	tokenizer.LoadFromString(source)
	tokens := tokenizer.Tokenize()

	parser := woklang.MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := woklang.MakeInterpreter()
	return interpreter.Interpret(expressions)
}

func main() {

	tokenizer := woklang.MakeTokenizer()
	tokenizer.LoadFromFile("demo.lisp")
	tokens := tokenizer.Tokenize()

	parser := woklang.MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := woklang.MakeInterpreter()
	interpreter.Interpret(expressions)
}
