package woklang

import (
	"fmt"
)

func Eval(source string) (result WokData) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[Runtime Error] Oops! Unhandled Error")
			result = NewWokException("Unhandled Exception")
		}
	}()
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromString(source)
	tokens := tokenizer.Tokenize()

	parser := MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := MakeInterpreter()
	result = interpreter.Interpret(expressions)
	return result
}

func Exec(filename string) (result WokData) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[Runtime Error] Oops! Unhandled Error")
			result = NewWokException("Unhandled Exception")
		}
	}()
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromFile(filename)
	tokens := tokenizer.Tokenize()

	parser := MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := MakeInterpreter()
	return interpreter.Interpret(expressions)
}
