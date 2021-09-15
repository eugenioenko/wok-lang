package main

func main() {

	tokenizer := MakeTokenizer()
	tokenizer.LoadFromFile("demo.lisp")
	tokenizer.Tokenize()

	parser := MakeParser()
	parser.Parse(tokenizer.tokens)

	interpreter := MakeInterpreter()
	interpreter.Interpret(parser.expressions)
}
