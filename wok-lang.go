package main

import "fmt"

func main() {
	/*
		tokenizer := MakeTokenizer()
		tokenizer.LoadFromFile("demo.wok")
		tokenizer.Tokenize()
		parser := MakeParser()
		parser.Parse(tokenizer.tokens)
		fmt.Println("Dones")
	*/

	num := NewWokInteger(22)
	num1 := NewWokInteger(22)
	same := num1.Equals(num)
	fmt.Println(same)

	fmt.Println("Dones")
}
