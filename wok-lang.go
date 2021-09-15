package main

func main() {

	tokenizer := MakeTokenizer()
	tokenizer.LoadFromFile("demo.lisp")
	tokenizer.Tokenize()

	parser := MakeParser()
	parser.Parse(tokenizer.tokens)

	println("end")

	// interpreter := MakeInterpreter()
	// interpreter.Interpret(parser.expressions)
}
