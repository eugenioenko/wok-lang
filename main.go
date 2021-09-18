package main

import "wok/woklang"

func main() {

	tokenizer := woklang.MakeTokenizer()
	tokenizer.LoadFromFile("demo.lisp")
	tokens := tokenizer.Tokenize()

	parser := woklang.MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := woklang.MakeInterpreter()
	interpreter.Interpret(expressions)
}
