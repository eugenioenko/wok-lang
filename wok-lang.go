package main

import (
	"fmt"
)

func main() {
	tknzr := MakeTokenizer()
	tknzr.LoadFromFile("demo.wok")
	tknzr.Tokenize()
	fmt.Println("Dones")
}
