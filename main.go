package main

import (
	"fmt"
	"wok/woklang"
)

func main() {
	v := woklang.Eval("(debug true)")
	fmt.Print(v.ToString())
	tokenizer := woklang.MakeTokenizer()
	tokenizer.LoadFromFile("demo.lisp")
	tokens := tokenizer.Tokenize()

	parser := woklang.MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := woklang.MakeInterpreter()
	interpreter.Interpret(expressions)
}
