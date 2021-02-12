package main

import "fmt"

/*
func call() (result string) {

	defer func() {
		if r := recover(); r != nil {
			result = r.(string)
		}
	}()
	panic("works!")
}
*/
func main() {
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromFile("demo.wok")
	tokenizer.Tokenize()

	parser := MakeParser()
	parser.Parse(tokenizer.tokens)

	interpreter := MakeInterpreter()
	interpreter.Interpret(parser.statements)

	fmt.Println("Dones")
}
