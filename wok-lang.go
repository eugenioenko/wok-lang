package main

import (
	"fmt"
)

func main() {
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromFile("demo.wok")
	tokenizer.Tokenize()
	parser := MakeParser()
	parser.Parse(tokenizer.tokens)
	fmt.Println("Dones")
}
