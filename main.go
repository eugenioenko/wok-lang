package main

import "woklang/woklang"

func WokEvaluate(source string) woklang.WokData {
	tokenizer := woklang.MakeTokenizer()
	tokenizer.LoadFromString(source)
	tokenizer.Tokenize()

	parser := woklang.MakeParser()
	parser.Parse(tokenizer.tokens)

	interpreter := woklang.MakeInterpreter()
	return interpreter.Interpret(parser.expressions)
}

func main() {

	tokenizer := woklang.MakeTokenizer()
	tokenizer.LoadFromFile("demo.lisp")
	tokenizer.Tokenize()

	parser := woklang.MakeParser()
	parser.Parse(tokenizer.tokens)

	interpreter := woklang.MakeInterpreter()
	interpreter.Interpret(parser.expressions)
}
